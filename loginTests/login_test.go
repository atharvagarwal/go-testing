package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

var userData = LoginRequest{
	Credential: "+917389998681",
	Password:   "Atharv@12",
	CircleName: "capcons",
}

// LoginRequest represents the JSON request structure for login
type LoginRequest struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
	CircleName string `json:"circleName"`
}

func TestLoginHandler(t *testing.T) {
	// Create a new HTTP request with JSON body
	requestBody := userData
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://capcons.com/go-auth/login", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Use a real HTTP client to send the request to the actual API route
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("API returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body (optional)
	// You might need to adjust this based on your API's response structure
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
		return
	}

	// Define a struct to unmarshal the JSON response (assuming it includes a token)
	var response struct {
		Token string `json:"token"`
	}

	// Unmarshal the response body
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v and the message is %v", err, string(bodyBytes))
		return
	}

	t.Logf("Token: %s", response.Token)
	// Check if the token field exists and is not empty
	if response.Token == "" {
		t.Errorf("API response does not contain a token")
	}
}
