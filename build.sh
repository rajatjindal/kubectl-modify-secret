#!/bin/bash

version=$1

if [ -z "$version" ]; then
    echo "version is needed"
    exit -1
fi

rm -rf _dist || true
mkdir -p _dist
# env GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/rajatjindal/kubectl-modify-secret/cmd.Version=$version" -o _dist/kubectl-modify-secret-windows-amd64-$version.exe main.go
env GOOS=darwin  GOARCH=amd64 go build -ldflags "-X github.com/rajatjindal/kubectl-modify-secret/cmd.Version=$version" -o _dist/darwin-amd64/kubectl-modify-secret      main.go
tar -cvzf _dist/darwin-amd64-$version.tar.gz _dist/darwin-amd64

env GOOS=linux   GOARCH=amd64 go build -ldflags "-X github.com/rajatjindal/kubectl-modify-secret/cmd.Version=$version" -o _dist/linux-amd64/kubectl-modify-secret       main.go
tar -cvzf _dist/linux-amd64-$version.tar.gz _dist/linux-amd64
