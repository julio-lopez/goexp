FROM golang:1.26.2@sha256:b54cbf583d390341599d7bcbc062425c081105cc5ef6d170ced98ef9d047c716

# Install required tools and signing keys for Google SDK
RUN \
  apt-get update && \
  apt-get install --yes --no-install-recommends \
    bat \
    build-essential \
    ca-certificates \
    cargo \
    curl \
    emacs-nox \
    file \
    git \
    gnupg \
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
    util-linux-extra \
    wget \
  && \
  ln -s /usr/bin/batcat /usr/local/bin/bat && \
  apt-get autoclean && \
  apt-get clean && \
  apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false && \
  rm -rf /var/lib/apt/lists/* && \
  rm -rf /var/cache/apt/* && \
  rm -rf /etc/ssh/ssh_host_*_key*

RUN \
  echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | \
    tee -a /etc/apt/sources.list.d/google-cloud-sdk.list && \
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | \
    gpg --dearmor -o /usr/share/keyrings/cloud.google.gpg

# RUN \
#   apt-get update && \
#   apt-get install -y --no-install-recommends \
#     google-cloud-cli google-cloud-sdk-gke-gcloud-auth-plugin google-cloud-cli-gke-gcloud-auth-plugin && \
#   apt-get autoclean && \
#   apt-get clean && \
#   apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false && \
#   rm -rf /var/lib/apt/lists/* && \
#   rm -rf /var/cache/apt/* && \
#   rm -rf /etc/ssh/ssh_host_*_key*

RUN \
  adduser --home /home/user --ingroup users --disabled-password --gecos "User" user

  # && \
  # mkdir -p /service/.s6-svscan /services && \
  # chown -R user /service /services

USER user:users

WORKDIR /home/user

COPY --chown=user:users home/ /home/

# Add s6 services directories and shutdown mechanisms
RUN mkfifo stop && \
mkdir -p /home/s6/active-svc/.s6-svscan /home/user/go/

# COPY --chown=user:users ssh/ services/ssh/

ENV PATH=/usr/local/go/bin:/user/go/bin:$PATH

ENTRYPOINT ["/usr/bin/s6-svscan", "/home/s6/active-svc"]

#ENTRYPOINT ["/usr/bin/tini", "--"]

#ENTRYPOINT [ "tail", "/user/stop" ]

ARG imageVersion="unknown"

LABEL imageVersion="${imageVersion}}"
