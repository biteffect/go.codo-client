package codo_cash

import (
	gmfin "github.com/biteffect/go.gm-fin"
)

type TransferFee struct {
	FeeRules []TransferFeeRule `json:",omitempty"`
	Fund     TransferFeeItem   `json:",omitempty"`
	Invoice  TransferFeeItem   `json:",omitempty"`
	Fee      TransferFeeItem   `json:"TransferFee,omitempty"`
}

type TransferFeeItem struct {
	Amount   gmfin.Amount   `json:",omitempty"`
	Currency gmfin.Currency `json:",omitempty"`
}

type TransferFeeRule struct {
	Currency  gmfin.Currency `json:",omitempty"`
	MaxAmount gmfin.Amount   `json:",omitempty"`
	MinAmount gmfin.Amount   `json:",omitempty"`
	Scope     string         `json:",omitempty"`
	Type      string         `json:",omitempty"`
	Value     gmfin.Amount   `json:",omitempty"`
}
