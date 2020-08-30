package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var arrays = [...]string{"水晶","达利","凤凰","玛瑙","山庄"}


func rollHouseName() string{
	return arrays[rand.Intn(len(arrays)-1)+1]
}

func houseMsg(n node){
	fmt.Println("==========================\n当前房子信息")
	if strings.Compare(n.name,notMoney)==0 {
		fmt.Println(notMoney)
		fmt.Printf("价值: %d\n",getMoeny(n.id))
	}else {
		fmt.Printf("名称： %s\n",n.name)
		fmt.Printf("楼房等级: %d\n",n.lv)
		fmt.Printf("金钱加成: %d\n",getAddMoney(n.id))
		fmt.Println("升级所需:暂未开放")
	}
}

