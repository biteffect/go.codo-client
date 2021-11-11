package client

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/biteffect/go.codo_cash"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Client GM API client
type Client struct {
	apiUrl     *url.URL
	clientId   uuid.UUID
	httpClient *http.Client
	privateKey *rsa.PrivateKey
}

func (c *Client) call(method string, req interface{}, resp interface{}) error {
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
	}
	if req != nil {
		data["params"] = req
	}
	httpBody, err := json.Marshal(data)
	if err != nil {
		return err
	}

	httpRequest, err := http.NewRequest("POST", c.apiUrl.String(), bytes.NewReader(httpBody))
	if err != nil {
		return err
	}

	hash := sha256.New()
	hash.Write(httpBody)

	httpRequest.Header.Set("Content-Type", "application/json")
	httpRequest.Header.Set("Accept", "application/json")
	httpRequest.Header.Set("Content-Length", fmt.Sprint(len(httpBody)))

	httpRequest.Header.Set("Host", c.apiUrl.Host)
	httpRequest.Header.Set("Date", time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT"))
	httpRequest.Header.Set("Digest", fmt.Sprintf("SHA-256=%v", base64.StdEncoding.EncodeToString(hash.Sum(nil))))

	headers := []string{
		fmt.Sprintf(
			"(request-target): %s %s",
			strings.ToLower(httpRequest.Method), httpRequest.URL.RequestURI(),
		),
		fmt.Sprintf("host: %s", httpRequest.Header.Get("host")),
		fmt.Sprintf("date: %s", httpRequest.Header.Get("date")),
		fmt.Sprintf("digest: %s", httpRequest.Header.Get("digest")),
	}

	hash = sha256.New()
	hash.Write([]byte(strings.Join(headers, "\n")))

	signed, err := rsa.SignPKCS1v15(rand.Reader, c.privateKey, crypto.SHA256, hash.Sum(nil))
	if err != nil {
		return err
	}

	httpRequest.Header.Set("Authorization", fmt.Sprintf(
		`Signature keyId="%v",algorithm="rsa-sha256",headers="(request-target) host date digest",signature="%v"`,
		c.clientId.String(),
		base64.StdEncoding.EncodeToString(signed),
	))

	httpResp, err := c.httpClient.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResp.Body.Close()

	respBody, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}

	result := struct {
		JsonRpc string          `json:"jsonrpc"`
		Result  interface{}     `json:"result,omitempty"`
		ID      int             `json:"id"`
		Error   *codo.CodoError `json:"error,omitempty"`
	}{
		Result: resp,
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return err
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
