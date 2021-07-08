package codo_cash

import (
	"crypto/rsa"
	"crypto/tls"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"net/url"
	"time"
)

// NewClient returns Codo Cash API Client
func NewClient(u url.URL, id uuid.UUID, key *rsa.PrivateKey) (*Client, error) {
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
