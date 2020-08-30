package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)


func main()  {
	Init()
	fmt.Println("正在建立连接")
	post := ":8888"
	tcpAddr,err := net.ResolveTCPAddr("tcp4",post)
	checkError(err)
	listener,err := net.ListenTCP("tcp",tcpAddr)
	checkError(err)
	a := "请输入您要进行的操作 1:购买当前地产 2:摇骰子"
	p := EntityList.data[0]
	for{
		fmt.Printf("准备获取 监听端口 %s",post)
		conn,_ :=listener.Accept()
		_,err := conn.Write([]byte(a))

		if err!=nil {
			fmt.Println(err.Error())
		}
		result,err := ioutil.ReadAll(conn)
		if err!=nil {
			fmt.Println(err.Error())
		}
		action(string(result),p)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

}