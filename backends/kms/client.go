package kms

import (
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
)

type Client struct {
	svc *kms.KMS
}

func New() (*Client, error) {

	s := session.New()
	region := *s.Config.Region

	metaSvc := ec2metadata.New(s)

	if len(region) < 1 {
		if metaSvc.Available() {
			r, err := metaSvc.Region()
			var _ = err
			region = r
		} else {
			region = "us-east-1"
		}
	}

	svc := kms.New(s, &aws.Config{Region: aws.String(region)})

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

	return string(resp.Plaintext), err
}

