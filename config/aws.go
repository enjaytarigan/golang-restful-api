package config

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getAccessKeyId() string {
	return os.Getenv("AWS_ACCESS_KEY_ID")
}

func getAccessKeySecret() string {
	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func getRegion() string {
	return os.Getenv("AWS_REGION")
}

func NewS3Client() *s3.S3 {
	config := &aws.Config{
		Region:      aws.String(getRegion()),
		Credentials: credentials.NewStaticCredentials(getAccessKeyId(), getAccessKeySecret(), ""),
	}

	session, err := session.NewSession(config)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return s3.New(session)
}
