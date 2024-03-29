package seaweedfs

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"icesfs/command/vars"
	"icesfs/errors"
	"icesfs/log"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type deleteObjectInfo struct {
	Size uint64 `json:"size"`
}

func (se *StorageEngine) DeleteObject(ctx context.Context, fid string) error {
	se.deletionQueue.EnQueue(fid)
	return nil
}

func (se *StorageEngine) loopProcessingDeletion() {
	var deleteCnt int
	for {
		deleteCnt = 0
		se.deletionQueue.Consume(func(fids []string) {
			for _, fid := range fids {
				ctx := context.Background()

				link, err := se.delLink(ctx, fid)
				if err != nil {
					log.Errorw("seaweedfs delete object: get link error", vars.ErrorKey, err.Error(), "fid", fid)
					continue
				}

				if link == 0 {
					err = se.deleteActualObject(ctx, fid)
					if err != nil {
						log.Errorw("seaweedfs delete object: delete actual object error", vars.ErrorKey, err.Error(), "fid", fid)
						continue
					}
				} else if link < 0 {
					_, err := se.kvStore.ClrNum(ctx, fidLinkKey(fid))
					if err != nil {
						log.Errorw("seaweedfs delete object: clear err fid link", vars.ErrorKey, err.Error())
					}
				}

				deleteCnt++
			}
		})
		if deleteCnt == 0 {
			time.Sleep(1234 * time.Millisecond)
		}
	}
}

func (se *StorageEngine) deleteActualObject(ctx context.Context, fullFid string) error {
	_, err := se.kvStore.ClrNum(ctx, fidLinkKey(fullFid))
	if err != nil {
		log.Errorw("seaweedfs delete actual object: clear fid link", vars.ErrorKey, err.Error())
		return err
	}

	volumeId, fid, err := se.parseFid(ctx, fullFid)
	if err != nil {
		log.Errorw("seaweedfs delete actual object: parse fid error", vars.ErrorKey, err.Error(), "fid", fid)
		return err
	}

	volumeIp, err := se.getVolumeHost(ctx, volumeId)
	if err != nil || volumeIp == "" {
		return err
	}

	req, err := http.NewRequest("DELETE", "http://"+volumeIp+"/"+strconv.FormatUint(volumeId, 10)+","+fid, nil)
	if err != nil {
		log.Errorw("seaweedfs delete actual object: new request delete error", vars.ErrorKey, err.Error(), "request url", "http://"+volumeIp+"/"+strconv.FormatUint(volumeId, 10)+","+fid)
		return errors.GetAPIErr(errors.ErrSeaweedFSVolume)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorw("seaweedfs delete actual object: do request delete error", vars.ErrorKey, err.Error(), "request url", "http://"+volumeIp+"/"+strconv.FormatUint(volumeId, 10)+","+fid, "request", req, "response", resp)
		return errors.GetAPIErr(errors.ErrSeaweedFSVolume)
	}
	if resp.StatusCode != http.StatusAccepted {
		log.Errorw("seaweedfs delete actual object: request error", "request url", "http://"+volumeIp+"/"+strconv.FormatUint(volumeId, 10)+","+fid, "http code", resp.StatusCode, "request", req, "response", resp)
		return errors.GetAPIErr(errors.ErrSeaweedFSVolume)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	httpBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorw("seaweedfs delete actual object: get http body error", vars.ErrorKey, err.Error(), "request url", "http://"+volumeIp+"/"+strconv.FormatUint(volumeId, 10)+","+fid, "response", resp)
		return errors.GetAPIErr(errors.ErrSeaweedFSVolume)
	}

	info := &deleteObjectInfo{}
	err = jsoniter.Unmarshal(httpBody, info)
	if err != nil {
		log.Errorw("seaweedfs delete actual object: get http body error", vars.ErrorKey, err.Error(), "request url", "http://"+volumeIp+"/"+strconv.FormatUint(volumeId, 10)+","+fid)
		return errors.GetAPIErr(errors.ErrSeaweedFSVolume)
	}

	return nil
}
