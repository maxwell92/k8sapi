package main

import (
	"appdeploy"
	"applist"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	//    "navlist"
)

func Get(url string) (body []byte, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Println(err)
		panic(err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func Post(url string, body io.Reader) (rtn []byte, err error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	//resp, err := client.Get(url)
	resp, err := client.Post(url, "application/json;charset=utf-8", body)
	if err != nil {
		log.Println(err)
		panic(err)
		return nil, err
	}

	defer resp.Body.Close()

	rtn, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rtn, nil
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//注意:如果没有调用ParseForm方法, 下面无法获取表单数据
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello myce")

}

func appDeploy(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//注意:如果没有调用ParseForm方法, 下面无法获取表单数据

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("template/html/appdeploy.html")
		//t.Execute(w, token) //used for session keep
		t.Execute(w, nil)

	} else {
		tempDeploy := new(appdeploy.BasicDeployment)
		tempApp := new(appdeploy.AppDeployment)
		fmt.Println("post assign")
		// tempApp assignment

		r.ParseForm()

		//注意:如果没有调用ParseForm方法, 下面无法获取表单数据
		fmt.Println(r.Form)
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])

		fmt.Println("appname:", r.Form["appname"])

		tempApp.Name = r.Form.Get("appname")
		tempApp.Namespace = r.Form.Get("namespace")

		dcSlice := []string{"Shijihulian", "Dianxin", "M6"}
		for _, v := range dcSlice {
			if v == r.Form.Get("datacenter") {
				//    tempApp.Datacenter = v
			}
		}

		tempApp.Image = r.Form.Get("image")
		tmp, _ := strconv.Atoi(r.Form.Get("replicas"))
		tempApp.Replicas = float64(tmp)

		specSlice := []string{"521M 1C", "4G 1C", "16G 8C"}
		for _, v := range specSlice {
			if v == r.Form.Get("spec") {
				//tempApp.Spec = v
			}
		}

		defaultlabels := make(map[string]string, 1)
		defaultlabels["appname"] = tempApp.Name

		fmt.Println(tempApp)

		// tempDeploy assignment
		tempDeploy.ApiVersion = "extensions/v1beta1"
		tempDeploy.Kind = "Deployment"

		tempDeploy.Metadata.Name = tempApp.Name
		tempDeploy.Spec.Replicas = tempApp.Replicas
		tempDeploy.Spec.Template.Metadata.Labels = defaultlabels
		tempDeploy.Spec.Template.Spec.Containers = make([]appdeploy.ContainersSTSDL, 1)
		tempDeploy.Spec.Template.Spec.Containers[0].Name = tempApp.Name
		tempDeploy.Spec.Template.Spec.Containers[0].Image = tempApp.Image

		var result []byte
		result, _ = json.MarshalIndent(tempDeploy, "", "    ")
		fmt.Fprintln(w, string(result))

		rep, err := Post("http://172.21.1.11:8080/apis/extensions/v1beta1/namespaces/default/deployments", strings.NewReader(string(result)))
		if err != nil {
			log.Println(err)
		}

		fmt.Println(string(rep))
		getApplist(w, r)

	}
}

func appList(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//t, err := template.ParseFiles("template/html/index.html")
		t, err := template.ParseFiles("template/index.html")
		if err != nil {
			return
		}
		err = t.Execute(w, nil)

	}

}

func getApplist(w http.ResponseWriter, r *http.Request) {
	var response []byte
	var err error

	response, err = Get("http://172.21.1.11:8080/api/v1/pods")
	//defer response.Body.Close()
	var rs applist.PodList
	err = json.Unmarshal(response, &rs)
	if err != nil {
		log.Println(err)
	}
	num := len(rs.Items)
	fmt.Println(num)

	podlist := make(applist.PodlistType, num)

	for i := 0; i < len(rs.Items); i++ {
		podlist[i].Name = rs.Items[i].Metadata.Name
		podlist[i].Ready = rs.Items[i].Status.ContainerStatuses[0].Ready
		podlist[i].Status = rs.Items[i].Status.Phase
		podlist[i].Restarts = int(rs.Items[i].Status.ContainerStatuses[0].RestartCount)
		podlist[i].StartTime = rs.Items[i].Status.StartTime
	}

	applist := make(applist.AppListType, num)

	for i := 0; i < len(rs.Items); i++ {
		var dc []string
		dc = make([]string, 1)
		dc[0] = "世纪互联"

		applist[i].Healthz.PodsAvailable = "全部"
		applist[i].Name = podlist[i].Name
		applist[i].Label = rs.Items[i].Metadata.Labels
		applist[i].Datacenter = dc
		applist[i].Replicas = 3
		//applist[i].Worktime = rs.Items[i].Metadata.CreationTimeStamp
		start := rs.Items[i].Status.ContainerStatuses[0].State.Running.StartedAt
		startTime, _ := time.Parse(time.RFC3339, start)
		now := time.Now()
		elapsedDuration := now.Sub(startTime)
		hour := elapsedDuration.Hours()
		min := elapsedDuration.Minutes()
		sec := elapsedDuration.Seconds()

		if hour > 24 {
			applist[i].Worktime = fmt.Sprintf("%d 天", int(hour/24))
		} else if hour < 1 {
			if min > 1 {
				applist[i].Worktime = fmt.Sprintf("%d 分钟", int(min))
			} else {
				applist[i].Worktime = fmt.Sprintf("%d 秒", int(sec))
			}

		} else {
			applist[i].Worktime = fmt.Sprintf("%d 小时", int(hour))
		}

	}

	//json.NewEncoder(w).Encode(applist)
	result, _ := json.MarshalIndent(applist, "", "   ")
	fmt.Fprintln(w, string(result))
}

func StaticServer(prefix string, staticDir string) {
	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(staticDir))))
	return
}

func navList(w http.ResponseWriter, r *http.Request) {
	/*
	   nl := new(navlist.Navlist)
	   nl.List = make([]navlist.ListType, 2)

	   nl.List[0].Id = 1
	   nl.List[0].Name = "Dashboard"
	   nl.List[0].State = "main.dashboard"
	   nl.List[0].IncludeState = "main.dashboard"
	   nl.List[0].ClassName = "fa-dashboard"

	   nl.List[1].Id = 2
	   nl.List[2].Name = "应用管理"
	   nl.List[2].State = "main.appManage"
	   nl.List[2].IncludeState = "main.appManage"
	   nl.List[2].ClassName = "fa-adn"
	   nl.List[2].Items = make([]navlist.ItemsType, 2)
	   nl.List[2].Items[0].Id = 22
	   nl.List[2].Items[0].Name = "发布"
	   nl.List[2].Items[0].State = "main.appManageDeployment"
	   nl.List[2].Items[0].IncludeState = "main.appManageDeployment"

	   nl.List[2].Items[1].Id = 23
	   nl.List[2].Items[1].Name = "回滚"
	   nl.List[2].Items[1].State = "main.appManageRollback"
	   nl.List[2].Items[1].IncludeState = "main.appManageRollback"

	   result, _ := json.MarshalIndent(nl,"", " ")
	   fmt.Fprintln(w, string(result))
	*/

	fmt.Fprintln(w, nv)
}

var nv = `
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

func main() {
	StaticServer("/static/", "./template")
	http.HandleFunc("/applist", appList)
	// http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/getapplist", getApplist)
	http.HandleFunc("/appdeploy", appDeploy)
	http.HandleFunc("/navlist", navList)

	err := http.ListenAndServe(":10000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
