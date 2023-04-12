package client

import (
	"github.com/biteffect/go.codo_cash"
)

func (c *Client) TransferCancel(sid string, message string) (*codo.Transfer, error) {
	req := map[string]interface{}{
		"TransferId": sid,
		"Message":    message,
	}
	res := &codo.Transfer{}
	if err := c.call("transfer.cancel", req, res); err != nil {
		return nil, err
	}
	return res, nil
}
