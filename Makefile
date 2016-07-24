build:
	GO15VENDOREXPERIMENT=1 go build

deps:
		GO15VENDOREXPERIMENT=1 go get -u ./...

clean:
	rm rancher-server-secure-properties
