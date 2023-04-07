package cloudflare

import (
	"encoding/json"
	"time"
)

type ResponseSingle struct {
	Result   *Result    `json:"result,omitempty"`
	Success  bool       `json:"success,omitempty"`
	Errors   []*Error   `json:"errors,omitempty"`
	Messages []*Message `json:"messages,omitempty"`
}

type ResponseMany struct {
	Result   []*Result  `json:"result,omitempty"`
	Success  bool       `json:"success,omitempty"`
	Errors   []*Error   `json:"errors,omitempty"`
	Messages []*Message `json:"messages,omitempty"`
}

type Result struct {
	Id string `json:"id,omitempty"`
}

type Message struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Type    string `json:"type,omitempty"`
}

type Error struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

type DnsUpdate struct {
	Type    string        `json:"type,omitempty"`
	Name    string        `json:"name,omitempty"`
	Content string        `json:"content,omitempty"`
	TTL     time.Duration `json:"ttl,omitempty"`
	Comment string        `json:"comment,omitempty"`
}

func (u *DnsUpdate) MarshalJSON() ([]byte, error) {
	type Alias DnsUpdate
	return json.Marshal(&struct {
		TTL uint64 `json:"ttl,omitempty"`
		*Alias
	}{
		TTL:   uint64(u.TTL.Seconds()),
		Alias: (*Alias)(u),
	})
}
