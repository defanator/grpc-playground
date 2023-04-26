#!/usr/bin/make -f

default:
	@echo "try: make (proto|deps)"

hc/hc.pb.go hc/hc_grpc.pb.go: hc/hc.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hc/hc.proto

proto: hc/hc.pb.go hc/hc_grpc.pb.go

deps:
	go mod tidy

request:
	grpcurl -plaintext 127.0.0.1:50051 hc.HealthCheck/GetStatus

request-with-proto:
	grpcurl -plaintext -proto ./hc/hc.proto 127.0.0.1:50051 hc.HealthCheck/GetStatus
