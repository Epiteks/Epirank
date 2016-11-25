package config

import (
	"encoding/json"
	"os"

	"github.com/Shakarang/Epirank/models"
	log "github.com/Sirupsen/logrus"
)

// DatabasePath is database path...
const DatabasePath = "./students.db"

const tek1 = "tek1"
const tek2 = "tek2"
const tek3 = "tek3"
const tek4 = "tek4"
const tek5 = "tek5"

// Cities list
var Cities = []models.City{
	{
		Name:       "Strasbourg",
		ID:         "STG",
		Promotions: []string{tek3},
	},
	// {
	// 	Name:       "Nice",
	// 	ID:         "NCE",
	// 	Promotions: []string{tek1, tek2, tek3},
	// },
}

// Authentication JSON file
type Authentication struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Token    string `json:"-"`
}

// LoadAuthenticationData loads the authentication file
func LoadAuthenticationData() (*Authentication, error) {

	configFile, err := os.Open("authentication.json")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error opening authentication file")
		return nil, err
	}

	defer configFile.Close()

	var authentication Authentication

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&authentication); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("Error parsing authentication file")
		return nil, err
	}

	return &authentication, nil
}
