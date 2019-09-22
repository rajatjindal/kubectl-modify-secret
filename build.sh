#!/bin/bash

version=$1

if [ -z "$version" ]; then
    echo "version is needed"
    exit -1
fi

mkdir -p _dist
env GOOS=windows GOARCH=amd64 go build -ldflags "-X github.com/rajatjindal/kubectl-modify-secret/cmd.Version=$version" -o _dist/kubectl-modify-secret-windows-amd64-$version.exe main.go
env GOOS=darwin  GOARCH=amd64 go build -ldflags "-X github.com/rajatjindal/kubectl-modify-secret/cmd.Version=$version" -o _dist/kubectl-modify-secret-darwin-amd64-$version      main.go
env GOOS=linux   GOARCH=amd64 go build -ldflags "-X github.com/rajatjindal/kubectl-modify-secret/cmd.Version=$version" -o _dist/kubectl-modify-secret-linux-amd64-$version       main.go