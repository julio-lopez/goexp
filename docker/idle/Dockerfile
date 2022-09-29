FROM ubuntu:22.04

# Add CA certs
RUN apt-get update && \
  apt-get install --yes --no-install-recommends \
    ca-certificates fio jed emacs-nox sudo time tmux s6 wget curl tini \
    openssh-server unzip man less groff bat hstr jq pip cargo make git && \
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

USER user:users

WORKDIR /user

# Add s6 services directories and shutdown mechanisms
RUN mkfifo stop && \
  mkdir -p service/.s6-svscan services/ssh go/src go/bin && \
  echo 's6-svscanctl -q /user/service' > shutdown && \
  chmod 755 shutdown go

COPY --chown=user:users ssh/ services/ssh/

ENTRYPOINT ["/usr/bin/s6-svscan", "/user/service"]

#ENTRYPOINT ["/usr/bin/tini", "--"]

#ENTRYPOINT [ "tail", "/user/stop" ]

ARG imageVersion="unknown"

LABEL imageVersion="${imageVersion}}"

ENV PATH /usr/local/go/bin:/user/go/bin:$PATH
