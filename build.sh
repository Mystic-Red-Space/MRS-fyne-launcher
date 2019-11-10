#!/usr/bin/env bash

export GOOS=linux
export GOARCH=amd64
go build -v -o $GOOS-$GOARCH-MRS-launcher
export GOARCH=386
go build -v -o $GOOS-$GOARCH-MRS-launcher
export GOOS=darwin
go build -v -o $GOOS-$GOARCH-MRS-launcher
export GOARCH=amd64
go build -v -o $GOOS-$GOARCH-MRS-launcher
