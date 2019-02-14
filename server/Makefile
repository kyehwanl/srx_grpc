all:
	go build -gcflags '-N -l' srx_grpc_server.go
	go build -gcflags '-N -l' -buildmode=c-shared -o srx_grpc_server.so  srx_grpc_server.go
	gcc -g -O0 -o CImple_srx_grpc_server CImple_srx_grpc_server.c ./srx_grpc_server.so -I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/srx_install/include/srx
