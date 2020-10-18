#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

TAG="rconway/requestlogger"

echo "Building image to tag: $TAG"
docker build -t "$TAG" .

if test "$1" = "push"
then
    echo "Pushing tag '$TAG' to DockerHub..."
    docker push "$TAG"
fi
