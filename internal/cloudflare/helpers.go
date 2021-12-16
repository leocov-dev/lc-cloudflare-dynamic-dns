package cloudflare

import (
	"encoding/json"
	"fmt"
	"io"
	"lc-cloudflare-dynamic-dns/config"
	"net/http"
)

func decodeBody(body io.ReadCloser) (response interface{}, err error) {
	defer body.Close()

	err = json.NewDecoder(body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func parseBody(body io.ReadCloser) (response Response, err error) {
	defer body.Close()

	err = json.NewDecoder(body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}

func setDefaultHeaders(r *http.Request) {
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.C.Cloudflare.ApiToken))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", config.UserAgent)
}
