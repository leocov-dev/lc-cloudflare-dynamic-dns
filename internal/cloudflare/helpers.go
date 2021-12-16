package cloudflare

import (
	"encoding/json"
	"fmt"
	"io"
	"lc-cloudflare-dynamic-dns/config"
	"lc-cloudflare-dynamic-dns/internal/misc"
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

func parseBody(body io.ReadCloser) (response ResponseMany, err error) {
	if config.IsDebug {
		prettyPrintResponse(body)
	}
	defer body.Close()

	err = json.NewDecoder(body).Decode(&response)
	if err != nil {
		return response, fmt.Errorf("failed to parse api response: %s", err)
	}

	return response, nil
}

func validResponseCode(resp *http.Response, desiredCode int) error {
	if resp.StatusCode != desiredCode {
		if config.IsDebug {
			prettyPrintResponse(resp.Body)
		}
		return fmt.Errorf("status code error: %d", resp.StatusCode)
	}

	return nil
}

func prettyPrintResponse(body io.ReadCloser) {
	decoded, _ := decodeBody(body)
	misc.PrettyPrintInterface(decoded)
}

func setDefaultHeaders(authToken string, r *http.Request) {
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", config.UserAgent)
}
