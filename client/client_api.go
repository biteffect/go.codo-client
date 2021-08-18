package client

import (
	codo_cash "github.com/biteffect/go.codo-client"
	uuid "github.com/satori/go.uuid"
)

func (c *Client) Profile() (*codo_cash.Profile, error) {
	profile := &codo_cash.Profile{}
	err := c.call("profile.info", nil, profile)
	switch err.(type) {
	case *codo_cash.CodoError:
		return nil, err
	case error:
		return nil, err
	default:
	}
	return profile, nil
}

func (c *Client) Balances() (*codo_cash.Profile, error) {
	profile := &codo_cash.Profile{}
	err := c.call("profile.info", nil, profile)
	switch err.(type) {
	case *codo_cash.CodoError:
		return nil, err
	case error:
		return nil, err
	default:
	}
	return profile, nil
}

func prepareTransferIdRequest(sid string) (map[string]interface{}, error) {
	id, err := toUuid(sid)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"TransferId": id,
	}, nil
}

func toUuid(sid string) (uuid.UUID, error) {
	return uuid.FromString(sid)
}
