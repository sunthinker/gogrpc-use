#!/bin/bash

protoc --go_out=../ --go-grpc_out=../ --experimental_allow_proto3_optional ./userinfo.proto
