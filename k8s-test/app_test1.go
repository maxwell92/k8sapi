package main 
import (
    "testing"
    "net/http"
    "net/http/httptest"

)
func Test_appDeploy1(t *testing.T) {
    var tempApp appDeployment
	tempApp.Name = "nginx-deployment-test"
	tempApp.Namespace = "default"

    var dc []string
    dc = make([]string, 5)
    dc[0] = "shijilulian"
    dc[1] = "dianxin"

	tempApp.Datacenter = dc 
	tempApp.Image = "nginx:1.7.9"
	tempApp.Replicas = 3

	spec := make(map[string]string, 2)
	spec["cpu"] = "2"
	spec["mem"] = "8G"

	tempApp.Spec.Request = spec

    configfile := make([]string, 1)
    configfile[0] = "/home/ubuntu/nginx.conf"
	tempApp.ConfigFile = configfile

    mountpoints := make([]string, 1)
    mountpoints[0] = "/usr/share/nginx"
    tempApp.MountPoints = make([]mountPointsType, 1)
	tempApp.MountPoints[0].MountPath = "/usr/share/nginx" 

	env := make(map[string]string, 2)
	env["username"] = "maxwell"
	env["password"] = "123456"

    tempApp.Env = make([]envType, 2)
	tempApp.Env[0].Name = "username"
    tempApp.Env[0].Value = "maxwell"
    tempApp.Env[1].Name = "password"
    tempApp.Env[1].Value = "123456" 

	labels := make(map[string]string, 1)
	labels["app"] = "nginx"

	tempApp.Label = labels


	tempDeploy := new(basicDeployment)
	tempDeploy.ApiVersion = "extensions/v1beta1"
	tempDeploy.Kind = "Deployment"
	tempDeploy.Metadata.Name = tempApp.Name
    tempDeploy.Metadata.Labels = tempApp.Label 
	tempDeploy.Spec.Replicas = tempApp.Replicas
    tempDeploy.Spec.Template.Metadata.Labels = tempApp.Label
    tempDeploy.Spec.Template.Spec.Containers = make([]containersSTSDL, 1)
	tempDeploy.Spec.Template.Spec.Containers[0].Name = tempApp.Name
    tempDeploy.Spec.Template.Spec.Containers[0].Image = tempApp.Image


    var result []byte
    result, _ = json.MarshalIndent(tempDeploy, "", "  ") 
    

    handler := func(w http.ResponseWriter, r *http.Request) {
        http.Error(w, "something failed", http.StatusInternalServerError)
    }


    req, err := http.NewRequest("POST", "http://172.21.1.11:8080/apis/extensions/v1beta1/namespaces/default/deployments", strings.NewReader(string(result)))

    if err != nil {
        t.Error("Didn't pass")
    }

    w := httptest.NewRecorder()
    handler(w, req)

    t.Error(w.Code," - ", w.Body.String())

}
