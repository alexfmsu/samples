#!/bin/bash
protoc ./proto/OrdersService.proto --go_out=./proto --go-grpc_out=require_unimplemented_servers=false:./proto
