package discord

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUserInfo(token string) (*UserInfoResponse, error) {
	url := BASE_API_URL + USER_ENDPOINT
	userinfo := &UserInfoResponse{}
	if err := makeRequest("GET", url, token, userinfo); err != nil {
		return nil, err
	}
	return userinfo, nil
}

func makeRequest(method, url, token string, target interface{}) error {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return fmt.Errorf("Error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response: %v", err)
	}

	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("Error unmarshaling response: %v", err)
	}

	return nil
}
