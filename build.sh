#!/bin/bash

version=$1

if [ -z "$version" ]; then
    echo "version is needed"
    exit -1
fi

go build -ldflags "-X github.com/rajatjindal/kubectl-modify-secret/cmd.Version=$version" -o kubectl-modify-secret main.go