package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
)

func operateBucket(client *s3.S3, bucketName, method string) error {
	request := client.NewRequest(&request.Operation{
		HTTPMethod: method,
		HTTPPath:   fmt.Sprintf("/%s", bucketName),
		BeforePresignFn: func(r *request.Request) error {
			return nil
		},
	}, nil, nil)
	return request.Send()
}

// CreateBucket create bucket with bucketName by aws s3 client
func CreateBucket(client *s3.S3, bucketName string) error {
	return operateBucket(client, bucketName, http.MethodPut)
}

// DeleteBucket delete bucket with bucketName by aws s3 client
func DeleteBucket(client *s3.S3, bucketName string) error {
	return operateBucket(client, bucketName, http.MethodDelete)
}
