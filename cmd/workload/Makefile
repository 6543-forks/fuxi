
GOPATH:=$(shell go env GOPATH)



.PHONY: run
run:
	go run *.go  --registry etcd --server_address 0.0.0.0:8080 --registry_address fuxi.io:12379

.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/workload/workload.proto

.PHONY: build
build: proto
	go build -o workload-api main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t workload-api:latest
