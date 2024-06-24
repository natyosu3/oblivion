package discord

import (
	"bytes"
	"encoding/json"
	"fmt"

	"net/http"
	"os"
)

const Oauth2_URL = "https://discord.com/api/oauth2/authorize?client_id=1209121185405866004&response_type=code&redirect_uri=https%3A%2F%2Foblivion.natyosu.com%2Fauth%2Fcallback&scope=identify+email"

const BASE_API_URL = "https://discordapp.com/api"
const USER_ENDPOINT = "/users/@me"
const GUILDS_ENDPOINT = "/users/@me/guilds"

var Client_Id string
var Client_Secret string

type UserInfoResponse struct {
	Id                     string      `json:"id"`
	Username               string      `json:"username"`
	Avatar                 string      `json:"avatar"`
	Discriminator          string      `json:"discriminator"`
	PublicFlags            int         `json:"public_flags"`
	Premium_type           int         `json:"premium_type"`
	Flags                  int         `json:"flags"`
	Banner                 string      `json:"banner"`
	Accent_color           string      `json:"accent_color"`
	Global_name            string      `json:"global_name"`
	Avatar_decoration_data interface{} `json:"avatar_decoration_data"`
	Banner_color           string      `json:"banner_color"`
	Mfa_enabled            bool        `json:"mfa_enabled"`
	Locale                 string      `json:"locale"`
	Email                  string      `json:"email"`
	Verified               bool        `json:"verified"`
}

type GuildsInfoResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func init() {
	load_env()
}

func load_env() {
	Client_Id = os.Getenv("CLIENT_ID")
	Client_Secret = os.Getenv("CLIENT_SECRET")
}

func Oauth2(payload []byte) (string, error) {
	resp, err := http.Post("https://discordapp.com/api/oauth2/token", "application/x-www-form-urlencoded", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("HTTP POST request failed:", err)
		return "", err
	}
	defer resp.Body.Close()

	var responseMap map[string]interface{}

	if err := json.NewDecoder(resp.Body).Decode(&responseMap); err != nil {
		return "", fmt.Errorf("Error decoding response body: %v", err)
	}

	accessToken, ok := responseMap["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("Access token not found in response")
	}

	return accessToken, nil
}
