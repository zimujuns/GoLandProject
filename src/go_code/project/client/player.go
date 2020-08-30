package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main(){
	tcpAdd,_:=net.ResolveTCPAddr("tcp4","127.0.0.1:8888")
	conn,err := net.DialTCP("tcp",nil,tcpAdd)
	if err!=nil{
		fmt.Println(err.Error())
	}
	result,err := ioutil.ReadAll(conn)
	if err!=nil{
		fmt.Println(err.Error())
	}
	fmt.Println(string(result))

	//var a string
	//for{
	//	result,_ := ioutil.ReadAll(conn)
	//
	//	fmt.Scanf("%s",&a)
	//	conn.Write([]byte(a))
	//}
}


