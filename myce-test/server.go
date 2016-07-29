package main
import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
    "strconv"
    "io"
    "encoding/json"
    "crypto/tls"
    "io/ioutil"
    "applist"
    "appdeploy"
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

        specSlice := []string {"521M 1C", "4G 1C", "16G 8C"}
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
        appList(w, r)

    }
}


func getApplist(w http.ResponseWriter, r *http.Request) {

    if r.Method == "GET" {
        t, err := template.ParseFiles("template/html/index.html")
        if err != nil {
            return 
        }
        err = t.Execute(w, nil)
    
    }


}

func appList(w http.ResponseWriter, r *http.Request) {
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

    
    podlist := make(applist.PodlistType, 20)

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
        dc[0] = "shijilulian"

        applist[i].Healthz.PodsAvailable = "all"
        applist[i].Name = podlist[i].Name
        applist[i].Label = rs.Items[i].Metadata.Labels
        applist[i].Datacenter = dc 
        applist[i].Replicas = 3
        applist[i].Worktime = rs.Items[i].Metadata.CreationTimeStamp
    }
		
	//json.NewEncoder(w).Encode(applist)
    result, _ := json.MarshalIndent(applist, "", "   ")
    fmt.Fprintln(w, string(result))	
}

func StaticServer(prefix string, staticDir string) {
    http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(staticDir))))
    return 
}

func main() {
    StaticServer("/template/", "./template")
    http.HandleFunc("/applist", appList)
    http.HandleFunc("/", sayhelloName)
    http.HandleFunc("/getapplist", getApplist)
    http.HandleFunc("/appdeploy", appDeploy)

    err := http.ListenAndServe(":10000", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}
