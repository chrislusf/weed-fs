package topology

import (
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/seaweedfs/seaweedfs/weed/glog"
	"github.com/seaweedfs/seaweedfs/weed/operation"
	"github.com/seaweedfs/seaweedfs/weed/security"
	"github.com/seaweedfs/seaweedfs/weed/stats"
	"github.com/seaweedfs/seaweedfs/weed/storage"
	"github.com/seaweedfs/seaweedfs/weed/storage/needle"
	"github.com/seaweedfs/seaweedfs/weed/storage/types"
	"github.com/seaweedfs/seaweedfs/weed/util"
	"github.com/seaweedfs/seaweedfs/weed/util/buffer_pool"
	util_http "github.com/seaweedfs/seaweedfs/weed/util/http"
)

type ReplicatedWriteParams struct {
	VolumeId needle.VolumeId
	Needle        *needle.Needle
	Jwt 	 security.EncodedJwt
	Replicate bool
	Fsync    bool
	ContentMd5 string
}

type ReplicatedDeleteParams struct {
	VolumeId needle.VolumeId
	Needle        *needle.Needle
	Jwt 	 security.EncodedJwt
	Replicate bool
}

func ReplicatedWrite(masterFn operation.GetMasterFn, grpcDialOption grpc.DialOption, s *storage.Store, params ReplicatedWriteParams) (isUnchanged bool, err error) {

	//check JWT
	jwt := params.Jwt

	// check whether this is a replicated write request
	var remoteLocations []operation.Location
	if !params.Replicate {
		// this is the initial request
		remoteLocations, err = GetWritableRemoteReplications(s, grpcDialOption, params.VolumeId, masterFn)
		if err != nil {
			glog.V(0).Infoln(err)
			return
		}
	}

	// read fsync value
	fsync := params.Fsync

	if s.GetVolume(params.VolumeId) != nil {
		start := time.Now()

		inFlightGauge := stats.VolumeServerInFlightRequestsGauge.WithLabelValues(stats.WriteToLocalDisk)
		inFlightGauge.Inc()
		defer inFlightGauge.Dec()

		isUnchanged, err = s.WriteVolumeNeedle(params.VolumeId, params.Needle, true, fsync)
		stats.VolumeServerRequestHistogram.WithLabelValues(stats.WriteToLocalDisk).Observe(time.Since(start).Seconds())
		if err != nil {
			stats.VolumeServerHandlerCounter.WithLabelValues(stats.ErrorWriteToLocalDisk).Inc()
			err = fmt.Errorf("failed to write to local disk: %v", err)
			glog.V(0).Infoln(err)
			return
		}
	}

	if len(remoteLocations) > 0 { //send to other replica locations
		start := time.Now()

		inFlightGauge := stats.VolumeServerInFlightRequestsGauge.WithLabelValues(stats.WriteToReplicas)
		inFlightGauge.Inc()
		defer inFlightGauge.Dec()

		err = DistributedOperation(remoteLocations, func(location operation.Location) error {
			fileId := needle.NewFileIdFromNeedle(params.VolumeId, params.Needle)

			u := url.URL{
				Scheme: "http",
				Host:   location.Url,
				Path:   fileId.String(),
			}
			q := url.Values{
				"type": {"replicate"},
				"ttl":  {params.Needle.Ttl.String()},
			}
			if params.Needle.LastModified > 0 {
				q.Set("ts", strconv.FormatUint(params.Needle.LastModified, 10))
			}
			if params.Needle.IsChunkedManifest() {
				q.Set("cm", "true")
			}
			u.RawQuery = q.Encode()

			pairMap := make(map[string]string)
			if params.Needle.HasPairs() {
				tmpMap := make(map[string]string)
				err := json.Unmarshal(params.Needle.Pairs, &tmpMap)
				if err != nil {
					stats.VolumeServerHandlerCounter.WithLabelValues(stats.ErrorUnmarshalPairs).Inc()
					glog.V(0).Infoln("Unmarshal pairs error:", err)
				}
				for k, v := range tmpMap {
					pairMap[needle.PairNamePrefix+k] = v
				}
			}
			bytesBuffer := buffer_pool.SyncPoolGetBuffer()
			defer buffer_pool.SyncPoolPutBuffer(bytesBuffer)

			// volume server do not know about encryption
			// TODO optimize here to compress data only once
			uploadOption := &operation.UploadOption{
				UploadUrl:         u.String(),
				Filename:          string(params.Needle.Name),
				Cipher:            false,
				IsInputCompressed: params.Needle.IsCompressed(),
				MimeType:          string(params.Needle.Mime),
				PairMap:           pairMap,
				Jwt:               jwt,
				Md5:               params.ContentMd5,
				BytesBuffer:       bytesBuffer,
			}

			uploader, err := operation.NewUploader()
			if err != nil {
				glog.Errorf("replication-UploadData, err:%v, url:%s", err, u.String())
				return err
			}
			_, err = uploader.UploadData(params.Needle.Data, uploadOption)
			if err != nil {
				glog.Errorf("replication-UploadData, err:%v, url:%s", err, u.String())
			}
			return err
		})
		stats.VolumeServerRequestHistogram.WithLabelValues(stats.WriteToReplicas).Observe(time.Since(start).Seconds())
		if err != nil {
			stats.VolumeServerHandlerCounter.WithLabelValues(stats.ErrorWriteToReplicas).Inc()
			err = fmt.Errorf("failed to write to replicas for volume %d: %v", params.VolumeId, err)
			glog.V(0).Infoln(err)
			return false, err
		}
	}
	return
}

func ReplicatedDelete(masterFn operation.GetMasterFn, grpcDialOption grpc.DialOption, store *storage.Store, params ReplicatedDeleteParams) (size types.Size, err error) {

	//check JWT
	jwt := params.Jwt

	var remoteLocations []operation.Location
	if params.Replicate {
		remoteLocations, err = GetWritableRemoteReplications(store, grpcDialOption, params.VolumeId, masterFn)
		if err != nil {
			glog.V(0).Infoln(err)
			return
		}
	}

	size, err = store.DeleteVolumeNeedle(params.VolumeId, params.Needle)
	if err != nil {
		glog.V(0).Infoln("delete error:", err)
		return
	}

	if len(remoteLocations) > 0 { //send to other replica locations
		if err = DistributedOperation(remoteLocations, func(location operation.Location) error {
			fileId := needle.NewFileIdFromNeedle(params.VolumeId, params.Needle)

			u := url.URL{
				Scheme: "http",
				Host:   location.Url,
				Path:   fileId.String(),
			}

			q := url.Values{
				"type": {"replicate"},
			}
			u.RawQuery = q.Encode()

			return util_http.Delete(u.String(), string(jwt))
		}); err != nil {
			size = 0
		}
	}
	return
}

type DistributedOperationResult map[string]error

func (dr DistributedOperationResult) Error() error {
	var errs []string
	for k, v := range dr {
		if v != nil {
			errs = append(errs, fmt.Sprintf("[%s]: %v", k, v))
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}

type RemoteResult struct {
	Host  string
	Error error
}

func DistributedOperation(locations []operation.Location, op func(location operation.Location) error) error {
	length := len(locations)
	results := make(chan RemoteResult)
	for _, location := range locations {
		go func(location operation.Location, results chan RemoteResult) {
			results <- RemoteResult{location.Url, op(location)}
		}(location, results)
	}
	ret := DistributedOperationResult(make(map[string]error))
	for i := 0; i < length; i++ {
		result := <-results
		ret[result.Host] = result.Error
	}

	return ret.Error()
}

func GetWritableRemoteReplications(s *storage.Store, grpcDialOption grpc.DialOption, volumeId needle.VolumeId, masterFn operation.GetMasterFn) (remoteLocations []operation.Location, err error) {

	v := s.GetVolume(volumeId)
	if v != nil && v.ReplicaPlacement.GetCopyCount() == 1 {
		return
	}

	// not on local store, or has replications
	lookupResult, lookupErr := operation.LookupVolumeId(masterFn, grpcDialOption, volumeId.String())
	if lookupErr == nil {
		selfUrl := util.JoinHostPort(s.Ip, s.Port)
		for _, location := range lookupResult.Locations {
			if location.Url != selfUrl {
				remoteLocations = append(remoteLocations, location)
			}
		}
	} else {
		err = fmt.Errorf("replicating lookup failed for %d: %v", volumeId, lookupErr)
		return
	}

	if v != nil {
		// has one local and has remote replications
		copyCount := v.ReplicaPlacement.GetCopyCount()
		if len(lookupResult.Locations) < copyCount {
			err = fmt.Errorf("replicating operations [%d] is less than volume %d replication copy count [%d]",
				len(lookupResult.Locations), volumeId, copyCount)
		}
	}

	return
}
