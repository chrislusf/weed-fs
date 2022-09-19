/*
 * MinIO Cloud Storage, (C) 2019 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package s3_constants

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

// Standard S3 HTTP request constants
const (
	// S3 storage class
	AmzStorageClass = "x-amz-storage-class"

	// S3 user-defined metadata
	AmzUserMetaPrefix    = "X-Amz-Meta-"
	AmzUserMetaDirective = "X-Amz-Metadata-Directive"

	// S3 object tagging
	AmzObjectTagging          = "X-Amz-Tagging"
	AmzObjectTaggingPrefix    = "X-Amz-Tagging-"
	AmzObjectTaggingDirective = "X-Amz-Tagging-Directive"
	AmzTagCount               = "x-amz-tagging-count"

	// S3 ACL headers
	AmzCannedAcl      = "X-Amz-Acl"
	AmzAclFullControl = "X-Amz-Grant-Full-Control"
	AmzAclRead        = "X-Amz-Grant-Read"
	AmzAclWrite       = "X-Amz-Grant-Write"
	AmzAclReadAcp     = "X-Amz-Grant-Read-Acp"
	AmzAclWriteAcp    = "X-Amz-Grant-Write-Acp"

	X_SeaweedFS_Header_Directory_Key = "x-seaweedfs-is-directory-key"
)

// Non-Standard S3 HTTP request constants
const (
	AmzIdentityId        = "s3-identity-id"
	AmzIdentityAccountId = "s3-identity-account-id"
	AmzAuthType          = "s3-auth-type"
	AmzIsAdmin           = "s3-is-admin" // only set to http request header as a context
)

func GetBucketAndObject(r *http.Request) (bucket, object string) {
	vars := mux.Vars(r)
	bucket = vars["bucket"]
	object = vars["object"]
	if !strings.HasPrefix(object, "/") {
		object = "/" + object
	}

	return
}

var PassThroughHeaders = map[string]string{
	"response-cache-control":       "Cache-Control",
	"response-content-disposition": "Content-Disposition",
	"response-content-encoding":    "Content-Encoding",
	"response-content-language":    "Content-Language",
	"response-content-type":        "Content-Type",
	"response-expires":             "Expires",
}
