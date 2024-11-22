package main

import (
	"crypto/tls"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// NewS3Client create aws s3 client witn config
func NewS3Client(config *aws.Config) (*s3.S3, error) {
	config.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	session, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}

	return s3.New(session), nil
}
