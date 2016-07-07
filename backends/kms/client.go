package kms

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type Client struct {
	svc *kms.KMS
}

func New() (*Client, error) {
	svc := kms.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

  client := &Client{
		svc: svc,
	}

	return client, nil
}

func (c *Client) Decrypt(ciphertext string) (string, error) {
	fmt.Printf("Hello, KMS.\n")
	c.svc.Decrypt(nil)
	return "", nil
}


