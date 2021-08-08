server:
	cd cmd && go run main.go -port 8899
proto:
	protoc --proto_path=proto proto/*.proto --go_out=:proto --go-grpc_out=:proto

.PHONY:server proto
