#!/bin/bash
protoc ./proto/MathService.proto --go_out=./proto --go-grpc_out=require_unimplemented_servers=false:./proto
