FROM ubuntu:22.04

# Add CA certs
RUN apt-get update && \
  apt-get install --yes --no-install-recommends \
    ca-certificates fio jed emacs-nox sudo time tmux s6 wget tini \
    openssh-server unzip man less groff bat hstr jq pip cargo && \
  pip install yq && \
  apt-get autoclean && \
  apt-get clean && \
  apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false && \
  rm -rf /var/lib/apt/lists/* && \
  rm -rf /var/cache/apt/* && \
  adduser --home /user --ingroup users --disabled-password --gecos "User" user && \
  mkdir -p /services/.s6-svscan && \
  chown -R user /services

ARG golangVersion="1.18.1"
ARG golangTarSHA="b3b815f47ababac13810fc6021eb73d65478e0b2db4b09d348eefad9581a2334"

ENV GOPATH /go

RUN \
  set -eux; \
  url="https://dl.google.com/go/go${golangVersion}.linux-amd64.tar.gz"; \
  wget -O go.tgz "$url" --progress=dot:giga; \
  echo "$golangTarSHA *go.tgz" | sha256sum -c -; \
  tar -C /usr/local -xzf go.tgz; \
  rm go.tgz; \
  mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

USER user:users

WORKDIR /user

RUN mkfifo stop && \
  mkdir -p services/.s6-svscan

ENTRYPOINT ["/usr/bin/tini", "--"]

#ENTRYPOINT [ "tail", "/user/stop" ]

ARG imageVersion="unknown"

LABEL imageVersion="${imageVersion}}"

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
