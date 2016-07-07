FROM golang:1.6.2

ADD . /go/src/github.com/SungardAS/rancher-server-secure-properties

RUN mkdir -p /go/src/github.com/SungardAS/rancher-server-secure-properties
RUN go get github.com/SungardAS/rancher-server-secure-properties
RUN go install github.com/SungardAS/rancher-server-secure-properties
RUN mkdir -p /usr/share/cattle/war

ENTRYPOINT ["/go/src/github.com/SungardAS/rancher-server-secure-properties/scripts/entry"]
CMD ["/go/bin/rancher-server-secure-properties"]
