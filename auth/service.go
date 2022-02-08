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
)
//type AuthStore interface {}

type AuthService struct {
}

//
func (s *AuthService) Login() (*AuthResponse, error) {
	Client := config.CONFIGURATION.KeycloakClientId
	ClientSecret := config.CONFIGURATION.KeycloakSecret
	Realm := config.CONFIGURATION.KeycloakRealm
	//
	fmt.Printf("Client %v | ClientSecret %v | Realm %v", Client, ClientSecret, Realm)
	//jwt, err := GCloakClient.Login(Client, ClientSecret, Realm, username, password)
	return nil, nil
}
//