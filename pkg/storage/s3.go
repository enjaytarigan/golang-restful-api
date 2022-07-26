package storage

import (
	"brodo-demo/service/uploader"
	"bytes"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getBucketName() string {
	return os.Getenv("AWS_BUCKET_NAME")
}

func getRegion() string {
	return os.Getenv("AWS_REGION")
}

type AwsUploader interface {
	PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)
}

type Storage struct {
	s3 AwsUploader
}

func NewStorage(awsUploader AwsUploader) uploader.Uploader {
	return &Storage{
		s3: awsUploader,
	}
}

func (storage *Storage) Upload(fileHeader *multipart.FileHeader, file multipart.File) (string, error) {
	region := getRegion()
	bucket := getBucketName()

	size := fileHeader.Size

	buffer := make([]byte, size)

	file.Read(buffer)

	defer file.Close()

	filename := strconv.Itoa(int(time.Now().Unix())) + fileHeader.Filename

	_, err := storage.s3.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		ACL:           aws.String("public-read"),
		Body:          bytes.NewReader(buffer),
		Key:           aws.String(filename),
		ContentLength: aws.Int64(int64(size)),
		ContentType:   aws.String(http.DetectContentType(buffer)),
	})

	if err != nil {
		return "", err
	}

	fileurl := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucket, region, filename)

	return fileurl, nil
}
