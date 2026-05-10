package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	req, _ := http.NewRequest("GET", "ad", nil)
	req.Header.Add("Authorization", "")
	_, err := GetAPIKey(req.Header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatal("Failed to get failure on empy auth string")
	}
	req.Header.Set("Authorization", "123124")
	_, err = GetAPIKey(req.Header)
	if err == nil {
		t.Fatal("Failed to catch split failure")
	}
	req.Header.Set("Authorization", "ApiKey 12344")
	apiKey, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatal("Failed to get api key")
	}
	fmt.Print(apiKey)
}
