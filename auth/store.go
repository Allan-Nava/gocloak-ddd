package auth

/*
* Copyright © 2022 Allan Nava <>
* Created 08/02/2022
* Updated 08/02/2022
*
 */

import (
	"fmt"

	"github.com/Allan-Nava/gocloak-ddd/config"
	"github.com/Nerzal/gocloak"
)

type AuthStore interface {
	Login(username string, password string) // need to add jwt
}

type GoCloakAuthStore struct {
	GCloakClient gocloak.GoCloak
}

//
func (s *GoCloakAuthStore) Login(username string, password string) {
	//
	Client := config.CONFIGURATION.KeycloakClientId
	ClientSecret := config.CONFIGURATION.KeycloakSecret
	Realm := config.CONFIGURATION.KeycloakRealm
	//
	jwt, err := s.GCloakClient.Login(Client, ClientSecret, Realm, username, password)
	if err != nil {
		// need to change the status code 400
		fmt.Printf("err %v", err)
		//return nil, err
	}
	fmt.Printf("jwt %v", jwt)
}
