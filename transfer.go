package codo

import (
	gmfin "github.com/biteffect/go.gm-fin"
	"time"
)

type Transfer struct {
	Id         string         `json:",omitempty"`
	CreatedAt  time.Time      `json:",omitempty"`
	Method     MethodType     `json:",omitempty" pg:"method,fk,notnull,type:payout_method"`
	Status     TransferStatus `pg:"status,notnull,type:transfer_status"`
	Amount     gmfin.Amount   `pg:",type:numeric(20,4),notnull,default:0"`
	Currency   gmfin.Currency `json:",omitempty" pg:",type:varchar(3),notnull,default:'UAH'"`
	InitiateAt *time.Time     `json:",omitempty"`
	AcceptAt   *time.Time     `json:",omitempty"`

	Fields []*Field `json:"Fields"`
}
