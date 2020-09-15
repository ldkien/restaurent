package app

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"restaurant/backend-base/entity"
	log "restaurant/backend-base/logger"
)

var GlobalConfig entity.GlobalConfig

func init() {
	fileConfig := os.Getenv(ConfigPath)
	jsonFile, err := os.Open(fileConfig)
	if err != nil {
		log.Logger.Error(err)
		panic(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &GlobalConfig)

}
