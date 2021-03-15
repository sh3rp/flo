package model

type Task struct {
	Id          string
	Name        string
	Type        TaskType
	TriggeredBy string
	Inputs      []DataObject
	Outputs     []DataObject
}

type TaskType int
