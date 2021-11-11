package codo

import "fmt"

type CodoError struct {
	Code    int
	Message string
	Data    interface{} `json:",omitempty"`
}

func (ce *CodoError) Error() string {
	return fmt.Sprintf("[%v] %s", ce.Code, ce.Message)
}
