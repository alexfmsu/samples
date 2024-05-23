#!/bin/bash
protoc ./proto/LotsService.proto --go_out=./proto --go-grpc_out=require_unimplemented_servers=false:./proto
