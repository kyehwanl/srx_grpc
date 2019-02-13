all:
	go build -gcflags '-N -l' srxapi_server.go
	go build -gcflags '-N -l' -buildmode=c-shared -o srxapi_server.so  srxapi_server.go
	gcc -g -O0 -o CImple_srxapi_server CImple_srxapi_server.c ./srxapi_server.so
