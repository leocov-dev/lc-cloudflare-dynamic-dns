package cloudflare

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"lc-cloudflare-dynamic-dns/config"
	"lc-cloudflare-dynamic-dns/internal/misc"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Api struct {
	http    *http.Client
	apiRoot string
}

func NewApiClient() *Api {
	return &Api{
		http:    &http.Client{},
		apiRoot: "https://api.cloudflare.com/client/v4",
	}
}

func (c *Api) getUrl(path string, params map[string]string) *url.URL {
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

func (c *Api) VerifyAuthToken() error {
	req, err := http.NewRequest("GET", c.getUrl("/user/tokens/verify", nil).String(), nil)
	setDefaultHeaders(req)

	if err != nil {
		return err
	}

	resp, err := c.http.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		if config.IsDebug {
			body, _ := decodeBody(resp.Body)
			misc.PrettyPrintInterface(body)
		}
		return fmt.Errorf("status code error: %d", resp.StatusCode)
	}

	data, err := parseBody(resp.Body)

	if err != nil {
		return err
	}

	if len(data.Messages) > 0 && data.Messages[0].Code == 10000 {
		return nil
	}

	return errors.New("failed to verify auth token")
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
	req, err := http.NewRequest(
		http.MethodGet,
		c.getUrl(
			"zones",
			map[string]string{"name": name},
		).String(),
		nil,
	)
	if err != nil {
		return "", err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}

	data, err := parseBody(resp.Body)
	if err != nil {
		return "", err
	}

	return data.Result.Id, nil
}

// GET https://api.cloudflare.com/client/v4/zones/{{zone_id}}/dns_records?type=A
func (c *Api) getRecordId(zoneId string) (string, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		c.getUrl(
			fmt.Sprintf("zones/%s/dns_records", zoneId),
			map[string]string{"type": "A"},
		).String(),
		nil,
	)
	if err != nil {
		return "", err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return "", err
	}

	data, err := parseBody(resp.Body)
	if err != nil {
		return "", err
	}

	return data.Result.Id, nil
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

	req, err := http.NewRequest(
		http.MethodPut,
		c.getUrl(
			fmt.Sprintf("zones/%s/dns_records/%s", zoneId, recordId),
			nil,
		).String(),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}

	_, err = c.http.Do(req)
	if err != nil {
		return err
	}

	return nil
}
