all:
	go build -gcflags '-N -l' srx_grpc_client.go
	go build -gcflags '-N -l' -buildmode=c-shared -o libsrx_grpc_client.so  srx_grpc_client.go
	gcc -g -O0 -o CImple_srx_grpc_client CImple_srx_grpc_client.c ./libsrx_grpc_client.so \
		-I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/srx_install/include/srx \
		-I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/ \
		-I/opt/project/srx_test1/srx/../_inst//include \
		-I/opt/project/gobgp_test/gowork/src/srx_grpc/srx/src/../extras/local/include \
		-Wl,--unresolved-symbols=ignore-all -lpthread
