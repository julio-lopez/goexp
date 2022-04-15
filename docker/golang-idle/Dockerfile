ARG golangVersion="1.18"

FROM golang:${golangVersion}

# Add CA certs
RUN apt-get update && \
  apt-get install --yes ca-certificates \
    fio jed emacs-nox sudo time tmux s6 wget tini openssh-server unzip man \
    less groff bat hstr jq pip && \
  pip install yq && \
  apt-get autoclean && \
  apt-get clean && \
  rm -rf /var/cache/apt/* && \
  adduser --home /user --ingroup users --disabled-password --gecos "User" user

USER user:users

WORKDIR /user

RUN mkfifo stop

ENTRYPOINT ["/usr/bin/tini", "--"]

#ENTRYPOINT [ "tail", "/user/stop" ]

#COPY --chown=user:users clone-kopia.sh /user/

ARG imageVersion="unknown"

LABEL imageVersion="${imageVersion}}"
