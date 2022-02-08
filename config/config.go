package config
/*
* Copyright © 2022 Allan Nava <>
* Created 08/02/2022
* Updated 08/02/2022
*
 */

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/sirupsen/logrus"
)

var CONFIGURATION *Configuration

type Configuration struct {
	AppEnv             			string `env:"APP_ENV"`
	LogLevel           			string `env:"LOG_LEVEL"`
	KeycloakAdmin				string `env:"KEYCLOAK_ADMIN"`
	KeycloakAdminPassword		string `env:"KEYCLOAK_ADMIN_PASSWORD"`
	KeycloakUrl					string `env:"KEYCLOAK_URL"`
	KeycloakSecret				string `env:"CLIENT_SECRET"`
	KeycloakRealm				string `env:"KEYCLOAK_REALM"`
	KeycloakClientId			string `env:"KEYCLOAK_CLIENT_ID"`
}

//
func SetEnvConfig() {
	cfg := Configuration{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	logrus.Info("\nload configuration OK")
	CONFIGURATION = &cfg
}
//