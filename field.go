package codo

type FieldType string
type FieldResolveStatus string

const (
	FldTypeMoney        FieldType = "money"
	FldTypeNumber       FieldType = "number"
	FldTypePhone        FieldType = "phone"
	FldTypeEmail        FieldType = "email"
	FldTypeText         FieldType = "text"
	FldTypeBool         FieldType = "boolean"
	FldTypeDate         FieldType = "date"
	FldTypeTaxonomy     FieldType = "taxonomy"
	FldTypeVerification FieldType = "verification"

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
	Type     FieldType
	Name     string
	Required bool `json:",omitempty"`
}

func FieldTypes() []FieldType {
	return []FieldType{
		FldTypeTaxonomy,
		FldTypeMoney,
		FldTypeNumber,
		FldTypePhone,
		FldTypeText,
		FldTypeDate,
		FldTypeBool,
	}
}
