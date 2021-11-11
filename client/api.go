package client

import (
	"crypto/rsa"
	"crypto/tls"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/url"
	"time"
)

// NewClient returns Codo Cash API Client
func New(u *url.URL, sid string, key *rsa.PrivateKey) (*Client, error) {
	id, err := uuid.FromString(sid)
	if err != nil {
		return nil, err
	}
	out := Client{
		apiUrl:     u,
		clientId:   id,
		privateKey: key,
		httpClient: &http.Client{
			Timeout: 120 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					Renegotiation:      tls.RenegotiateOnceAsClient,
					InsecureSkipVerify: true,
				},
			}},
	}
	return &out, nil
}
