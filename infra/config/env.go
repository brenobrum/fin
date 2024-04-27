package config

import (
	"fmt"
	"os"
)

var errors []string

func SetupEnv() {
	validateEnv("MONGO_URI")

	defaultEnv("JWT_AUTH_EXP_MINUTES", "30")
	validateEnv("JWT_AUTH_SECRET_KEY")

	defaultEnv("REFRESH_AUTH_EXP_MINUTES", "1440")
	validateEnv("REFRESH_AUTH_SECRET_KEY")

	defaultEnv("MONGO_DATABASE", "develop")

	for _, err := range errors {
		fmt.Println(err)
	}

	if len(errors) > 0 {
		os.Exit(0)
	}
}

func validateEnv(envName string) {
	env := os.Getenv(envName)
	if env == "" {
		errors = append(errors, fmt.Sprintf("no env %s", envName))
	}
}

func defaultEnv(envName, defaultValue string) {
	env := os.Getenv(envName)
	if env == "" {
		os.Setenv(envName, defaultValue)
	}
}
