package pkg

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMakeAuthRequest(t *testing.T) {
	// Set up test environment variables
	os.Setenv("OPTII_CLIENT_ID", "test_client_id")
	os.Setenv("OPTII_CLIENT_SECRET", "test_client_secret")
	os.Setenv("OPTII_AUTH_URL", "http://test-auth-url.com")
	issuedAt := time.Now().UnixMilli()

	// Create a test server to mock the HTTP endpoint
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request method and URL
		fmt.Printf("Request: %+v\n", r)
		if r.Method != "POST" {
			t.Errorf("Unexpected request: %s %s", r.Method, r.URL.String())
		}

		// Verify the request body
		expectedBody := "client_id=test_client_id&client_secret=test_client_secret&grant_type=client_credentials&scope=openapi"
		bodyBytes, _ := io.ReadAll(r.Body)
		if string(bodyBytes) != expectedBody {
			t.Errorf("Unexpected request body: %s", string(bodyBytes))
		}

		// Return a mock response
		response := fmt.Sprintf(`{"access_token": "test_access_token", "expires_in": 3600, "token_type": "Bearer", "issued_at": "%d"}`, issuedAt)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Override the authURL with the test server URL
	os.Setenv("OPTII_AUTH_URL", server.URL)

	// Call the function under test
	auth, err := makeAuthRequest()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify the returned auth object
	expectedAuth := &Auth{
		AccessToken: "test_access_token",
		IssuedAt:    fmt.Sprintf("%d", issuedAt),
		ExpireIn:    3600,
		TokenType:   "Bearer",
	}
	assert.Equal(t, expectedAuth, auth)
}
