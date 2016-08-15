package deployment

import (
	"github.com/kataras/iris"
	//hc "httpclient"
	//"log"
)

type DeploymentController struct {
	*iris.Context
}

// GET /api/v1/organizations/{oid}/deployments
func (dc *DeploymentController) Get() {
	oid := dc.Param("oid")
	dc.Write("Get from /api/v1/organizations/%s/deployments", oid)

}

// GET /api/v1/organizations/{oid}/deployments/{id}
func (dc *DeploymentController) GetById() {
	oid := dc.Param("oid")
	id := dc.Param("id")
	dc.Write("Get from /api/v1/organizations/%s/deployments/%s", oid, id)
}

// POST /api/v1/deployments
func (dc *DeploymentController) Post() {
	dc.Write("Post from /api/v1/deployments")
}

func (dc *DeploymentController) Encode() {

}

func (dc *DeploymentController) Decode() {

}
