package main


var EntityList Entitys = Entitys{data: map[int]node{}}

type Entitys struct {
	data map[int]node
}

type node struct {
	id int
	name string
	age int
}


func addEntity(entityList *Entitys,id *int,age int,name string)  {
	*id +=1
	addentityN(entityList, node{id: *id,age: age,name: name})
}

func addentityN(entityList *Entitys,n node)  {
	entityList.data[n.id] = n
}


//用来判断是否有值
func getName(entityList Entitys,id int) string{

	if _,ok := entityList.data[id];ok {
		return entityList.data[id].name
	}
	return "不存在"

}
