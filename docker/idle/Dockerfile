FROM ubuntu:23.04

# Install required tools and signing keys for Google SDK
RUN \
  echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | \
    tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | \
    apt-key --keyring /usr/share/keyrings/cloud.google.gpg add - && \
  apt-get update && \
  apt-get install --yes --no-install-recommends \
    bat \
    build-essential \
    ca-certificates \
    cargo \
    curl \
    emacs-nox \
    file \
    fio \
    git \
    hstr \
    jed \
    jq \
    less \
    make \
    man \
    openssh-server \
    pip \
    pipx \
    procps \
    s6 \
    sudo \
    time \
    tmux \
    tini \
    unzip \
    wget  \
    && \
  apt-get autoclean && \
  apt-get clean && \
  apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false && \
  rm -rf /var/lib/apt/lists/* && \
  rm -rf /var/cache/apt/* && \
  rm -rf /etc/ssh/ssh_host_*_key* && \
  adduser --home /user --ingroup users --disabled-password --gecos "User" user && \
  mkdir -p /service/.s6-svscan /services && \
  chown -R user /service /services

# Install Google SDK
# RUN apt-get install google-cloud-cli -y

#pip install yq && \

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
