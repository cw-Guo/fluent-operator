#!/bin/bash


# AUTOMATICALLY GENERATED
# DO NOT EDIT THIS FILE DIRECTLY, USE /post_checkout.erb


set -e

HOST_ARCH=$(uname -m)

if [ x"${HOST_ARCH}" == x"aarch64" ]; then
    echo "Building arm64 image natively"
    exit
fi

# Enable cross-platform builds https://github.com/multiarch/qemu-user-static
docker run --rm --privileged multiarch/qemu-user-static --reset -p yes
