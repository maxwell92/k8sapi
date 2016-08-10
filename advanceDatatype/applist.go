package applist 

type PodinfoType struct {
	Name 	string 	`json:"Name"`
	Ready 	bool 	`json:"Ready"`
	Status 	string 	`json:"Status"`
	Restarts	int	`json:"Restarts"`
	StartTime	string	`json:"StartTime"`
}

type PodlistType []PodinfoType



type AppInfoType struct {
    Healthz AppHealthzType `json: "appHealthz"`
    Name string `json: "appName"`
    Label map[string] string `json: "appLabel"`
    Datacenter []string `json: "appDatacenter"`
    Replicas float64 `json: "appReplicas"` 
    Worktime string `json: "appWorktime"`
}

type AppHealthzType struct {
    PodsAvailable string `json: "podsAvailable"`
}

type AppListType []AppInfoType
