#!/usr/bin/env bash
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

pushd $SCRIPT_DIR

pushd application
    pushd frontend
        pub build
    popd
    go build
    rm -rf dist/*
    mv application dist/ensaios
    cp -r frontend/build/web dist/web
    pushd ./dist
        ./ensaios
    popd
popd

