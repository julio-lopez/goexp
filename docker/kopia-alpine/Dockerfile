FROM golang:1.22 AS builder

# PWD: /go

WORKDIR /go/src/

RUN git clone https://github.com/kopia/kopia

WORKDIR /go/src/kopia

ARG kopiaCommit="master"

RUN git checkout "${kopiaCommit}"

# Install static executable, by default it ends up at /go/bin/kopia
RUN GO111MODULE=on CGO_ENABLED=0 GO_EXTLINK_ENABLED=0 go install \
  -ldflags="-s -w -X github.com/kopia/kopia/repo.BuildVersion=${kopiaCommit} -X github.com/kopia/kopia/repo.BuildInfo=$(date +%Y%m%d-%H%M%S)-$(git rev-parse --short=7 HEAD)" .

FROM alpine:3.20.1@sha256:b89d9c93e9ed3597455c90a0b88a8bbb5cb7188438f70953fede212a0c4394e0

WORKDIR /kopia

# Add CA certs
RUN apk add --no-cache --verbose ca-certificates && \
  rm -rf /var/cache/apk/* && \
  adduser -D kopia && addgroup kopia kopia && \
  chown kopia /kopia

USER kopia:kopia

ENTRYPOINT [ "/bin/kopia" ]

COPY --from=builder --chown=kopia:kopia /go/bin/kopia /bin/kopia

ARG imageVersion="unknown"
LABEL imageVersion="${imageVersion}"
LABEL kopiaCommit="${kopiaCommit}"
