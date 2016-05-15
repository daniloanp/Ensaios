#!/usr/bin/env bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

pushd $SCRIPT_DIR

pushd app
    pushd frontend
    pub build
    popd
    go build
    ./app
popd

