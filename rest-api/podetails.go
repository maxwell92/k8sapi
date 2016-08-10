package main

type PodetailsType struct {
	Name       string `json:"Name"`
	Namespace  string `json:"Namespace"`
	Node       string `json:"Node"`
	Start_Time string `json:"Start Time"`
	//	Labels	string	`json:"Labels"`
	Status string `json:"Status"`
	IP     string `json:"IP"`
	//	Controllers	string	`json:"Controllers"`
	//	Containers	[]ContainerType	`json:"Containers"`
	//	Conditions	[]ConditionType	`json:"Conditions"`
}
