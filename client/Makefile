all:
	go build -gcflags '-N -l' srxapi_client.go
	go build -gcflags '-N -l' -buildmode=c-shared -o srxapi_client.so  srxapi_client.go
	gcc -g -O0 -o CImple_srxapi_client CImple_srxapi_client.c ./srxapi_client.so
