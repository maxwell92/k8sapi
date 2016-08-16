package main

import (
	ac "controller/app"
	"github.com/kataras/iris"
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "", 0)

}

const (
	SERVER string = "http://172.21.1.11:8080"
)

func main() {
	appc := ac.NewAppController(SERVER)
	iris.Get("/api/v1/organization/:oid/deployments/:id", appc.Describe)
	iris.Get("/api/v1/organization/:oid/deployments", appc.List)
	iris.Post("/api/v1/deployments", appc.Create)
	iris.Get("/test", appc.Test)
	iris.Listen(":10000")
}
