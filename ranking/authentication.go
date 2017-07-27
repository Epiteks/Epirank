package ranking

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Epiteks/Epirank/config"
	"github.com/Epiteks/Epirank/ranking/urls"
)

// Authentication connects current user to retrieve its token
func Authentication(authentication *config.Authentication) error {

	// Create Form URL-Encoded with credentials
	authenticationData := url.Values{}
	authenticationData.Set("login", authentication.Login)
	authenticationData.Add("password", authentication.Password)

	// Create POST request with required header and data
	request, err := http.NewRequest("POST", urls.EpitechIntranet+"/?format=json", strings.NewReader(authenticationData.Encode()))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	for _, element := range resp.Cookies() {
		if element.Name == "user" {
			authentication.Token = element.Value
		}
	}

	if authentication.Token == "" {
		return fmt.Errorf("Token not received")
	}

	return nil
}
