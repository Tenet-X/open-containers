FROM ubuntu:20.04
LABEL descritption="Golang builder, how to create the binary in this env"
LABEL maintainer="binwu@google.com"

USER root

ARG GO_VER=1.16.4

ENV PROFILE=release
ARG DEBIAN_FE=noninteractive
ARG TZID=Asia/Shanghai

# install ubuntu packages
RUN apt update \
    && DEBIAN_FRONTEND=$DEBIAN_FE TZ=$TZID apt install -y \
    build-essential \
    git \
    clang \
    libclang-dev \
    pkg-config \
    libssl-dev \
    curl \
    && apt autoremove -y; apt clean -y

# install golang build env
ENV GO_HOME=/opt/go
ENV PATH=$GO_HOME/bin:$PATH 

RUN set -eux; \
    cd /opt \ 
    && curl -L https://golang.org/dl/go${GO_VER}.linux-amd64.tar.gz -o go.tar.gz \
    && tar xzf go.tar.gz \
    && go version

RUN mkdir /builder
VOLUME builder /builder
WORKDIR /builder
