#docker run --rm -v $(pwd):/app -w /app <image_name>:<tag> golangci-lint run -c strict_config.yaml --timeout 2m0s
FROM golang:1.13.1 as builder
RUN apt-get update && apt-get install -y \
    wget \
    ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
# Install golangci-lint
RUN wget https://install.goreleaser.com/github.com/golangci/golangci-lint.sh  && \
    chmod +x ./golangci-lint.sh && \
    ./golangci-lint.sh -b $GOPATH/bin && \
    golangci-lint linters
ADD relaxed_config.yaml /tmp/golangcilint_relaxed_config.yaml
ADD strict_config.yaml /tmp/golangcilint_strict_config.yaml
ENTRYPOINT ["golangci-lint"]
CMD ["run", "-c", "/tmp/golangcilint_relaxed_config.yaml", "--timeout", "2m0s"]
