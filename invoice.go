package codo

import (
	"github.com/biteffect/go.gm-fin"
	"time"
)

type Invoice struct {
	Id        string         `json:",omitempty"`
	CreatedAt time.Time      `json:",omitempty"`
	Method    MethodType     `json:",omitempty" pg:"method,fk,notnull,type:payout_method"`
	Status    InvoiceStatus  `pg:"status,notnull,type:transfer_status"`
	Amount    gmfin.Amount   ``
	Currency  gmfin.Currency `json:",omitempty"`
	Fields    []*Field       `json:"Fields"`
}

func (i *Invoice) FieldByKey(key CodoFieldName) *Field {
	for _, f := range i.Fields {
		if f.Key == key {
			return f
		}
	}
	return nil
}
