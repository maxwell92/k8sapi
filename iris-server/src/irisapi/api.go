package irisapi

import "github.com/kataras/iris"

type DeploymentAPI struct {
	*iris.Context
}

func (da DeploymentAPI) Get() {
	da.Write("Get from /deployment")
}

func (da DeploymentAPI) Post() {
	da.Write("Post from /deployment")
}

func (da DeploymentAPI) GetById(id string) {
	da.Write("Get from /deployment/%s", id)
}

func (da DeploymentAPI) DeleteById(id string) {
	da.Write("Delete from /deployment/%s", id)
}
