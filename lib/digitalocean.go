package lib

import (
	"log"
	"mime/multipart"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3

func init() {
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(appConfig.Digitalocean.ID, appConfig.Digitalocean.Secret, ""),
		Endpoint:    aws.String(appConfig.Digitalocean.Endpoint),
		Region:      aws.String(appConfig.Digitalocean.Region),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		log.Fatalf("ERROR: %v\n", err)
		return
	}

	s3Client = s3.New(newSession)
}

func S3UploadFile(userName string, filePath string, img multipart.File, contentType string) (*s3.PutObjectOutput, error) {
	object := s3.PutObjectInput{
		Bucket:      aws.String(GetBucketName()),
		Key:         aws.String(filePath),
		Body:        img,
		ACL:         aws.String("private"),
		ContentType: aws.String(contentType),
		Metadata: map[string]*string{
			"created-at":    aws.String(time.Now().String()),
			"uploader-name": aws.String(userName),
		},
	}
	return s3Client.PutObject(&object)
}

func S3DownloadFile(fileName string) (string, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(appConfig.Digitalocean.Bucket),
		Key:    aws.String(fileName),
	}

	res, _ := s3Client.GetObjectRequest(input)
	urlStr, err := res.Presign(time.Minute * 180)
	if err != nil {
		log.Printf("ERROR: %v\n", err)
		return "", err
	}

	return urlStr, nil
}

func S3DeleteFile(fileName string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(appConfig.Digitalocean.Bucket),
		Key:    aws.String(fileName),
	}

	_, err := s3Client.DeleteObject(input)
	if err != nil {
		return err
	}
	return nil
}
