package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

func InitOssBucket() (*oss.Bucket, error){
	client, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		return nil, err
	}
	bucket, err := client.Bucket(MyBucket)
	if err != nil {
		return nil, err
	}
	return bucket, err
}

func PutObject(objectKey string, fileReader io.Reader) error {
	bucket, err := InitOssBucket()
	if err != nil {
		return err
	}
	if err = bucket.PutObject(objectKey, fileReader); err != nil {
		return err
	}
	return nil
}
