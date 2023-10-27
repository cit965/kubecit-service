package qiniucloud

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
	"kubecit-service/internal/pkg/provider/oss"
)

type ossQiniuCloud struct {
	Bucket    string
	AccessKey string
	SecretKey string
	Domain    string
	Region    string
}

func (o *ossQiniuCloud) UploadFile(reader io.Reader, filename string, size int64) (interface{}, error) {
	cfg := storage.Config{
		Region:        &storage.ZoneHuadongZheJiang2,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	putPolicy := storage.PutPolicy{
		Scope: o.Bucket,
	}
	mac := qbox.NewMac(o.AccessKey, o.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err := formUploader.Put(context.Background(), &ret, upToken, filename, reader, size, &putExtra)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return ret, nil
}

func (o *ossQiniuCloud) DownloadFile() {

}

func NewossQiniuCloud(bucket, ak, sk string) oss.Oss {
	return &ossQiniuCloud{
		Bucket:    bucket,
		AccessKey: ak,
		SecretKey: sk,
	}
}
