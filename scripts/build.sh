#!/bin/sh
CurrentVersion=1.0.1

Project=futuq

Path=$Project"/version"
GitCommit=$(git rev-parse --short HEAD || echo unsupported)

go build -ldflags "-X $Path.Version=$CurrentVersion   \
  -X '$Path.BuildTime=`date "+%Y-%m-%d %H:%M:%S"`'    \
  -X '$Path.GoVersion=`go version`' -X $Path.GitCommit=$GitCommit"  \
  api/futuq.go

echo "build finish !!"
echo "Version:" $CurrentVersion
echo "Git commit:" $GitCommit
echo "Go version:" `go version`