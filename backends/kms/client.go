package kms

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type Client struct {
	svc *kms.KMS
}

func New() (*Client, error) {
	svc := kms.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})

  client := &Client{
		svc: svc,
	}

	return client, nil
}

func (c *Client) Decrypt(ciphertext string) (string, error) {

	data, err := base64.StdEncoding.DecodeString(ciphertext)

	decryptInput := &kms.DecryptInput{
    CiphertextBlob: data,
	}

	resp, err := c.svc.Decrypt(decryptInput)
	if err != nil {
		panic(err)
  }

	return string(resp.Plaintext), nil
}


