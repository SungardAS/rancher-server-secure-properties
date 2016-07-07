package backends

import (
	"fmt"
	"errors"

	"github.com/SungardAS/rancher-server-secure-properties/backends/kms"
)

type DecryptClient interface {
	Decrypt(ciphertext string) (string, error)
}

func New(config Config) (DecryptClient, error) {
	fmt.Printf("Hello, world.\n")

	switch config.Backend {
	case "kms":
		return kms.New()
	}

	return nil, errors.New("Invalid backend")
}
