#!/bin/bash -ex

echo "phala_build.sh"

[ -d phala-blockchain ] || git clone --depth 1 --recurse-submodules --shallow-submodules -j 8 https://github.com/Tenet-X/phala-blockchain.git; \
    cd phala-blockchain \
    && git pull \
    && git checkout $BRANCH \
    && git pull \
    && git submodule update --recursive \
    && cargo +$NIGHTLY build --$PROFILE
