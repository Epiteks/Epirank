package ranking

import (
	"database/sql"
	"fmt"
	"github.com/Shakarang/Epirank/config"
	"github.com/Shakarang/Epirank/database"
	"github.com/Shakarang/Epirank/ticker"
	log "github.com/Sirupsen/logrus"
	"time"
)

var (
	auth *config.Authentication

	// DB instance
	DB *sql.DB

	rankingTicker ticker.Ticker

	// LastRankUpdate is when the rank was updated
	LastRankUpdate time.Time
)

func rankingTick() {

	// First ranking download
	// if err := UpdateRanking(); err != nil {
	// 	log.Error(err)
	// } else {
	// 	LastRankUpdate = time.Now()
	// }

	for {
		<-rankingTicker.Timer.C
		log.Info(time.Now().Unix(), " Update rank")
		if err := UpdateRanking(); err != nil {
			log.Error(err)
		}
		rankingTicker.Update()
	}
}

// InitRanking inits students ranking
func InitRanking(db *sql.DB) error {

	// Create authentication object based on auth file
	auth = config.AuthenticationDataFromEnvironment()

	// If getting data in env failed, quit.
	if auth == nil {
		return fmt.Errorf("Authentification data nil")
	}

	if db == nil {
		return fmt.Errorf("Database nil")
	}

	DB = db

	// Authenticate current user
	if err := Authentication(auth); err != nil {
		return fmt.Errorf("Authentification failed")
	}

	var tickerParameters = ticker.Parameters{
		Hour:     3,
		Minute:   0,
		Interval: 24 * time.Hour,
	}

	rankingTicker = ticker.New(tickerParameters)
	go rankingTick()

	return nil
}

// UpdateRanking updates students ranking data
func UpdateRanking() error {

	if auth == nil || DB == nil {
		return fmt.Errorf("Authentication or Database nil")
	}

	//Retrieve all students data
	data, _ := RequestAllData(auth.Token)

	// Insert in database
	if err := database.InsertData(DB, data); err != nil {
		log.Error(err)
	}

	LastRankUpdate = time.Now()
	return nil
}
