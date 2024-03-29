package app

import (
	"log"
	"os"
	"strconv"
)

type config struct {
	BOTTOKEN_CELEBOT                 string
	LONGPOLLING_WORKERS              int
	DEFAULT_DELAY_BETWEEN_CHECKS_SEC int
	BD_NOTIFICATION_HOUR_MOSCOW_TZ   int
	CLUBCODE                         string
	ADMINCODE                        string
	LOG_FILE                         string
}

var config_ config

func init() {
	config_.BOTTOKEN_CELEBOT = os.Getenv("BOTTOKEN_CELEBOT")
	longPolingWorkers, err := strconv.Atoi(os.Getenv("LONGPOLLING_WORKERS"))
	if err != nil {
		log.Fatalf("Parse var error for: LONGPOLLING_WORKERS")
	}
	config_.LONGPOLLING_WORKERS = longPolingWorkers
	seconds, err := strconv.Atoi(os.Getenv("DEFAULT_DELAY_BETWEEN_CHECKS_SEC"))
	if err != nil {
		log.Fatalf("Parse var error for: DEFAULT_DELAY_BETWEEN_CHECKS_SEC")
	}
	config_.DEFAULT_DELAY_BETWEEN_CHECKS_SEC = seconds
	config_.BD_NOTIFICATION_HOUR_MOSCOW_TZ, err = strconv.Atoi(os.Getenv("BD_NOTIFICATION_HOUR_MOSCOW_TZ"))
	if err != nil {
		log.Fatalf("Parse var error for: BD_NOTIFICATION_HOUR_MOSCOW_TZ")
	}
	config_.CLUBCODE = os.Getenv("CLUBCODE")
	config_.ADMINCODE = os.Getenv("ADMINCODE")
	config_.LOG_FILE = os.Getenv("LOG_FILE")
}

func GetConfig() *config {
	return &config_
}
