FROM ubuntu:22.04

# Add CA certs
RUN apt-get update && \
  apt-get install --yes --no-install-recommends \
    ca-certificates git fio jed emacs-nox sudo time tmux s6 wget curl tini \
    make openssh-server unzip man less groff bat hstr jq pip cargo && \
  pip install yq && \
  apt-get autoclean && \
  apt-get clean && \
  apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false && \
  rm -rf /var/lib/apt/lists/* && \
  rm -rf /var/cache/apt/* && \
  rm -rf /etc/ssh/ssh_host_*_key* && \
  adduser --home /user --ingroup users --disabled-password --gecos "User" user && \
  mkdir -p /service/.s6-svscan /services && \
  chown -R user /service /services

ARG golangVersion="1.18.3"
ARG golangTarSHA="956f8507b302ab0bb747613695cdae10af99bbd39a90cae522b7c0302cc27245"

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

# Add s6 services directories and shutdown mechanisms
RUN mkfifo stop && \
  mkdir -p service/.s6-svscan services/ssh && \
  echo 's6-svscanctl -q /user/service' > shutdown && \
  chmod 755 shutdown

COPY --chown=user:users ssh/ services/ssh/

ENTRYPOINT ["/usr/bin/s6-svscan", "/user/service"]

#ENTRYPOINT ["/usr/bin/tini", "--"]

#ENTRYPOINT [ "tail", "/user/stop" ]

ARG imageVersion="unknown"

LABEL imageVersion="${imageVersion}}"

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
