package client

import (
	"github.com/biteffect/go.codo_cash"
	uuid "github.com/satori/go.uuid"
)

func (c *Client) Profile() (*codo.Profile, error) {
	profile := &codo.Profile{}
	err := c.call("profile.info", nil, profile)
	if err == nil {
		return profile, nil
	}
	return nil, err
}

func (c *Client) Balances() (*codo.Profile, error) {
	profile := &codo.Profile{}
	err := c.call("profile.info", nil, profile)
	if err == nil {
		return profile, nil
	}
	return nil, err
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
