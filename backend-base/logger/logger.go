package logger

import (
	log "github.com/jeanphorn/log4go"
	"os"
)

var Logger *log.Filter
func init() {
	log.LoadConfiguration(os.Getenv("LOG4GO"))
	Logger = log.LOGGER("SERVER")
}
