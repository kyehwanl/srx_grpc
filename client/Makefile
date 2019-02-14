all:
	go build -gcflags '-N -l' srx_grpc_client.go
	go build -gcflags '-N -l' -buildmode=c-shared -o srx_grpc_client.so  srx_grpc_client.go
	gcc -g -O0 -o CImple_srx_grpc_client CImple_srx_grpc_client.c ./srx_grpc_client.so
