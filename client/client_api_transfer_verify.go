package client

import (
	"github.com/biteffect/go.codo_cash"
)

func (c *Client) TransferExecuteVerify(sid string, field codo.CodoFieldName) (*codo.Verification, error) {
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return nil, err
	}
	req["FieldKey"] = field
	res := &codo.Verification{}
	if err := c.call("transfer.verify.field", req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) VerifyStatus(sid string) (*codo.Verification, error) {
	res := &codo.Verification{}
	if err := c.call("verify.status", map[string]interface{}{
		"VerificationId": sid,
	}, res); err != nil {
		return nil, err
	}
	return res, nil
}
