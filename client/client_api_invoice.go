package client

import (
	"errors"
	"fmt"
	"github.com/biteffect/go.codo_cash"
)

func (c *Client) InvoiceCreate(i *codo.Invoice) (*codo.Invoice, error) {
	if i == nil {
		return nil, errors.New("transfer is empty")
	}
	if i.Method != codo.MethodByCode {
		return nil, fmt.Errorf("only %s method now supported", codo.MethodByCode)
	}
	if i.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater then %v", 0)
	}
	res := &codo.Invoice{}
	if err := c.call("invoice.create", map[string]interface{}{
		"invoice": i,
	}, res); err != nil {
		return nil, err
	}
	return res, nil

}

func (c *Client) InvoiceStatus(sid string) (*codo.Invoice, error) {
	return c.callInvoiceSimpleCommand("status", sid)
}

func (c *Client) InvoiceCancel(sid string, message string) (*codo.Invoice, error) {
	req := map[string]interface{}{
		"TransferId": sid,
		"Message":    message,
	}
	res := &codo.Invoice{}
	if err := c.call("invoice.cancel", req, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (c *Client) callInvoiceSimpleCommand(cmd string, sid string) (*codo.Invoice, error) {
	invoice := &codo.Invoice{}
	if err := c.callSimpleCommand("invoice."+cmd, sid, invoice); err != nil {
		return nil, err
	}
	return invoice, nil
}
