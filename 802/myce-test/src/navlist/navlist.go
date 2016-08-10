package navlist

type Navlist struct {
	List []ListType `json:"list"`
}

type ListType struct {
	Id           float64     `json:"id"`
	Name         string      `json:"name"`
	State        string      `json:"state"`
	IncludeState string      `json:"includeState"`
	ClassName    string      `json:"className"`
	Items        []ItemsType `json:"items,omitempty"`
}

type ItemsType struct {
	Id           float64 `json:"id"`
	Name         string  `json:"name"`
	State        string  `json:"state"`
	IncludeState string  `json:"includeState"`
}
