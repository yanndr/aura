proto:
	protoc -I pb/ pb/aura.proto --go_out=plugins=grpc:pb

server:
	go run cmd/auraSvc/main.go

client:
	go run cmd/client/main.go
	