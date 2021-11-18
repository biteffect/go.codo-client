package client

import (
	"errors"
	"fmt"
	"github.com/biteffect/go.codo_cash"
)

func (c *Client) TransferCreate(t *codo.Transfer) (*codo.Transfer, error) {
	if t == nil {
		return nil, errors.New("transfer is empty")
	}
	if t.Method != codo.MethodByCode {
		return nil, fmt.Errorf("only %s method now supported", codo.MethodByCode)
	}
	if t.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater then %v", 0)
	}
	res := &codo.Transfer{}
	if err := c.call("transfer.create", map[string]interface{}{
		"transfer": t,
	}, res); err != nil {
		return nil, err
	}
	return res, nil

}

func (c *Client) TransferAcceptFee(sid string) (*codo.TransferFee, error) {
	fee := &codo.TransferFee{}
	if err := c.callSimpleCommand("transfer.accept.fee", sid, fee); err != nil {
		return nil, err
	}
	return fee, nil
}

func (c *Client) TransferAcceptFields(sid string) ([]codo.FieldSchema, error) {
	res := make([]codo.FieldSchema, 0)
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return nil, err
	}
	if err = c.call("transfer.accept.fields", req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) TransferSetAcceptFields(sid string, fields []codo.Field) ([]codo.FieldState, error) {
	res := make([]codo.FieldState, 0)
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return nil, err
	}
	req["Fields"] = fields
	if err = c.call("transfer.accept.set_fields", req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) TransferExecuteAccept(sid string) (*codo.Transfer, error) {
	return c.callTransferSimpleCommand("accept.execute", sid)
}
