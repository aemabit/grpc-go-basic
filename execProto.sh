#!/bin/sh
echo "what is the source of the proto file?"
read source
protoc --go_out=./pb --go-grpc_out=./pb $source
echo "protoc executed"