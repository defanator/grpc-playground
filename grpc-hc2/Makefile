#!/usr/bin/make -f

default:
	@echo "try: make (proto|deps)"

hc/hc.pb.go hc/hc_grpc.pb.go: hc/hc.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hc/hc.proto

proto: hc/hc.pb.go hc/hc_grpc.pb.go

deps:
	go mod tidy

run:
	go run hc_server/hc_server.go

request:
	grpcurl -plaintext -d '{"service": "foo"}' 127.0.0.1:50051 hc.Health/Check
	grpcurl -plaintext -d '{"service": "bar"}' 127.0.0.1:50051 hc.Health/Check
	grpcurl -plaintext -d '{"service": "baz"}' 127.0.0.1:50051 hc.Health/Check

request-with-proto:
	grpcurl -plaintext -proto ./hc/hc.proto -d '{"service": "foo"}' 127.0.0.1:50051 hc.Health/Check
	grpcurl -plaintext -proto ./hc/hc.proto -d '{"service": "bar"}' 127.0.0.1:50051 hc.Health/Check
	grpcurl -plaintext -proto ./hc/hc.proto -d '{"service": "baz"}' 127.0.0.1:50051 hc.Health/Check
