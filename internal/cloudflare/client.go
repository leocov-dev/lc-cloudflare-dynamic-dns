package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Api struct {
	http      *http.Client
	authToken string
}

func NewApiClient(authToken string) *Api {
	return &Api{
		http:      &http.Client{},
		authToken: authToken,
	}
}

func (c Api) getUrl(path string, params map[string]string) *url.URL {
	u := &url.URL{
		Scheme: "https",
		Host:   "api.cloudflare.com",
		Path:   fmt.Sprintf("/client/v4/%s", strings.TrimLeft(path, "/")),
	}

	q := u.Query()

	for k, v := range params {
		q.Set(k, v)
	}

	u.RawQuery = q.Encode()
	return u
}

func (c *Api) DoRequest(
	method string,
	path string,
	query map[string]string,
	body io.Reader,
) (*http.Response, error) {
	req, err := http.NewRequest(
		method,
		c.getUrl(path, query).String(),
		body,
	)
	if err != nil {
		return nil, err
	}
	setDefaultHeaders(c.authToken, req)

	return c.http.Do(req)
}

// GET https://api.cloudflare.com/client/v4/user/tokens/verify
func (c *Api) VerifyAuthToken() error {
	resp, err := c.DoRequest(
		http.MethodGet,
		"/user/tokens/verify",
		nil,
		nil,
	)
	if err != nil {
		return err
	}

	return validResponseCode(resp, http.StatusOK)
}

func (c *Api) DoUpdate(name, ip string, ttl int) error {
	zoneId, err := c.getZoneId(name)
	if err != nil {
		return err
	}

	recordId, err := c.getRecordId(zoneId)
	if err != nil {
		return err
	}

	return c.updateDns(name, ip, ttl, zoneId, recordId)
}

// GET https://api.cloudflare.com/client/v4/zones?name={{name}}
func (c *Api) getZoneId(name string) (string, error) {
	resp, err := c.DoRequest(
		http.MethodGet,
		"zones",
		map[string]string{"name": name},
		nil,
	)
	if err != nil {
		return "", err
	}

	if err = validResponseCode(resp, http.StatusOK); err != nil {
		return "", err
	}

	data, err := parseBody(resp.Body)
	if err != nil {
		return "", err
	}

	return data.Result[0].Id, nil
}

// GET https://api.cloudflare.com/client/v4/zones/{{zone_id}}/dns_records?type=A
func (c *Api) getRecordId(zoneId string) (string, error) {
	resp, err := c.DoRequest(
		http.MethodGet,
		fmt.Sprintf("zones/%s/dns_records", zoneId),
		map[string]string{"type": "A"},
		nil,
	)
	if err != nil {
		return "", err
	}

	if err = validResponseCode(resp, http.StatusOK); err != nil {
		return "", err
	}

	data, err := parseBody(resp.Body)
	if err != nil {
		return "", err
	}

	return data.Result[0].Id, nil
}

// PUT https://api.cloudflare.com/client/v4/zones/{{zone_id}}/dns_records/{{a_record_id}}
func (c *Api) updateDns(name, ip string, ttl int, zoneId, recordId string) error {
	update := &DnsUpdate{
		Type:    "A",
		Name:    name,
		Content: ip,
		TTL:     time.Duration(ttl) * time.Second,
	}

	body, err := json.Marshal(update)
	if err != nil {
		return err
	}

	resp, err := c.DoRequest(
		http.MethodPut,
		fmt.Sprintf("zones/%s/dns_records/%s", zoneId, recordId),
		nil,
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	if err = validResponseCode(resp, http.StatusOK); err != nil {
		return err
	}

	return nil
}
