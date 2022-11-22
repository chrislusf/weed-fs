package s3api

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/seaweedfs/seaweedfs/weed/glog"
	"github.com/seaweedfs/seaweedfs/weed/pb/filer_pb"
	"github.com/seaweedfs/seaweedfs/weed/s3api/s3_constants"
	"github.com/seaweedfs/seaweedfs/weed/s3api/s3account"
	"github.com/seaweedfs/seaweedfs/weed/s3api/s3acl"
	"github.com/seaweedfs/seaweedfs/weed/s3api/s3err"
	"github.com/seaweedfs/seaweedfs/weed/util"
	"net/http"
	"path/filepath"
)

func getAccountId(r *http.Request) string {
	id := r.Header.Get(s3_constants.AmzAccountId)
	if len(id) == 0 {
		return s3account.AccountAnonymous.Id
	} else {
		return id
	}
}

func (s3a *S3ApiServer) checkAccessByOwnership(r *http.Request, bucket string) s3err.ErrorCode {
	metadata, errCode := s3a.bucketRegistry.GetBucketMetadata(bucket)
	if errCode != s3err.ErrNone {
		return errCode
	}
	accountId := getAccountId(r)
	if accountId == s3account.AccountAdmin.Id || accountId == *metadata.Owner.ID {
		return s3err.ErrNone
	}
	return s3err.ErrAccessDenied
}

//Check access for PutBucketAclHandler
func (s3a *S3ApiServer) checkAccessForPutBucketAcl(accountId, bucket string) (*BucketMetaData, s3err.ErrorCode) {
	bucketMetadata, errCode := s3a.bucketRegistry.GetBucketMetadata(bucket)
	if errCode != s3err.ErrNone {
		return nil, errCode
	}

	if bucketMetadata.ObjectOwnership == s3_constants.OwnershipBucketOwnerEnforced {
		return nil, s3err.AccessControlListNotSupported
	}

	if accountId == s3account.AccountAdmin.Id || accountId == *bucketMetadata.Owner.ID {
		return bucketMetadata, s3err.ErrNone
	}

	if len(bucketMetadata.Acl) > 0 {
		reqGrants := s3acl.DetermineReqGrants(accountId, s3_constants.PermissionWriteAcp)
		for _, bucketGrant := range bucketMetadata.Acl {
			for _, reqGrant := range reqGrants {
				if s3acl.GrantEquals(bucketGrant, reqGrant) {
					return bucketMetadata, s3err.ErrNone
				}
			}
		}
	}
	glog.V(3).Infof("acl denied! request account id: %s", accountId)
	return nil, s3err.ErrAccessDenied
}

func updateBucketEntry(s3a *S3ApiServer, entry *filer_pb.Entry) error {
	return s3a.updateEntry(s3a.option.BucketsPath, entry)
}

// Check Bucket/BucketAcl Read related access
// includes:
// - HeadBucketHandler
// - GetBucketAclHandler
// - ListObjectsV1Handler
// - ListObjectsV2Handler
func (s3a *S3ApiServer) checkAccessForReadBucket(r *http.Request, bucket, aclAction string) (*BucketMetaData, s3err.ErrorCode) {
	bucketMetadata, errCode := s3a.bucketRegistry.GetBucketMetadata(bucket)
	if errCode != s3err.ErrNone {
		return nil, errCode
	}

	if bucketMetadata.ObjectOwnership == s3_constants.OwnershipBucketOwnerEnforced {
		return bucketMetadata, s3err.ErrNone
	}

	accountId := s3acl.GetAccountId(r)
	if accountId == s3account.AccountAdmin.Id || accountId == *bucketMetadata.Owner.ID {
		return bucketMetadata, s3err.ErrNone
	}

	if len(bucketMetadata.Acl) > 0 {
		reqGrants := s3acl.DetermineReqGrants(accountId, aclAction)
		for _, bucketGrant := range bucketMetadata.Acl {
			for _, reqGrant := range reqGrants {
				if s3acl.GrantEquals(bucketGrant, reqGrant) {
					return bucketMetadata, s3err.ErrNone
				}
			}
		}
	}

	glog.V(3).Infof("acl denied! request account id: %s", accountId)
	return nil, s3err.ErrAccessDenied
}

//Check ObjectAcl-Read related access
// includes:
// - GetObjectAclHandler
func (s3a *S3ApiServer) checkAccessForReadObjectAcl(r *http.Request, bucket, object string) (acp *s3.AccessControlPolicy, errCode s3err.ErrorCode) {
	bucketMetadata, errCode := s3a.bucketRegistry.GetBucketMetadata(bucket)
	if errCode != s3err.ErrNone {
		return nil, errCode
	}

	getAcpFunc := func() (*s3.AccessControlPolicy, s3err.ErrorCode) {
		entry, err := getObjectEntry(s3a, bucket, object)
		if err != nil {
			if err == filer_pb.ErrNotFound {
				return nil, s3err.ErrNoSuchKey
			} else {
				return nil, s3err.ErrInternalError
			}
		}

		if entry.IsDirectory {
			return nil, s3err.ErrExistingObjectIsDirectory
		}

		acpOwnerId := s3acl.GetAcpOwner(entry.Extended, *bucketMetadata.Owner.ID)
		acpOwnerName := s3a.accountManager.IdNameMapping[acpOwnerId]
		acpGrants := s3acl.GetAcpGrants(entry.Extended)
		acp = &s3.AccessControlPolicy{
			Owner: &s3.Owner{
				ID:          &acpOwnerId,
				DisplayName: &acpOwnerName,
			},
			Grants: acpGrants,
		}
		return acp, s3err.ErrNone
	}

	if bucketMetadata.ObjectOwnership == s3_constants.OwnershipBucketOwnerEnforced {
		return getAcpFunc()
	} else {
		accountId := s3acl.GetAccountId(r)

		acp, errCode := getAcpFunc()
		if errCode != s3err.ErrNone {
			return nil, errCode
		}

		if accountId == *acp.Owner.ID {
			return acp, s3err.ErrNone
		}

		//find in Grants
		if acp.Grants != nil {
			reqGrants := s3acl.DetermineReqGrants(accountId, s3_constants.PermissionReadAcp)
			for _, requiredGrant := range reqGrants {
				for _, grant := range acp.Grants {
					if s3acl.GrantEquals(requiredGrant, grant) {
						return acp, s3err.ErrNone
					}
				}
			}
		}

		glog.V(3).Infof("acl denied! request account id: %s", accountId)
		return nil, s3err.ErrAccessDenied
	}
}

// Check Object-Read related access
// includes:
// - GetObjectHandler
//
// offload object access validation to Filer layer
// - s3acl.CheckObjectAccessForReadObject
func (s3a *S3ApiServer) checkBucketAccessForReadObject(r *http.Request, bucket string) s3err.ErrorCode {
	bucketMetadata, errCode := s3a.bucketRegistry.GetBucketMetadata(bucket)
	if errCode != s3err.ErrNone {
		return errCode
	}

	if bucketMetadata.ObjectOwnership != s3_constants.OwnershipBucketOwnerEnforced {
		//offload object acl validation to filer layer
		r.Header.Set(s3_constants.XSeaweedFSHeaderAmzBucketOwnerId, *bucketMetadata.Owner.ID)
	}

	return s3err.ErrNone
}

// Check ObjectAcl-Write related access
// includes:
// - PutObjectAclHandler
func (s3a *S3ApiServer) checkAccessForWriteObjectAcl(accountId, bucket, object string) (bucketMetadata *BucketMetaData, objectEntry *filer_pb.Entry, objectOwner string, errCode s3err.ErrorCode) {
	bucketMetadata, errCode = s3a.bucketRegistry.GetBucketMetadata(bucket)
	if errCode != s3err.ErrNone {
		return nil, nil, "", errCode
	}

	if bucketMetadata.ObjectOwnership == s3_constants.OwnershipBucketOwnerEnforced {
		return nil, nil, "", s3err.AccessControlListNotSupported
	}

	//bucket acl
	bucketAclAllowed := false
	reqGrants := s3acl.DetermineReqGrants(accountId, s3_constants.PermissionWrite)
	if accountId == *bucketMetadata.Owner.ID {
		bucketAclAllowed = true
	} else if bucketMetadata.Acl != nil {
	bucketLoop:
		for _, bucketGrant := range bucketMetadata.Acl {
			for _, requiredGrant := range reqGrants {
				if s3acl.GrantEquals(bucketGrant, requiredGrant) {
					bucketAclAllowed = true
					break bucketLoop
				}
			}
		}
	}
	if !bucketAclAllowed {
		return nil, nil, "", s3err.ErrAccessDenied
	}

	//object acl
	objectEntry, err := getObjectEntry(s3a, bucket, object)
	if err != nil {
		if err == filer_pb.ErrNotFound {
			return nil, nil, "", s3err.ErrNoSuchKey
		}
		return nil, nil, "", s3err.ErrInternalError
	}

	if objectEntry.IsDirectory {
		return nil, nil, "", s3err.ErrExistingObjectIsDirectory
	}

	objectOwner = s3acl.GetAcpOwner(objectEntry.Extended, *bucketMetadata.Owner.ID)
	if accountId == objectOwner {
		return bucketMetadata, objectEntry, objectOwner, s3err.ErrNone
	}

	objectGrants := s3acl.GetAcpGrants(objectEntry.Extended)
	if objectGrants != nil {
		for _, objectGrant := range objectGrants {
			for _, requiredGrant := range reqGrants {
				if s3acl.GrantEquals(objectGrant, requiredGrant) {
					return bucketMetadata, objectEntry, objectOwner, s3err.ErrNone
				}
			}
		}
	}

	glog.V(3).Infof("acl denied! request account id: %s", accountId)
	return nil, nil, "", s3err.ErrAccessDenied
}

func updateObjectEntry(s3a *S3ApiServer, bucket, object string, entry *filer_pb.Entry) error {
	dir, _ := filepath.Split(object)
	return s3a.updateEntry(util.Join(s3a.option.BucketsPath, bucket, dir), entry)
}

// CheckAccessForPutObject Check ACL for PutObject API
// includes:
// - PutObjectHandler
func (s3a *S3ApiServer) CheckAccessForPutObject(r *http.Request, bucket, object string) s3err.ErrorCode {
	accountId := s3acl.GetAccountId(r)
	return s3a.checkAccessForWriteObject(r, bucket, object, accountId)
}

// CheckAccessForNewMultipartUpload Check Acl for InitiateMultipartUploadResult API
// includes:
// - NewMultipartUploadHandler
func (s3a *S3ApiServer) CheckAccessForNewMultipartUpload(r *http.Request, bucket, object string) s3err.ErrorCode {
	accountId := s3acl.GetAccountId(r)
	if accountId == IdentityAnonymous.AccountId {
		return s3err.ErrAccessDenied
	}
	return s3a.checkAccessForWriteObject(r, bucket, object, accountId)
}

// CheckAccessForCompleteMultipartUpload Check Acl for CompleteMultipartUpload API
// includes:
// - CompleteMultipartUploadHandler
func (s3a *S3ApiServer) CheckAccessForCompleteMultipartUpload(r *http.Request, bucket, object string) s3err.ErrorCode {
	bucketMetadata, errCode := s3a.bucketRegistry.GetBucketMetadata(bucket)
	if errCode != s3err.ErrNone {
		return errCode
	}

	//bucket access allowed
	accountId := s3acl.GetAccountId(r)
	if accountId == *bucketMetadata.Owner.ID {
		return s3err.ErrNone
	} else {
		if len(bucketMetadata.Acl) > 0 {
			reqGrants := s3acl.DetermineReqGrants(accountId, s3_constants.PermissionWrite)
			for _, bucketGrant := range bucketMetadata.Acl {
				for _, requiredGrant := range reqGrants {
					if s3acl.GrantEquals(bucketGrant, requiredGrant) {
						return s3err.ErrNone
					}
				}
			}
		}
	}
	glog.V(3).Infof("acl denied! request account id: %s", accountId)
	return s3err.ErrAccessDenied
}

func (s3a *S3ApiServer) checkAccessForWriteObject(r *http.Request, bucket, object, accountId string) s3err.ErrorCode {
	bucketMetadata, errCode := s3a.bucketRegistry.GetBucketMetadata(bucket)
	if errCode != s3err.ErrNone {
		return errCode
	}

	if bucketMetadata.ObjectOwnership == s3_constants.OwnershipBucketOwnerEnforced {
		// validate grants (only bucketOwnerFullControl acl is allowed)
		_, grants, errCode := s3acl.ParseAndValidateAclHeaders(r, s3a.accountManager, bucketMetadata.ObjectOwnership, *bucketMetadata.Owner.ID, accountId, false)
		if errCode != s3err.ErrNone {
			return errCode
		}
		if len(grants) > 1 {
			return s3err.AccessControlListNotSupported
		}
		bucketOwnerFullControlGrant := &s3.Grant{
			Permission: &s3_constants.PermissionFullControl,
			Grantee: &s3.Grantee{
				Type: &s3_constants.GrantTypeCanonicalUser,
				ID:   bucketMetadata.Owner.ID,
			},
		}
		if len(grants) == 0 {
			return s3err.ErrNone
		}

		if s3acl.GrantEquals(bucketOwnerFullControlGrant, grants[0]) {
			return s3err.ErrNone
		}

		return s3err.AccessControlListNotSupported
	}

	//bucket access allowed
	bucketAclAllowed := false
	if accountId == *bucketMetadata.Owner.ID {
		bucketAclAllowed = true
	} else {
		if len(bucketMetadata.Acl) > 0 {
			reqGrants := s3acl.DetermineReqGrants(accountId, s3_constants.PermissionWrite)
		bucketLoop:
			for _, bucketGrant := range bucketMetadata.Acl {
				for _, requiredGrant := range reqGrants {
					if s3acl.GrantEquals(bucketGrant, requiredGrant) {
						bucketAclAllowed = true
						break bucketLoop
					}
				}
			}
		}
	}
	if !bucketAclAllowed {
		glog.V(3).Infof("acl denied! request account id: %s", accountId)
		return s3err.ErrAccessDenied
	}

	//object access allowed
	entry, err := getObjectEntry(s3a, bucket, object)
	if err != nil {
		if err != filer_pb.ErrNotFound {
			return s3err.ErrInternalError
		}
	} else {
		if entry.IsDirectory {
			return s3err.ErrExistingObjectIsDirectory
		}

		//Only the owner of the bucket and the owner of the object can overwrite the object
		objectOwner := s3acl.GetAcpOwner(entry.Extended, *bucketMetadata.Owner.ID)
		if accountId != objectOwner && accountId != *bucketMetadata.Owner.ID {
			glog.V(3).Infof("acl denied! request account id: %s, expect account id: %s", accountId, *bucketMetadata.Owner.ID)
			return s3err.ErrAccessDenied
		}
	}

	ownerId, grants, errCode := s3acl.ParseAndValidateAclHeadersOrElseDefault(r, s3a.accountManager, bucketMetadata.ObjectOwnership, *bucketMetadata.Owner.ID, accountId, false)
	if errCode != s3err.ErrNone {
		return errCode
	}

	s3acl.SetAcpOwnerHeader(r, ownerId)
	s3acl.SetAcpGrantsHeader(r, grants)

	return s3err.ErrNone
}

func getObjectEntry(s3a *S3ApiServer, bucket, object string) (*filer_pb.Entry, error) {
	return s3a.getEntry(util.Join(s3a.option.BucketsPath, bucket), object)
}

func (s3a *S3ApiServer) ExtractBucketAcp(r *http.Request) (owner string, grants []*s3.Grant, errCode s3err.ErrorCode) {
	accountId := s3acl.GetAccountId(r)

	ownership := s3_constants.DefaultOwnershipForCreate
	if ownership == s3_constants.OwnershipBucketOwnerEnforced {
		return accountId, []*s3.Grant{
			{
				Permission: &s3_constants.PermissionFullControl,
				Grantee: &s3.Grantee{
					Type: &s3_constants.GrantTypeCanonicalUser,
					ID:   &accountId,
				},
			},
		}, s3err.ErrNone
	} else {
		return s3acl.ParseAndValidateAclHeadersOrElseDefault(r, s3a.accountManager, ownership, accountId, accountId, false)
	}
}
