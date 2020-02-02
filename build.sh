#!/usr/bin/env bash

VERSION=2.0
COMMIT_ID=bbbbb
BUILD_DATE=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
GO_VERSION=$(go version | awk '{print $3}')

go build -v -ldflags="-X 'github.com/kenchaaan/dnatctl/pkg/cmd/version.GitVersion=$VERSION' \
    -X 'github.com/kenchaaan/dnatctl/pkg/cmd/version.GitCommit=$COMMIT_ID' \
    -X 'github.com/kenchaaan/dnatctl/pkg/cmd/version.BuildDate=$BUILD_DATE' \
    -X 'github.com/kenchaaan/dnatctl/pkg/cmd/version.GoVersion=$GO_VERSION'"


#
#var (
#	gitVersion = "1.0"
#	gitCommit = "aaaaaa"
#	buildDate = "1970-01-01T00:00:00Z"
#	goVersion = "1.5"
#)