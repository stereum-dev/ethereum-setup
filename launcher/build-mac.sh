#!/bin/bash

RELEASE=$1
if [ -z "$RELEASE" ];then
  echo "Please specify base-installer release (with buildnumber) as first parameter"
  exit 8
fi

go get github.com/pkg/browser

mkdir -p build
cp -r *.go *.mod *.sum build
cd build
sed -i '' s/%%RELEASE%%/${RELEASE}/ launcher.go
go build
