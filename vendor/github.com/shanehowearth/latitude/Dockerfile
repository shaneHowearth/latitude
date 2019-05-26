ARG GO_VERSION=1.10.8

FROM golang:$GO_VERSION as builder

COPY . latitude$GOPATH/src/latitude
WORKDIR $GOPATH/src/latitude
ADD . $GOPATH/src/latitude
COPY Gopkg.toml Gopkg.lock ./
# Install dep
ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
# Fetch dependencies.
RUN dep ensure --vendor-only
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -i -v -o /go/bin/client latitude
