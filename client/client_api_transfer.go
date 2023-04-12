package client

import (
	"github.com/biteffect/go.codo_cash"
)

func (c *Client) TransferStatus(sid string) (*codo.Transfer, error) {
	return c.callTransferSimpleCommand("status", sid)
}

func (c *Client) callTransferSimpleCommand(cmd string, sid string) (*codo.Transfer, error) {
	transfer := &codo.Transfer{}
	if err := c.callSimpleCommand("transfer."+cmd, sid, transfer); err != nil {
		return nil, err
	}
	return transfer, nil
}

func (c *Client) callSimpleCommand(cmd string, sid string, res interface{}) error {
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return err
	}
	if err = c.call(cmd, req, res); err != nil {
		return err
	}
	return nil
}
