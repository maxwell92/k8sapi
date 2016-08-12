package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type AbcController struct {
	*iris.Context
}

func (ac AbcController) Get() {
	id := ac.Param("id")
	fmt.Printf("Get from /abc/%s\n", id)
	ac.Write("Get from /abc/" + id)
}

/*
func (ac AbcController) Get() {
	fmt.Println("get from Get")
	ac.Write("get from /abc")
}
*/
func (ac AbcController) Post() {
	oid := ac.Param("oid")
	lid := ac.Param("lid")
	fmt.Printf("post from /abc/%s/lalala/%s\n", oid, lid)
	ac.Write("post from /abc/" + oid + "/lalala/" + lid)
}

func main() {
	abc := new(AbcController)
	//iris.API("/abc/:id", *abc)
	iris.API("/abc", *abc)
	iris.API("/abc/:oid/lalala/:lid", *abc)
	iris.Listen(":10000")
}
