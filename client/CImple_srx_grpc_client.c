
#include <stdio.h>
#include "srx_grpc_client.h"
#include <stdlib.h>


int main ()
{
    printf(" Running C imple grpc client from C\n");

    char buff[10];
    buff[0] = 0xAB;
    buff[1] = 0xCD;
    buff[2] = 0xEF;


    GoSlice pdu = {(void*)buff, (GoInt)3, (GoInt)10};
    //GoSlice emp = {};

    int32_t result;
    result = Run(pdu);
    //Run(emp);

    printf(" validation result: %02x\n", result);

    return 0;
}

