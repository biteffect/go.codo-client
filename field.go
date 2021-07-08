package codo_cash

type FieldResolveStatus string

const (
	FldResolveAccept FieldResolveStatus = "accept"
	FldResolveSkip   FieldResolveStatus = "skip"
	FldResolveError  FieldResolveStatus = "error"
)

type Field struct {
	Key    CodoFieldName
	Value  interface{}
	Status FieldResolveStatus `json:",omitempty"`
	Error  string             `json:",omitempty"`
}
