package model

type DataObject struct {
	Name  string
	Type  DataObjectType
	Value string
}

const (
	DataObjectType_Int32 = iota
	DataObjectType_Int64
	DataObjectType_Float64
	DataObjectType_String
)

type DataObjectType int
