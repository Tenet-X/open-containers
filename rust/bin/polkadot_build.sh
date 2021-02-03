#!/bin/bash -ex

echo "polkadot_build.sh"

[ -d polkadot ] || git clone https://github.com/paritytech/polkadot.git; \
    cd polkadot \
    && git pull --all \
    && git checkout $BRANCH \
    && git pull \
    && cargo build --$PROFILE
