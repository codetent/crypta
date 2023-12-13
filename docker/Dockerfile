FROM golang:1.21.5-bookworm

RUN GO111MODULE=on GOBIN=/usr/local/bin \
    go install github.com/bufbuild/buf/cmd/buf@v1.28.1 && \
    go install github.com/goreleaser/goreleaser@v1.22.1 && \
    go install github.com/vektra/mockery/v2@v2.38.0 && \
    go install github.com/miniscruff/changie@v1.16.1 && \
    go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2

RUN echo "test"
