package env

import "os"

type EnvVars struct {
	HASURA_GRAPHQL_ADMIN_SECRET string
	GRAPHQL_SERVER_URL          string
	JWT_SECRETE                 string
	JWT_REFRESH_SECRETE         string
	SMTP_HOST                   string
	SMTP_PORT                   string
	SMTP_USERNAME               string
	SMTP_PASSWORD               string
	EMAIL_FROM                  string
	EMAIL_VERIFICATION_URL      string
	TICKET_URL                  string
}

var Env EnvVars

func InitEnv() {
	Env = EnvVars{
		HASURA_GRAPHQL_ADMIN_SECRET: os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"),
		GRAPHQL_SERVER_URL:          os.Getenv("GRAPHQL_SERVER_URL"),
		JWT_SECRETE:                 os.Getenv("JWT_SECRETE"),
		JWT_REFRESH_SECRETE:         os.Getenv("JWT_REFRESH_SECRETE"),
		SMTP_HOST:                   os.Getenv("SMTP_HOST"),
		SMTP_PORT:                   os.Getenv("SMTP_PORT"),
		SMTP_USERNAME:               os.Getenv("SMTP_USERNAME"),
		SMTP_PASSWORD:               os.Getenv("SMTP_PASSWORD"),
		EMAIL_FROM:                  os.Getenv("EMAIL_FROM"),
		EMAIL_VERIFICATION_URL:      os.Getenv("EMAIL_VERIFICATION_URL"),
		TICKET_URL:                  os.Getenv("TICKET_URL"),
	}
}
