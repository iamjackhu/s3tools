package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	fmt.Println("Starting")

	endpoint := "10.0.1.88:9000"
	AccessKey := "ZhAmNIEBEfTNpWGR"
	SecretKey := "znN1kMoSlEh8NTy4h3skqakSrg45R2a8"

	httpClient := http.Client{}

	cred := credentials.NewStaticCredentials(AccessKey, SecretKey, "")
	disableSSL := true
	config := &aws.Config{
		Region:      aws.String("default"),
		Endpoint:    &endpoint,
		Credentials: cred,
		DisableSSL:  &disableSSL,
		HTTPClient:  &httpClient,
	}

	session, err := session.NewSession(config)
	if err != nil {
		fmt.Printf("can't create aws session: %v\n", err)
		os.Exit(1)
	}

	client := s3.New(session)

	request := client.NewRequest(&request.Operation{
		HTTPMethod: "PUT",
		HTTPPath:   "/testbk",
		BeforePresignFn: func(r *request.Request) error {
			return nil
		},
	}, nil, nil)

	if err := request.Send(); err != nil {
		fmt.Printf("can't send request: %v\n", err)
		os.Exit(2)
	}
}
