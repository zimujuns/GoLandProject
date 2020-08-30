package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
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

/**
	一个类似大富翁的小游戏 但是玩家只有一个
	格子为30 每3秒增加一次金钱
	每一回合[摇骰子]增加一倍的房价[基础每个格子600]
	每个房子能为玩家增加每秒收入
 */
var rollNum int = 1
func mainv() {
	Init()
	player := EntityList.data[0]
	v1:=1
	i:=0
	for {
		fmt.Println("请输入您要进行的操作 1:购买当前地产 2:摇骰子")
		fmt.Scanf("%d",&i)
		switch i {
		case 1:
			msg(5)
			if strings.Compare(player.name ,notMoney)==0 {
				if money>=getMoeny(player.id){
					houseName := rollHouseName()
					fmt.Printf("您购买了当前房子,名称叫为%s\n",houseName)
					player.name = houseName
					money-= getMoeny(player.id)
					msg(3)
					true()
				}else{
					msg(2)
				}
			}else{
				msg(1)
			}
			break
		case 2:
			msg(4)
			v1 = roll()
			fmt.Printf("你摇出了%d点\n",v1)
			player = EntityList.data[player.id + v1]
			true()
			houseMsg(*player)
			break
		}
	}

}

var money int = 600

func true()  {
	theadMoney()
}

func theadMoney(){
	addMoney := getAddMoneySum()
	v1 := rollNum*3
	money+=10+addMoney+v1
	fmt.Printf("增加了金钱 当前金钱为%d\n房子资产额外增加 %d\n回合金额外增加 %d\n",money,addMoney,v1)
}

func roll() int{
	rollAction()
	i := int(rand.Intn(5)+1)
	return i
}

func rollAction(){
	rollNum++
}

func action(s string, player *node) {
	i,_ := strconv.Atoi(s)
	v1 := 1
	switch i {
	case 1:
		msg(5)
		if strings.Compare(player.name ,notMoney)==0 {
			if money>=getMoeny(player.id){
				houseName := rollHouseName()
				fmt.Printf("您购买了当前房子,名称叫为%s\n",houseName)
				player.name = houseName
				money-= getMoeny(player.id)
				msg(3)
				true()
			}else{
				msg(2)
			}
		}else{
			msg(1)
		}
		break
	case 2:
		msg(4)
		v1 = roll()
		fmt.Printf("你摇出了%d点\n",v1)
		player = EntityList.data[player.id + v1]
		true()
		houseMsg(*player)
		break
	}
}

func msg (i int){
	switch i {
	case 1:
		fmt.Println("您当前的地块已被购买!")
		break
	case 2:
		fmt.Println("您的金钱不足!")
		break
	case 3:
		fmt.Println("您已购买成功!")
		break
	case 4:
		fmt.Println("你选择了摇色子的操作")
		break
	case 5:
		fmt.Println("你选择了购买房子的操作")
		break
	default:
		fmt.Println("程序有误,请联系管理")
	}
}
