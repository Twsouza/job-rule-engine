package pkg

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type Auth struct {
	AccessToken string `json:"access_token"`
	ExpireIn    int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
}

var auth *Auth

// Authenticate is a function that authenticates the user.
// It checks if the authentication is valid and if not, makes an authentication request.
// It returns the authenticated user and any error that occurred during the authentication process.
func Authenticate() (*Auth, error) {
	var err error
	if !auth.isValid() {
		auth, err = makeAuthRequest()
		if err != nil {
			return nil, err
		}
	}

	return auth, nil
}

// makeAuthRequest sends an authentication request to the OPTII_AUTH_URL endpoint using the client credentials flow.
// It retrieves the client ID, client secret, and authentication URL from environment variables.
// The function returns an Auth struct containing the authentication response or an error if the request fails.
func makeAuthRequest() (*Auth, error) {
	clientID := os.Getenv("OPTII_CLIENT_ID")
	clientSecret := os.Getenv("OPTII_CLIENT_SECRET")
	authURL := os.Getenv("OPTII_AUTH_URL")

	data := url.Values{}
	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "openapi")

	request, err := http.NewRequest("POST", authURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	auth := &Auth{}
	if err := json.NewDecoder(resp.Body).Decode(auth); err != nil {
		return nil, err
	}

	return auth, nil
}

func (a *Auth) isValid() bool {
	if a == nil {
		return false
	}

	issuedAtMillis, err := strconv.ParseInt(a.IssuedAt, 10, 64)
	if err != nil {
		return false
	}

	// Convert issuedAt from milliseconds to time.Time
	issuedAtTime := time.UnixMilli(issuedAtMillis)

	// Calculate the expiration time
	expirationTime := issuedAtTime.Add(time.Duration(a.ExpireIn) * time.Second)

	// Compare the expiration time with the current time
	return time.Now().Before(expirationTime)
}
