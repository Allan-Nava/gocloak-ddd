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

type AuthService struct {
}

//
func (s *AuthService) Login(username string, password string) (*AuthResponse, error) {
	Client := config.CONFIGURATION.KeycloakClientId
	ClientSecret := config.CONFIGURATION.KeycloakSecret
	Realm := config.CONFIGURATION.KeycloakRealm
	//
	//client := gocloak.NewClient("https://mycool.keycloak.instance")
	fmt.Printf("Client %v | ClientSecret %v | Realm %v", Client, ClientSecret, Realm)
	return nil, nil
	/*jwt, err := GCloakClient.Login(Client, ClientSecret, Realm, username, password)

	if err != nil {
		// need to change the status code 400
		return nil, err
	}
	var resp = AuthResponse{Token: jwt.}

	return &resp, nil*/

}

//
