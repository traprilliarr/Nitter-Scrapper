package lib

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var S3_CLIENT *s3.S3

func InitS3() {
	key := os.Getenv("S3BUCKET_KEY")
	secret := os.Getenv("S3BUCKET_SECRET")
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(key, secret, ""),
		Endpoint:         aws.String(os.Getenv("S3BUCKET_HOST")),
		Region:           aws.String(os.Getenv("S3BUCKET_REGION")),
		S3ForcePathStyle: aws.Bool(false), // // Configures to use subdomain/virtual calling format. Depending on your version, alternatively use o.UsePathStyle = false
	}
	S3_CLIENT = s3.New(session.New(s3Config))
}
