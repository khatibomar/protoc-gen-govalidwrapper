#!/bin/bash
rm -rf gen
mkdir -p gen/go
cd .. && go build -o example/protoc-gen-govalidwrapper . && cd example
protoc \
  --go_out=paths=source_relative,default_api_level=API_OPAQUE:gen/go \
  --govalidwrapper_out=gen/go \
  --plugin=protoc-gen-govalidwrapper=./protoc-gen-govalidwrapper \
  -I. \
  user.proto
