FROM golang:1.17

RUN adduser --home /user --ingroup users --disabled-password --gecos "User" user

# Add CA certs
RUN apt-get update && \
  apt-get install --yes ca-certificates fio jed emacs-nox sudo time tmux && \
  apt-get autoclean && \
  apt-get clean && \
  rm -rf /var/cache/apt/*

USER user:users

WORKDIR /user

RUN mkfifo stop

ENTRYPOINT [ "tail", "/user/stop" ]

COPY --chown=user:users clone-kopia.sh /user/

ARG imageVersion="unknown"

LABEL imageVersion="${imageVersion}}"
