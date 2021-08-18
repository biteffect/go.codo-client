package codo

type FieldResolveStatus string

const (
	FldResolveAccept FieldResolveStatus = "accept"
	FldResolveSkip   FieldResolveStatus = "skip"
	FldResolveError  FieldResolveStatus = "error"
)

type Field struct {
	Key   CodoFieldName
	Value interface{}
}

type FieldState struct {
	Field
	Status FieldResolveStatus `json:",omitempty"`
	Error  string             `json:",omitempty"`
}

type FieldSchema struct {
	Key      CodoFieldName
	Name     string
	Type     string
	Required bool `json:",omitempty"`
}
