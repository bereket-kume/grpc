#!/bin/bash

# Add Go bin to PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Compile the proto file
protoc --go_out=. --go-grpc_out=. payment.proto

# Move generated files to proto directory
mkdir -p proto
mv github.com/recitelabs/grpc/proto/*.go proto/ 2>/dev/null || true
rm -rf github.com/ 2>/dev/null || true

echo "Proto compilation completed!" 