package utils

import (
	"fmt"
	"time"
	"unsafe"
)
func str2byte(str string) []byte{
	p := (*[2]uintptr)(unsafe.Pointer(&str))
	//var res [3]uintptr = {p[0],p[1],p[2]}
	res := [3]uintptr{p[0],p[1],p[1]}
	//res := {p[0],p[1],p[1]}
	return *(*[]byte)(unsafe.Pointer(&res))
}
func byte2str(b []byte) string{
	p := (*[3]uintptr)(unsafe.Pointer(&b))
	res:=[2]uintptr{p[0],p[1]}
	return *(*string)(unsafe.Pointer(&res))
}
func main(){
	str:="abcabcabc"
	//var p unsafe.Pointer
	//p := unsafe.Pointer(&str)
	p := (*[2]uintptr)((unsafe.Pointer(&str))  )
	//x := []byte(str)

	fmt.Println(&str)
	fmt.Println(uintptr(p[0]))
	//不是一个地址 证明发生了拷贝
	//fmt.Println(&str)
	//fmt.Println(&x[0])
	t1 := time.Now() // get current time
	//logic handlers

	for i:=0;i<1e9;i++{
		//temp:= str2byte(str)
		//str= byte2str(temp)

		//fmt.Println(temp)
		//fmt.Println(byte2str(temp))

		//慢的方法
		temp := []byte(str)
		str = string(temp)
	}
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
