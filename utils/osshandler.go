package utils

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func InitOssBucket() (b *oss.Bucket) {
	// 创建OSSClient实例。
	var accessKeyId = ""     // #TODO: 这里自己改自己的对应配置
	var accessKeySecret = "" // #TODO: 这里自己改自己的对应配置
	var endpoint = ""        // #TODO: 这里自己改自己的对应配置

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	// 获取存储空间。
	bucketName := "" // #TODO: 这里自己改自己的对应配置
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return bucket
}
