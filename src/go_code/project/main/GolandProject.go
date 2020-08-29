package main

import (
	"fmt"
)

//Go语言的文件都要归属于同一个包
//Go语言达到了静态语言的编译速度,又有动态语言开发维护的效率
//有指针的概念 c+python
//天然的支持并发 [重要特点]
//可以不用;
//可以使用切片
//自动回收内存
//延时执行
//返回多个值

func enums() {
	const (
		v1 = "kall"
		v2 = iota
		v3
		v4
		v5
	)
	fmt.Printf("%s %d %d %d %d",v1, v2, v3, v4, v5)
}

//在这里没有object类型 但是有一个 空接口
//空接口可以传任何类型的数据
//interface{} << 空接口
func a(bb interface{}) {
	fmt.Println(fmt.Sprint(bb))
}


//%t是表示 bool类型的输出
//%d %f %c 分别是 int float unicode
//%b 二进制 %s string []byte
//%q byte

func main() {
	//var v0 string= "www"
	//a(v0)
	//var v1 bool = true
	//a(v1)
	//var v2 float32 = 1000.0
	//a(v2)
	//var v3 byte = 97
	//fmt.Printf("%q",v3)
	//a(v3)


	for i:=1;i<10;i++{
		n := 1

		LOOP : if n<=i{
			fmt.Printf("%d x %d = %d  ",i,n,i*n)
			n++
			goto LOOP
		}else{
			fmt.Printf("\n")
		}

	}
	var i int =0
	addEntity(&EntityList,&i,12,"wwww")
	addEntity(&EntityList,&i,12,"kall")
	addEntity(&EntityList,&i,12,"faaa")

	Loop:
	fmt.Printf("请输入key 来获取数据")
	fmt.Scanf("%d",&i)
	fmt.Printf("获取到的数据为: " + getName(EntityList,i))
	goto Loop
}
