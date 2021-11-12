package client

import (
	"github.com/biteffect/go.codo_cash"
)

func (c *Client) TransferSearchByBearerCode(code string) ([]*codo.Transfer, error) {
	res := make([]*codo.Transfer, 0)
	if err := c.call("transfer.payout.find", map[string]interface{}{
		"PayoutCode": code,
	}, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) TransferHold(sid string) (*codo.Transfer, error) {
	return c.callTransferSimpleCommand("hold", sid)
}

func (c *Client) TransferUnHold(sid string) (*codo.Transfer, error) {
	return c.callTransferSimpleCommand("unhold", sid)
}

func (c *Client) TransferPayoutFee(sid string) (*codo.TransferFee, error) {
	fee := &codo.TransferFee{}
	if err := c.callSimpleCommand("transfer.payout.fee", sid, fee); err != nil {
		return nil, err
	}
	return fee, nil
}

func (c *Client) TransferPayoutFields(sid string) ([]codo.FieldSchema, error) {
	res := make([]codo.FieldSchema, 0)
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return nil, err
	}
	if err = c.call("transfer.payout.fields", req, &res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) TransferSetPayoutFields(sid string, fields []codo.Field) ([]codo.FieldState, error) {
	res := make([]codo.FieldState, 0)
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return nil, err
	}
	req["Fields"] = fields
	if err = c.call("transfer.payout.set_fields", req, &res); err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) TransferExecutePayout(sid string) (*codo.Transfer, error) {
	return c.callTransferSimpleCommand("payout.execute", sid)
}
