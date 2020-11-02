#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

TAG="rconway/requestlogger"
VERSION="1.0"

echo "Building image to tag: $TAG"
docker build -t "${TAG}:${VERSION}" .
docker tag "${TAG}:${VERSION}" "${TAG}:latest"

if test "$1" = "push"
then
    for v in "${VERSION}" "latest"
    do
        image="${TAG}:${v}"
        echo "Pushing tag '${image}' to DockerHub..."
        docker push "${image}"
    done
fi
