package main

import (
	//"database/sql"
	//"fmt"
	"github.com/Shakarang/Epirank/config"
	"github.com/Shakarang/Epirank/database"
	"github.com/Shakarang/Epirank/requests"
	log "github.com/Sirupsen/logrus"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func init() {

	// Log as ASCII Formatter
	log.SetFormatter(&log.TextFormatter{})

	// Output to stdout, could also be a file.
	log.SetOutput(os.Stdout)

	// Log everything
	log.SetLevel(log.InfoLevel)
}

func main() {

	// Create authentication object based on auth file
	var auth, err = config.LoadAuthenticationData()

	// If getting data in  authentication file failed, quit.
	if err != nil {
		os.Exit(-1)
	}

	// Authenticate current user
	if err := requests.Authentication(auth); err != nil {
		os.Exit(-1)
	}
	// Retrieve all students data
	data, _ := requests.RequestAllData(auth.Token)

	if db, err := database.Init(config.DatabasePath); err != nil {
		log.Fatal(err)
	} else {
		defer db.Close()
		database.CreateTable(db)

		if err := database.InsertData(db, data); err != nil {
			log.Error(err)
		}

		// var city = "STG"
		// var promo = "tek1"
		// database.GetStudentsFrom(db, &city, &promo)
	}

}
