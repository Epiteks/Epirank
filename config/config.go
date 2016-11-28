package config

import (
	"encoding/json"
	"os"

	"github.com/Shakarang/Epirank/models"
	log "github.com/Sirupsen/logrus"
)

// DatabasePath is database path...
const DatabasePath = "/tmp/students.db"

const tek1 = "tek1"
const tek2 = "tek2"
const tek3 = "tek3"
const tek4 = "tek4"
const tek5 = "tek5"

// Cities list
var Cities = []models.City{
	{
		Name:       "All",
		ID:         "",
		Promotions: []string{tek1, tek2, tek3, tek4, tek5},
	},
	{
		Name:       "Bordeaux",
		ID:         "BDX",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Lille",
		ID:         "LIL",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Lyon",
		ID:         "LYN",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Marseille",
		ID:         "MAR",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Montpellier",
		ID:         "MPL",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Nancy",
		ID:         "NCY",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Nantes",
		ID:         "NAN",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Nice",
		ID:         "NCE",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Paris",
		ID:         "PAR",
		Promotions: []string{tek1, tek2, tek3, tek4, tek5},
	},
	{
		Name:       "Rennes",
		ID:         "REN",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Strasbourg",
		ID:         "STG",
		Promotions: []string{tek1, tek2, tek3},
	},
	{
		Name:       "Toulouse",
		ID:         "TLS",
		Promotions: []string{tek1, tek2, tek3},
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
		}).Error("Error opening authentication file")
		return nil, err
	}

	defer configFile.Close()

	var authentication Authentication

	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(&authentication); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error parsing authentication file")
		return nil, err
	}

	return &authentication, nil
}

// AuthenticationDataFromEnvironment get data from Env
func AuthenticationDataFromEnvironment() *Authentication {

	var authentication Authentication

	authentication.Login = os.Getenv("EPIRANK_LOGIN")
	authentication.Password = os.Getenv("EPIRANK_PASSWORD")

	if authentication.Login == "" || authentication.Password == "" {
		log.Fatal("Could not get EPIRANK_LOGIN and EPIRANK_PASSWORD")
		return nil
	}
	return &authentication
}
