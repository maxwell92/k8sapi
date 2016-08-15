package main

import (
	"fmt"
	"github.com/kataras/iris"
)

type AbcController struct {
	*iris.Context
}

/*
func (ac AbcController) GetBy() {
	id := ac.Param("id")
	fmt.Printf("Get from /abc/%s\n", id)
	ac.Write("Get from /abc/" + id)
}
*/
/*
func (ac AbcController) GetBy(id string) {
	fmt.Printf("Get from /abc/%s\n", id)
	ac.Write("Get from /abc/" + id)
}
*/
func (ac AbcController) Get() {
	od := ac.Param("od")
	ld := ac.Param("ld")
	fmt.Printf("get from /abc/%s/lalala/%s\n", od, ld)

	//fmt.Println("get from Get")
	//ac.Write("get from /abc")
}

/*
func (ac AbcController) GetBy() {
	od := ac.Param("od")
	fmt.Printf("get from /abc/%s\n", od)
}
*/
func (ac AbcController) Post() {
	od := ac.Param("od")
	//ld := ac.Param("ld")
	fmt.Printf("post from /abc/%s/lalala\n", od)

}
func main() {
	abc := new(AbcController)
	//iris.API("/abc", *abc)
	iris.API("/abc/:od/lalala/:ld", *abc)
	iris.API("/abc/:od/lalala", *abc)
	iris.API("/abc", *abc)
	iris.Listen(":10000")
}
