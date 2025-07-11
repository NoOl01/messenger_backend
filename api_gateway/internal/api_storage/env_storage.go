package api_storage

import (
	"os"
)

type EnvStorage struct {
	ApiPort         string
	AuthHost        string
	AuthPort        string
	DbHost          string
	DbPort          string
	SmtpServiceHost string
	SmtpServicePort string
	DebugMode       bool
}

var Env = &EnvStorage{}

func LoadEnv() {
	//err := godotenv.Load("../.env")
	//if err != nil {
	//	log.Fatalf("%s: %s", errs.FailedLoadEnvFile, err.Error())
	//}

	Env.AuthHost = os.Getenv("AUTH_HOST")
	Env.AuthPort = os.Getenv("AUTH_PORT")
	Env.DbHost = os.Getenv("DB_SERVICE_HOST")
	Env.DbPort = os.Getenv("DB_SERVICE_PORT")
	Env.SmtpServiceHost = os.Getenv("SMTP_SERVICE_HOST")
	Env.SmtpServicePort = os.Getenv("SMTP_SERVICE_PORT")
	Env.ApiPort = os.Getenv("API_PORT")
	Env.DebugMode = os.Getenv("DEBUG_MODE") == "true"
}
