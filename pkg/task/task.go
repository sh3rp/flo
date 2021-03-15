package task

import "github.com/sh3rp/flo/pkg/model"

type TaskExec interface {
	Name() string
	Run([]model.DataObject) ([]model.DataObject, error)
}
