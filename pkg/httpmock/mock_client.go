package httpmock

import (
	"io"
	"net/http"
	"strings"
)

func NewMockClient(statusCode int, response string) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(func(_ *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: statusCode,
				Body:       io.NopCloser(strings.NewReader(response)),
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
			}, nil
		}),
	}
}
