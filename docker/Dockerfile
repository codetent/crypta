FROM debian:bookworm-slim

# Install basic tools
RUN apt-get update && \
    apt-get install -y \
        ca-certificates \
        curl \
        git \
        gnupg \
    && \
    rm -rf /var/lib/apt/lists/*

# Install nodejs
ARG NODE_MAJOR=20
RUN mkdir -p /etc/apt/keyrings && \
    (curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg) && \
    (echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | tee /etc/apt/sources.list.d/nodesource.list) && \
    apt-get update && \
    apt-get install nodejs -y && \
    rm -rf /var/lib/apt/lists/*

# Install eget
RUN curl https://zyedidia.github.io/eget.sh | sh && \
    mv eget /usr/local/bin/eget

# Install go
ARG GO_VERSION=1.21.5
RUN eget --to "/usr/local/lib/go" \
    https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz \
    -f go
ENV PATH=/usr/local/lib/go/bin:$PATH

# Install golangci-lint
ARG GOLANGCI_LINT_VERSION=1.55.2
RUN eget --to "/usr/local/bin" \
    golangci/golangci-lint \
    -t "v${GOLANGCI_LINT_VERSION}" \
    -a "golangci-lint-${GOLANGCI_LINT_VERSION}-linux-amd64.tar.gz"

# Install changie
ARG CHANGIE_VERSION=1.17.0
RUN eget --to "/usr/local/bin" \
    miniscruff/changie \
    -t "v${CHANGIE_VERSION}" \
    -a "changie_${CHANGIE_VERSION}_linux_amd64.tar.gz"

# Install goreleaser
ARG GORELEASER_VERSION=1.22.1
RUN eget --to "/usr/local/bin" \
    goreleaser/goreleaser \
    -t "v${GORELEASER_VERSION}" \
    -a "goreleaser_Linux_x86_64.tar.gz"

# Install mockery
ARG MOCKERY_VERSION=2.38.0
RUN eget --to "/usr/local/bin" \
    vektra/mockery \
    -t "v${MOCKERY_VERSION}" \
    -a "mockery_${MOCKERY_VERSION}_Linux_x86_64.tar.gz"

# Install buf
ARG BUF_VERSION=1.28.1
RUN eget --to "/usr/local/bin" \
    bufbuild/buf \
    -t "v${BUF_VERSION}" \
    -a "buf-Linux-x86_64.tar.gz"

# Install ginkgo
ARG GINKGO_VERSION=2.13.2
RUN GOBIN=/usr/local/bin \
    go install github.com/onsi/ginkgo/v2/ginkgo@v${GINKGO_VERSION} && \
    go clean -cache

# Install gh cli
ARG GH_VERSION=2.40.1
RUN eget --to "/usr/local/bin" \
    cli/cli \
    -t "v${GH_VERSION}" \
    -a "gh_${GH_VERSION}_linux_amd64.tar.gz"
