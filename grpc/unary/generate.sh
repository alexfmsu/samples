#!/bin/bash
protoc ./proto/LoginService.proto --go_out=./proto --go-grpc_out=require_unimplemented_servers=false:./proto
