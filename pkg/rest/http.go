package rest

import (
	"io"
	"net/http"
)

func DoHTTPRequest(client *http.Client, req *http.Request) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var restErr *Error
		restErr, err = ExtractRestErrorFromBody(body)
		if err != nil {
			return nil, err
		}

		return nil, restErr
	}

	return body, nil
}
