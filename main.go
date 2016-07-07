package main

import (
  "os"
  "strings"
  "fmt"

	"github.com/SungardAS/rancher-server-secure-properties/backends"
)

var (
  backendsConfig backends.Config
)

func main() {
	backendsConfig.Backend = "kms"
	decryptClient, err := backends.New(backendsConfig)
	if err != nil {
		return
	}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=",2)
		if strings.HasPrefix(pair[0],"CATTLE") {
			fmt.Println(pair[0])
			plaintext, err := decryptClient.Decrypt(pair[1]);
			if err != nil {
				panic(err)
			}

			fmt.Println(plaintext)

		}
	}

}
