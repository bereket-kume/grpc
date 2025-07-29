#!/bin/bash

# Add Go bin to PATH
export PATH=$PATH:$(go env GOPATH)/bin

# Compile the proto file
protoc --go_out=. --go-grpc_out=. payment.proto

echo "Proto compilation completed!" 