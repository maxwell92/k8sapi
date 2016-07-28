package main

type appDeployment struct {
   Name string `json: "name"`
   Namespace string `json: "namespace"`
   Datacenter []string `json: "datacenter"`
   Image string `json: "image"`
   Replicas float64 `json: "replicas"`
   Spec specType `json: "spec"`
   ConfigFile []string `json: "configFile"`
   MountPoints []mountPointsType `json: "mountPoints"`
   HealthCheck healthCheckType `json: "healthCheck"`
   Readiness readinessType `json: "readinessType"`
   PostStart postType   `json: "postStart"`
   PreStop preType `json: "preStart"`
   Env []envType `json: "env"`
   Label map[string] string `json: "label"`
   Command []string `json: "command"`
   Args []string `json: "args"`
}

type specType struct {
   Request map[string] string `json: "requests"`     
}

type mountPointsType struct {
   Name string `json: "name"`
   Readonly bool `json: "readOnly"`
   MountPath string `json: "mountPath"`
}


type healthCheckType struct {
   Exec execHealthType `json: "exec"`
   HttpGet httpHealthType `json: "httpGet"`
   TcpSocket tcpHealthType `json: "tcpSocket"`
   InitialDelaySeconds float64 `json: "initialDelaySeconds"`
   TimeoutSeconds float64 `json: "timeoutSeconds"`
   PeriodSeconds float64 `json: "periodSeconds"`
   SuccessThreshold float64 `json: "successThreshold"`
   FailureThreshold float64 `json: "failureThreshold"`
}

type execHealthType struct {
    Command []string `json: "command"` 
}

type httpHealthType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []headersHealthType `json: "httpHeaders"`    
}

type headersHealthType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type tcpHealthType struct {
    Port float64 `json: "port"`
}

type readinessType struct {
    Exec execReadType `json: "exec"`
    HttpGet httpReadType `json: "httpGet"`
    TcpSocket tcpReadType `json: "tcpSocket"`
    InitialDelaySeconds float64 `json: "initialDelaySeconds"`
    TimeoutSeconds float64 `json: "timeoutSeconds"`
    PeriodSeconds float64 `json: "periodSeconds"`
    SuccessThreshold float64 `json: "successThreshold"`
    FailureThreshold float64 `json: "failureThreshold"`
}

type execReadType struct {
    Command []string `json: "command"` 
}

type httpReadType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []headersRType `json: "httpHeaders"`    
}

type headersRType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type tcpReadType struct {
    Port float64 `json: "port"`
}

type postType struct {
    Exec execPostType `json: "exec"`
    HttpGet httpGetPostType `json: "httpGet"`
    TcpSocket tcpPostType `json: "tcpPSLCType"`
}

type execPostType struct {
    Command []string `json: "command"`
}

type httpGetPostType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []headersPostType `json: "httpHeaders"`    
}

type headersPostType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type tcpPostType struct {
    Port float64 `json: "port"`
}



type preType struct {
    Exec execPreType `json: "exec"`
    HttpGet httpGetPreType `json: "httpGet"`
    TcpSocket tcpPreType `json: "tcpPSLCType"`
}

type execPreType struct {
    Command []string `json: "command"`
}

type httpGetPreType struct {
    Path string `json: "path"`
    Port float64 `json: "port"`
    Host string `json: "host"`
    Scheme string `json: "scheme"`
    HttpHeaders []headersPreType `json: "httpHeaders"`    
}

type headersPreType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}

type tcpPreType struct {
    Port float64 `json: "port"`
}


type envType struct {
    Name string `json: "name"`
    Value string `json: "value"`
}


