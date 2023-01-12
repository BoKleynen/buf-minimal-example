FROM node:16-bullseye as development
ARG TARGETARCH
WORKDIR /workspace

# Install dependencies
RUN apt-get update
RUN apt-get install -y ca-certificates git openssh-client build-essential curl python3 jq zip postgresql-client

# Install golang for the correct architecture
ARG GO_VERSION=1.19
ARG PROTOC_VERSION=21.12
RUN set -eux; \
    ARCH="$(arch)"; \
    case "${ARCH}" in \
       amd64|x86_64) \
         BINARY_URL="https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz"; \
         ;; \
       aarch64|arm64) \
         BINARY_URL="https://golang.org/dl/go${GO_VERSION}.linux-arm64.tar.gz"; \
         ;; \
       *) \
         echo "Unsupported arch: ${ARCH}"; \
         exit 1; \
         ;; \
    esac; \
    wget ${BINARY_URL} -O "go${GO_VERSION}.tar.gz" && rm -rf /usr/local/go && tar -C /usr/local -xzf "go${GO_VERSION}.tar.gz";

RUN mkdir /go
ENV GOPATH="/go"
ENV PATH="/usr/local/go/bin:${GOPATH}/bin:${PATH}"


RUN set -eux; \
    ARCH="$(arch)"; \
    case "${ARCH}" in \
       amd64|x86_64) \
         BINARY_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"; \
         ;; \
       aarch64|arm64) \
         BINARY_URL="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-aarch_64.zip"; \
         ;; \
       *) \
         echo "Unsupported arch: ${ARCH}"; \
         exit 1; \
         ;; \
    esac; \
    curl -sfLo protoc.zip ${BINARY_URL} && \
    mkdir protoc && \
    unzip -q -d /protoc protoc.zip && \
    rm -r protoc protoc.zip
ENV PATH="/protoc/bin:${PATH}"

RUN BIN="/usr/local/bin" && \
    VERSION="1.11.0" && \
      curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" -o "${BIN}/buf" && \
      chmod +x "${BIN}/buf"
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 &&  \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0

RUN mkdir -p /.cache && chmod -R 777 /.cache
RUN chmod -R 777 /go
RUN echo "[safe]" >> /etc/gitconfig &&  \
    echo "    directory=*" >> /etc/gitconfig

ENTRYPOINT [ "/bin/bash", "-c" ]
