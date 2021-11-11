package codo

import (
	gmfin "github.com/biteffect/go.gm-fin"
	"time"
)

type Transfer struct {
	Id         string         `json:",omitempty"`
	CreatedAt  time.Time      `json:",omitempty"`
	Method     MethodType     `json:",omitempty" pg:"method,fk,notnull,type:payout_method"`
	Status     TransferStatus ``
	Amount     gmfin.Amount   ``
	Currency   gmfin.Currency `json:",omitempty"`
	InitiateAt *time.Time     `json:",omitempty"`
	AcceptAt   *time.Time     `json:",omitempty"`

	Fields []*Field `json:"Fields"`
}

func (t *Transfer) FieldByKey(key CodoFieldName) *Field {
	for _, f := range t.Fields {
		if f.Key == key {
			return f
		}
	}
	return nil
}
