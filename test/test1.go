package main

//#include<stdio.h>
//void PrintInternalCall(int i) {
// printf(" testing - %d  calling from server driver module\n", i);
//}
import "C"
import "fmt"

func main() {
	fmt.Println("testing...")
	C.PrintInternalCall(12)
}
