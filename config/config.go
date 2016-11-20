package config

import (
	"encoding/json"
	"os"

	"github.com/Shakarang/Epirank/models"
	log "github.com/Sirupsen/logrus"
)

// Promotions list
var Promotions = []models.Promotion{
	models.Promotion{
		Name:            "tek1",
		AcademicProgram: "master",
	},
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
