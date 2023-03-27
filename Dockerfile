FROM cjen1/reckon-mininet:latest as base

RUN apt-get update && apt-get install --no-install-recommends -yy -qq \
    build-essential \
    software-properties-common

RUN ln /usr/bin/ovs-testcontroller /usr/bin/controller

# Build dependencies
RUN apt-get update && apt-get install --no-install-recommends -yy -qq \
    autoconf \
    automake \
    libtool \
    curl \
    g++ \
    unzip

# Add stretch backports
RUN echo 'deb http://ftp.debian.org/debian stretch-backports main' | sudo tee /etc/apt/sources.list.d/stretch-backports.list
RUN apt-get update && apt-get install -yy -qq \
    openjdk-11-jdk

# Runtime dependencies
RUN pip3 install --upgrade wheel setuptools
ADD requirements.txt requirements.txt
RUN pip3 install -r requirements.txt
RUN apt-get update && apt-get install --no-install-recommends -yy -qq psmisc iptables

# Test dependencies
RUN apt-get update && apt-get install --no-install-recommends -yy -qq \
    tmux \
    screen \
    strace \
    linux-tools \
    tcpdump \
    lsof \
    vim \
    netcat \
    locales-all

# Add ocons dependencies
RUN bash -c "sh <(curl -L https://nixos.org/nix/install) --daemon"
RUN bash -c "echo 'experimental-features = nix-command flakes' > /etc/nix/nix.conf"

# Build ocons impl and client
# Required for silly reasons that ocaml has difficulty 
RUN git init . && git add .
RUN bash -l -c "nix build ./reckon/systems/ocons/ocons-src"
RUN bash -l -c "nix build ./reckon/systems/ocons/clients"

# Add reckon code
ADD . .
ENV PYTHONPATH="/root:${PYTHONPATH}"
ENV SHELL=/bin/bash

# Make directory for logs
RUN mkdir -p /results/logs

# Add built artefacts
COPY --from=etcd-image /reckon/systems/etcd reckon/systems/etcd
COPY --from=zk-image /reckon/systems/zookeeper/bins reckon/systems/zookeeper/bins
