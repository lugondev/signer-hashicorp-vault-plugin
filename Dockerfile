############################
# STEP 1 build executable plugin binary
############################
FROM golang:1.19-buster AS builder

ARG TARGETOS
ARG TARGETARCH

RUN apt-get update && \
	apt-get install --no-install-recommends -y \
	ca-certificates upx-ucl

WORKDIR /plugin

ENV GO111MODULE=on
COPY go.mod go.sum ./
##COPY LICENSE ./
RUN go mod tidy

COPY . .

#COPY ./vault/scripts/vault-init.sh .
#COPY ./build/bin/signer-hashicorp-vault-plugin .

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -a -v -o signer-hashicorp-vault-plugin
RUN upx signer-hashicorp-vault-plugin
RUN sha256sum -b signer-hashicorp-vault-plugin | cut -d' ' -f1 > SHA256SUM

############################
# STEP 2 build new vault image
############################
FROM library/vault:1.8.4

RUN apk add --no-cache \
    jq \
    curl

# Expose the plugin directory as a volume
VOLUME /vault/plugins

COPY --from=builder /plugin/signer-hashicorp-vault-plugin /vault/plugins/signer-hashicorp-vault-plugin
COPY --from=builder /plugin/vault/tls/* /vault/tls/
COPY --from=builder /plugin/vault/scripts/* /usr/local/bin/

RUN setcap cap_ipc_lock=+ep /vault/plugins/signer-hashicorp-vault-plugin

EXPOSE 8200
