package main

import (
  "os"
  "strings"

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

	os.MkdirAll("/usr/share/cattle/war",755)
	f, ferr := os.Create("/usr/share/cattle/war/cattle-local.properties")
	var _ = ferr

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=",2)
		if strings.HasPrefix(pair[0],"ENC_CATTLE") {
			plaintext, err := decryptClient.Decrypt(pair[1]);
			if err != nil {
				panic(err)
			}

			propertyName := strings.Replace(strings.Replace(pair[0],"ENC_CATTLE_","",1),"_",".",-1)
			propertyName = strings.ToLower(propertyName)

			f.WriteString(propertyName)
			f.WriteString("=")
			f.WriteString(plaintext)
			f.WriteString("\n")

		}
	}

}
