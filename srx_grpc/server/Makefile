all:
	go build -gcflags '-N -l' srx_grpc_server.go
	go build -gcflags '-N -l' -buildmode=c-shared -o libsrx_grpc_server.so  srx_grpc_server.go
	gcc -g -O0 -o CImple_srx_grpc_server CImple_srx_grpc_server.c ./libsrx_grpc_server.so \
	-I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/srx_install/include/srx \
	-I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/ \
	-I/opt/project/srx_test1/srx/../_inst//include \
	-I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/../extras/local/include \
	-Wl,--unresolved-symbols=ignore-all
