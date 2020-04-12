#!/usr/bin/env bash

if [ ! -f install.sh ]; then
    echo 'install must be run within its container folder' 1>&2
    exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

echo $CURDIR

gofmt -w src

echo `pwd`

go install main

export GOPATH="$OLDGOPATH"

echo 'finished'
