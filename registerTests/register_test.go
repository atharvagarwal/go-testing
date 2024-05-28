package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	//"regexp"
	"testing"
	//"github.com/ttacon/libphonenumber"
)

var userData = UserRegistrationRequest{
	Credential: "+917389998687",
	Password:   "Password1!",
	CircleName: "capcons",
}

// UserRegistrationRequest represents the JSON request structure
type UserRegistrationRequest struct {
	Credential string `json:"credential"`
	Password   string `json:"password"`
	CircleName string `json:"circleName"`
}

// (We don't need the registration handler here anymore)

func TestRegisterUserHandler(t *testing.T) {
	// Create a new HTTP request with JSON body
	requestBody := userData
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://capcons.com/go-auth/register", bytes.NewBuffer(requestBodyBytes))
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
	var expectedMessage = "User created successfully"

	// ... (rest of your test code)

	// Read the response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
		return
	}

	// Define a struct to unmarshal the JSON response
	var response struct {
		Message string `json:"message"`
	}

	// Unmarshal the response body (assuming it's JSON)
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v and the message is %v", err, string(bodyBytes))
		return
	}

	// Check if the expected message is present in the response
	if response.Message == expectedMessage {
		// Success: response contains the expected message
	} else {
		t.Errorf("API returned unexpected response: %v", string(bodyBytes))
	}

}

/*func validateUser(w http.ResponseWriter, r *http.Request) {
	user := userData
	// Validate credential (just checking length for example)
	num, err := libphonenumber.Parse(user.Credential, "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the phone number is valid
	if !libphonenumber.IsValidNumber(num) {
		http.Error(w, "Invalid Phone number", http.StatusBadRequest)
		return
	}

	// Validate password length
	if len(user.Password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	// Validate password complexity (at least one uppercase, digit, and special character)
	re := regexp.MustCompile(`[A-Z]`)
	digit := regexp.MustCompile(`[0-9]`)
	special := regexp.MustCompile(`[^a-zA-Z0-9]`)
	if !re.MatchString(user.Password) || !digit.MatchString(user.Password) || !special.MatchString(user.Password) {
		http.Error(w, "Password must contain at least one uppercase letter, digit, and special character", http.StatusBadRequest)
		return
	}
}*/
