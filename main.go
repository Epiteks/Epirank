package main

import (
	"os"

	"github.com/Shakarang/Epirank/config"
	"github.com/Shakarang/Epirank/requests"

	log "github.com/Sirupsen/logrus"
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
		return
	}

	requests.Authentication(auth)

}
