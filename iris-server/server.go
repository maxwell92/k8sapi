package main

import (
	//"fmt"
	color "github.com/iris-contrib/color"
	iris "github.com/kataras/iris"
	ia "irisapi"
)

const (
	banner = `
               *
              | |
              | |
           - -   - -
          | | | | | |
          | | | | | |
          -----------  
(o゜▽゜)o  is listening on :10000
`
	nv = `
{
  "list": [
    {
      "className": "fa-dashboard",
      "includeState": "main.dashboard",
      "state": "main.dashboard",
      "name": "Dashboard",
      "id": 1
    },
    {
      "item": [
        {
          "includeState": "main.appManageDeployment",
          "state": "main.appManageDeployment",
          "name": "发布",
          "id": 21
        },
        {
          "includeState": "main.appManageRollback",
          "state": "main.appManageRollback",
          "name": "回滚",
          "id": 22
        },
        {
          "includeState": "main.appManageRollup",
          "state": "main.appManageRollup",
          "name": "滚动升级",
          "id": 22
        },
        {
          "includeState": "main.appManageCancel",
          "state": "main.appManageCancel",
          "name": "撤销",
          "id": 22
        },
        {
          "includeState": "main.appManageHistory",
          "state": "main.appManageHistory",
          "name": "历史发布",
          "id": 22
        }
      ],
      "className": "fa-adn",
      "includeState": "main.appManage",
      "state": "main.appManage",
      "name": "应用管理",
      "id": 2
    },
    {
      "item": [
        {
          "includeState": "main.imageManageSearch",
          "state": "main.imageManageSearch",
          "name": "查找镜像",
          "id": 31
        },
        {
          "includeState": "main.imageManageDelete",
          "state": "main.imageManageDelete",
          "name": "删除镜像",
          "id": 32
        }
      ],
      "className": "fa-file-archive-o",
      "includeState": "main.imageManage",
      "state": "main.imageManage",
      "name": "镜像管理",
      "id": 3
    },
    {
      "className": "fa-cloud",
      "includeState": "main.rbdManage",
      "state": "main.rbdManage",
      "name": "云盘管理",
      "id": 4
    },
    {
      "item": [
        {
          "includeState": "main.extensionsService",
          "state": "main.extensionsService",
          "name": "创建服务",
          "id": 51
        },
        {
          "includeState": "main.extensionsEndpoint",
          "state": "main.extensionsEndpoint",
          "name": "创建访问点",
          "id": 52
        }
      ],
      "className": "fa-arrows",
      "includeState": "main.extensions",
      "state": "main.extensions",
      "name": "扩展功能",
      "id": 5
    },
    {
      "className": "fa-credit-card",
      "includeState": "main.costManage",
      "state": "main.costManage",
      "name": "计费&充值",
      "id": 6
    }
  ]
}
`
)

func navList(ctx *iris.Context) {
	//fmt.Fprintln(w, nv)
	ctx.Write(nv)
}
func init() {
	iris.StaticServe("./template", "/static")
	iris.Get("/navlist", navList)
	color.Cyan(banner)
}

func main() {
	iris.API("/deployment", ia.DeploymentAPI{})
	iris.Listen(":10000")
}
