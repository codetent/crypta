FROM golang:1.21.5-bookworm

RUN GO111MODULE=on GOBIN=/usr/local/bin \
    go install github.com/bufbuild/buf/cmd/buf@v1.28.1
