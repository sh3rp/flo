package task

import (
	"fmt"

	"github.com/sh3rp/flo/pkg/model"
)

func NewCliTask(name string, cliParams []model.DataObject) TaskExec {
	return cliTask{name, cliParams}
}

type cliTask struct {
	name      string
	cliParams []model.DataObject
}

func (ct cliTask) Name() string {
	return ct.name
}

func (ct cliTask) Run(data []model.DataObject) ([]model.DataObject, error) {
	fmt.Printf("Running cli command: %+v", ct.cliParams)
	return nil, nil
}
