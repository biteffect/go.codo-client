package codo_cash

func (c *Client) ProfileInfo() (*Profile, error) {
	profile := &Profile{}
	err := c.call("profile.info", nil, profile)
	switch err.(type) {
	case *CodoError:
		return nil, err
	case error:
		return nil, err
	default:
	}
	return profile, nil
}
