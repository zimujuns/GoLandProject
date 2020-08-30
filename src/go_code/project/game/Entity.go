package main

import "strings"

var EntityList Entitys = Entitys{data: map[int]*node{}}
var notMoney = "暂未建造"


type Entitys struct {
	data map[int]*node
}

type node struct {
	id int
	name string
	money int
	addMoney int
	lv int
}

func Init(){
	for i:=0;i<30;i++{
		addentityN(&EntityList,node{id: i,name: notMoney,money: 600* rollNum,addMoney: 300,lv: 1})
	}
}

func addEntity(entityList *Entitys,id *int,age,addMoney,lv int,name string)  {
	*id +=1
	addentityN(entityList, node{id: *id,money: age,name: name,addMoney: addMoney,lv: lv})
}

func addentityN(entityList *Entitys,n node)  {
	entityList.data[n.id] = &n
}

func getAddMoney(id int) int {
	return EntityList.data[id].addMoney * EntityList.data[id].lv
}

func getAddMoneySum() int{
	extral := 0
	for _,n:= range EntityList.data {
		if strings.Compare(n.name,notMoney)!=0 {
			extral+=n.addMoney*n.lv
		}
	}

	return extral

}


//用来判断是否有值
func getName(entityList Entitys,id int) string{

	if _,ok := entityList.data[id];ok {
		return entityList.data[id].name
	}
	return "不存在"
}

func getMoeny(id int) int{
	if _,ok := EntityList.data[id];ok {
		return EntityList.data[id].money * rollNum
	}
	return -1
}