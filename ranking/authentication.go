package ranking

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Shakarang/Epirank/config"
	"github.com/Shakarang/Epirank/ranking/urls"
)

// Authentication connects current user to retrieve its token
func Authentication(authentication *config.Authentication) error {

	// Create Form URL-Encoded with credentials
	authenticationData := url.Values{}
	authenticationData.Set("login", authentication.Login)
	authenticationData.Add("password", authentication.Password)

	// Create POST request with required header and data
	request, err := http.NewRequest("POST", urls.EpitechIntranet+"/?format=json", bytes.NewBufferString(authenticationData.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(authenticationData.Encode())))

	if err != nil {
		return err
	}

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	for _, element := range resp.Cookies() {
		if element.Name == "PHPSESSID" {
			authentication.Token = element.Value
		}
	}

	if authentication.Token == "" {
		return fmt.Errorf("Token not received")
	}

	return nil
}
