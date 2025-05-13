# httpmock

`httpmock` is a simple utility for mocking HTTP clients in Go. It provides a way to create mock HTTP clients with customizable responses, making it easy to test code that interacts with external APIs.

## Installation

```bash
go get github.com/divar-ir/go-http-mock
```

## Usage

Below is an example of how to use `httpmock` in a test suite with Go's native `testing` package.

### Example Test

```go
package yourpackage_test

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/divar-ir/go-http-mock/pkg/httpmock"
)

func TestAPIClient(t *testing.T) {
	// Create a mock HTTP client
	mockClient := httpmock.NewMockClient(200, `{"message": "success"}`)

	// Use the mock client in your code
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp, err := mockClient.Do(req)
	if err != nil {
		t.Fatalf("request failed: %v", err)
	}

	// Check the status code
	if resp.StatusCode != 200 {
		t.Errorf("expected status code 200, got %d", resp.StatusCode)
	}

	// Read and verify the response body
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}

	expectedBody := `{"message": "success"}`
	if strings.TrimSpace(string(body)) != expectedBody {
		t.Errorf("expected body %s, got %s", expectedBody, string(body))
	}
}
```
