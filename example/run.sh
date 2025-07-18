#!/bin/bash
rm -rf gen
mkdir -p gen/go
cd .. && go build -o example/protoc-gen-govalidwrapper . && cd example
protoc \
  --go_out=paths=source_relative,default_api_level=API_OPAQUE:gen/go \
  --go_opt=Mgovalid.proto=google.golang.org/protobuf/types/known/emptypb \
  --govalidwrapper_out=gen/go \
  --govalidwrapper_opt=Mgovalid.proto=google.golang.org/protobuf/types/known/emptypb \
  --plugin=protoc-gen-govalidwrapper=./protoc-gen-govalidwrapper \
  -I. \
  user.proto
