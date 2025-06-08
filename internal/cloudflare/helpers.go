package cloudflare

import (
	"encoding/json"
	"fmt"
	"io"
	"lc-cloudflare-dynamic-dns/config"
	"lc-cloudflare-dynamic-dns/internal/misc"
	"net/http"
)

func parseBody(body io.ReadCloser) (response ResponseMany, err error) {
	defer body.Close()

	err = json.NewDecoder(body).Decode(&response)

	if config.IsDebug {
		prettyPrintResponse(response)
	}

	if err != nil {
		return response, fmt.Errorf("failed to parse api response: %s", err)
	}

	return response, nil
}

func validResponseCode(resp *http.Response, desiredCode int) error {

	if resp.StatusCode != desiredCode {
		response, _ := parseBody(resp.Body)
		for _, item := range response.Errors {
			if item.Message == "A record with the same settings already exists." {
				return nil
			}
		}

		return fmt.Errorf("status code error: %d", resp.StatusCode)
	}

	return nil
}

func prettyPrintResponse(body interface{}) {
	misc.PrettyPrintInterface(body)
}

func setDefaultHeaders(authToken string, r *http.Request) {
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", config.UserAgent)
}
