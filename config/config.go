package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	//ServerTypeAPI - servers only API requests
	ServerTypeAPI = "api"

	//ServerTypeAll - serves API requests and does all exchange stuff
	ServerTypeAll = "all"

	//ServerTypeData - doesn't serve API, only backend stuff
	ServerTypeData = "data"

	//ServerTypeExchange - only exchange individual machine
	ServerTypeExchange = "exchange"

	//ServerTypeAnalysis - only for analysis
	ServerTypeAnalysis = "analysis"

	//ServerTypeTest - for test server
	ServerTypeTest = "test"
)

var ListenPort = ""
var EncryptionSecret = ""
var MongoURL = ""
var SubType = ""
var ServerType = ""

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	ListenPort = os.Getenv("LISTEN_PORT")
	EncryptionSecret = os.Getenv("ENCRYPT_SECRET")
	MongoURL = os.Getenv("READONLY_MONGO")
	SubType = os.Getenv("SUB_TYPE")
}
