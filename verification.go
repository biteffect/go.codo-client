package codo

import "time"

type Verification struct {
	Id       string     `json:",omitempty"`
	Method   string     `json:",omitempty"`
	Status   string     ``
	ExpireAt *time.Time `json:",omitempty"`
}
