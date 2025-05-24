package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"google.golang.org/genproto/googleapis/iam/credentials/v1"
	"google.golang.org/grpc/experimental/credentials"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			"your-access-key-id",
			"your-secret",
			"",
		),
	})
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "your-bucket-name"
}

func main() {

}