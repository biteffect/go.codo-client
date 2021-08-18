package codo

import uuid "github.com/satori/go.uuid"

type Profile struct {
	Account struct {
		Id    uuid.UUID
		Title string
	}
	Organization string `json:",omitempty"`
}
