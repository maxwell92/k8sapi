package app

import (
	"k8s.io/kubernetes/pkg/apis/extensions"
)

// AppDeployment for Frontend fulfillment
type AppDeployment struct {
	Datacenters []AppDc               `json:"dataCenters"`
	Deployment  extensions.Deployment `json:"deployment"`
}

type AppDc struct {
	DcID float64 `json:"dcID,omitempty"`
}
