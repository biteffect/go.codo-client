package codo

import (
	gmfin "github.com/biteffect/go.gm-fin"
)

type Balance struct {
	Balance  string         `json:",omitempty"`
	Amount   gmfin.Amount   `json:",omitempty"`
	Currency gmfin.Currency `json:",omitempty"`
}
