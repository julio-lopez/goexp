FROM debian:stable-slim AS builder

ARG golangVersion="1.17.8b7"
ARG tarSha256="4a1fa2c8d77309e1ef5bafe7e80e75c06e70c0ae1212d9f3d95485017155491d"

ADD "https://go-boringcrypto.storage.googleapis.com/go${golangVersion}.linux-amd64.tar.gz" go.tgz

RUN echo "${tarSha256}  go.tgz" | sha256sum -c -

RUN tar -xzf go.tgz

FROM debian:stable-slim

CMD ["bash"]

ENV GOPATH=/go

# Add CA certs and dev libs
RUN apt-get update && \
  apt-get install --yes --no-install-recommends \
    ca-certificates curl netbase wget gnupg dirmngr git mercurial \
    openssh-client subversion g++ gcc libc6-dev make pkg-config && \
  apt-get autoclean && \
  apt-get clean && \
  rm -rf /var/cache/apt/* && \
  adduser --home /user --ingroup users --disabled-password --gecos "User" user && \
  mkdir -p "$GOPATH/src" "$GOPATH/bin" && \
  chmod -R 777 "$GOPATH"

ENV PATH=/go/bin:/usr/local/go/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

COPY --from=builder /go/ /usr/local/go/

WORKDIR ${GOPATH}

LABEL org.golang.version="${golangVersion}"

ENV GOLANG_VERSION="${golangVersion}"

ARG imageVersion="unknown"

LABEL imageVersion="${imageVersion}}"
