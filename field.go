package codo

import (
	"fmt"
	"strconv"
)

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

func (f *Field) Int() int {
	switch f.Value.(type) {
	case string:
		v, _ := strconv.Atoi(f.Value.(string))
		return v
	case float64:
		return int(f.Value.(float64))
	default:
	}
	return 0
}

func (f *Field) Float() float64 {
	switch f.Value.(type) {
	case string:
		v, _ := strconv.ParseFloat(f.Value.(string), 64)
		return v
	case float64:
		return f.Value.(float64)
	default:
	}
	return 0
}

func (f *Field) String() string {
	switch f.Value.(type) {
	case string:
		return f.Value.(string)
	case float64:
		return fmt.Sprintf("%g", f.Value.(float64))
	default:
		return fmt.Sprintf("%v", f.Value)
	}
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
