all:
	protoc -I=. --go_out=plugins=grpc:. ./srxapi.proto
