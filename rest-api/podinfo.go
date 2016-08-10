package main

type PodinfoType struct {
	Name      string `json:"Name"`
	Ready     bool   `json:"Ready"`
	Status    string `json:"Status"`
	Restarts  int    `json:"Restarts"`
	StartTime string `json:"StartTime"`
}

type PodlistType []PodinfoType
