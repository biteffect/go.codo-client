package client

import (
	"errors"
	"fmt"
	codo_cash "github.com/biteffect/go.codo-client"
)

func (c *Client) TransferCreate(t *codo_cash.Transfer) (*codo_cash.Transfer, error) {
	if t == nil {
		return nil, errors.New("transfer is empty")
	}
	if t.Method != codo_cash.MethodByCode {
		return nil, fmt.Errorf("only %s method now supported", codo_cash.MethodByCode)
	}
	if t.Amount <= 0 {
		return nil, fmt.Errorf("amount must be greater then %v", 0)
	}
	res := &codo_cash.Transfer{}
	err := c.call("transfer.create", map[string]interface{}{
		"transfer": t,
	}, res)
	switch err.(type) {
	case *codo_cash.CodoError:
		return nil, err
	case error:
		return nil, err
	default:
	}
	return res, nil

}

func (c *Client) TransferStatus(sid string) (*codo_cash.Transfer, error) {
	return c.callTransferSimpleCommand("status", sid)
}

func (c *Client) TransferCancel(sid string) (*codo_cash.Transfer, error) {
	return c.callTransferSimpleCommand("cancel", sid)
}

func (c *Client) TransferAcceptFee(sid string) (*codo_cash.TransferFee, error) {
	fee := &codo_cash.TransferFee{}
	if err := c.callSimpleCommand("transfer.accept.fee", sid, fee); err != nil {
		return nil, err
	}
	return fee, nil
}

func (c *Client) TransferAcceptFields(sid string) ([]codo_cash.FieldSchema, error) {
	res := make([]codo_cash.FieldSchema, 0)
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return nil, err
	}
	err = c.call("transfer.accept.fields", req, &res)
	switch err.(type) {
	case *codo_cash.CodoError:
		return nil, err
	case error:
		return nil, err
	default:
	}
	return res, nil
}

func (c *Client) TransferSetAcceptFields(sid string, fields []codo_cash.Field) ([]codo_cash.FieldState, error) {
	res := make([]codo_cash.FieldState, 0)
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return nil, err
	}
	req["Fields"] = fields
	err = c.call("transfer.accept.fields", req, &res)
	switch err.(type) {
	case *codo_cash.CodoError:
		return nil, err
	case error:
		return nil, err
	default:
	}
	return res, nil
}

func (c *Client) TransferExecuteAccept(sid string) (*codo_cash.Transfer, error) {
	return c.callTransferSimpleCommand("accept.execute", sid)
}

func (c *Client) TransferPayoutFee(sid string) (*codo_cash.TransferFee, error) {
	fee := &codo_cash.TransferFee{}
	if err := c.callSimpleCommand("transfer.payout.fee", sid, fee); err != nil {
		return nil, err
	}
	return fee, nil
}

func (c *Client) TransferPayoutFields(sid string) ([]codo_cash.FieldSchema, error) {
	res := make([]codo_cash.FieldSchema, 0)
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return nil, err
	}
	err = c.call("transfer.payout.fields", req, &res)
	switch err.(type) {
	case *codo_cash.CodoError:
		return nil, err
	case error:
		return nil, err
	default:
	}
	return res, nil
}

func (c *Client) TransferSetPayoutFields(sid string, fields []codo_cash.Field) ([]codo_cash.FieldState, error) {
	res := make([]codo_cash.FieldState, 0)
	req, err := prepareTransferIdRequest(sid)
	if err != nil {
		return nil, err
	}
	req["Fields"] = fields
	err = c.call("transfer.payout.fields", req, &res)
	switch err.(type) {
	case *codo_cash.CodoError:
		return nil, err
	case error:
		return nil, err
	default:
	}
	return res, nil
}

func (c *Client) TransferExecutePayout(sid string) (*codo_cash.Transfer, error) {
	return c.callTransferSimpleCommand("payout.execute", sid)
}

func (c *Client) TransferHold(sid string) (*codo_cash.Transfer, error) {
	return c.callTransferSimpleCommand("hold", sid)
}

func (c *Client) callTransferSimpleCommand(cmd string, sid string) (*codo_cash.Transfer, error) {
	transfer := &codo_cash.Transfer{}
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
	err = c.call(cmd, req, res)
	switch err.(type) {
	case *codo_cash.CodoError:
		return err
	case error:
		return err
	default:
	}
	return nil
}
