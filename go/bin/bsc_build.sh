#!/bin/bash -ex

echo "bsc_build.sh"

[ -d bsc ] || git clone https://github.com/Tenet-X/bsc.git; \
    cd bsc \
    && git checkout $BRANCH \
    && git pull \
    && make geth