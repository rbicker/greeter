#!/bin/sh

# GO
protoc --proto_path=api/proto/v1 --go_out=plugins=grpc:api/proto/v1 greet.proto

# JS + D.TS
#protoc --proto_path=api/proto/v1 --js_out=import_style=commonjs+dts:web-client/greet \
#  --grpc-web_out=import_style=commonjs+dts,mode=grpcwebtext:web-client/greet greet.proto

# TS
protoc --proto_path=api/proto/v1 --js_out=import_style=commonjs:web-client/greet \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:web-client/greet greet.proto
