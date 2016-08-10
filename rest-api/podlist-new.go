package main

type PodList struct {
	Kind  string      `json: "kind"`
	Items []itemsType `json: "items"`
}

type itemsType struct {
	Metadata metadataItemsType `json: "metadata"`
	Spec     specItemsType     `json: "spec"`
}

type metadataItemsType struct {
	Name              string `json: "name"`
	Namespece         string `json: "namespace"`
	Uid               string `json: "uid"`
	CreationTimeStamp string `json: "creationTimestamp"`
	DeletionTimeStamp string `json: "deleteTimeStamp"`
	//  DeletionGracePeriodSeconds
	Labels map[string]string `json: "labels"`
	//  Annotations map[string]string `json: "annotations"`
}

type specItemsType struct {
	Volumes       []volumesSpecType    `json: "volumes"`
	Containers    []containersSpecType `json: "containers"`
	RestartPolicy string               `json: "restartPolicy"`
	//  TerminationGracePeriodSecond
	//  ActiveDeadlineSeconds
	DnsPolicy    string `json: "dnsPolicy"`
	NodeSelector string `json: "nodeSelector"`
	NodeName     string `json: "nodeName"`
}

type volumesSpecType struct {
	Name      string                   `json: "name"`
	HostPath  hostPathVolumesSpecType  `json: "hostPath"`
	GitRepo   gitRepoVolumesSpecType   `json: "gitRepo"`
	Secret    secretVolumesSpecType    `json: "secretName"`
	Rbd       rbdVolumesSpecType       `json: "rbd"`
	ConfigMap configMapVolumesSpecType `json: "configMap"`
}

type hostPathVolumesSpecType struct {
}

type gitRepoVolumesSpecType struct {
}

type secretVolumesSpecType struct {
}

type rbdVolumesSpecType struct {
}

type configMapVolumesSpecType struct {
}

type containersSpecType struct {
	Name            string                     `json: "name"`
	Image           string                     `json: "image"`
	Commands        []string                   `json: "command"`
	Args            []string                   `json: "args"`
	WorkingDir      string                     `json: "workingDir"`
	Ports           []portConSpecType          `json: "ports"`
	Env             []envConSpecType           `json: "env"`
	Resources       resourcesConSpecType       `json: "resources"`
	VolumeMounts    []volumeMountsConSpecType  `json: "volumeMounts"`
	LivenessProbe   lvProbeConSpecType         `json: "livenessProbe"`
	ReadnessProbe   rdProbeConSpecType         `json: "readnessProbe"`
	LifeCycle       lifeCycleConSpecType       `json: "lifecycle"`
	ImagePullPolicy imagePullPolicyConSpecType `json: "imagePullPolicy"`
}
