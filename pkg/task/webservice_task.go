package task

import (
	"fmt"

	"github.com/sh3rp/flo/pkg/model"
)

func NewWebserviceTask(name string, cfg WebserviceTaskConfig) TaskExec {
	return webserviceTask{name, cfg}
}

type WebserviceTaskConfig struct {
	URL        string
	Method     string
	Parameters []model.DataObject
	Body       []model.DataObject
}

type webserviceTask struct {
	name   string
	config WebserviceTaskConfig
}

func (wst webserviceTask) Name() string {
	return wst.name
}

func (wst webserviceTask) Run(data []model.DataObject) ([]model.DataObject, error) {
	fmt.Printf("Hitting ws url: %s", wst.config.URL)
	return nil, nil
}
