package auth

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/Nerzal/gocloak/v10"
	log "github.com/sirupsen/logrus"
	"gopkg.in/square/go-jose.v2"
)


var Server string
var Realm *string
var Client string
var Certificates jose.JSONWebKeySet
var GCloakClient KeycloakMutexWrapper
var AdminJWT *gocloak.JWT
var ClientSecret string
var ClientID string
var RealmManagementClientId string

// keycloak Mutex Wrapper to prevent leak memory and problem with concurrent access
type KeycloakMutexWrapper struct {
	Mu sync.Mutex
	gocloak.GoCloak
}


// Init Keycloak one time
func InitKeycloak(ctx context.Context, username string, password string, realm string, server string) error {
	//
	GCloakClient = KeycloakMutexWrapper{
		Mu:      sync.Mutex{},
		GoCloak: gocloak.NewClient(server),
	}

	GCloakClient.Mu.Lock()
	defer GCloakClient.Mu.Unlock()

	//
	grant := "password"
	clientId := "admin-cli"
	Client = "compress_api" // need to fix the client name not hardcoded
	// set realm global
	//Realm = &realm
	log.Debugf("InitKeycloak into keycloak GetAdminToken %v", realm)
	//
	var err error
	AdminJWT, err = GCloakClient.GetToken(ctx, realm, gocloak.TokenOptions{
		GrantType: &grant,
		ClientID:  &clientId,
		Username:  &username,
		Password:  &password,
	})
	if err != nil {
		log.Errorf("error during login admin for realm %s, err: %s", Realm, err.Error())
		return utils.NewError("AuthKeycloak", err.Error())
	}
	//
	clients, err := GCloakClient.GetClients(ctx, AdminJWT.AccessToken, realm, gocloak.GetClientsParams{})
	if err != nil {
		log.Error(err)
	}
	//
	for _, client := range clients {
		if client.ClientID != nil && *client.ClientID == Client {
			ClientID = *client.ID
			log.Debugf("clientId: %v , Client: %v \n ", ClientID, Client)
			cr, err := GCloakClient.GetClientSecret(ctx, AdminJWT.AccessToken, realm, ClientID)
			//
			if err == nil {
				//fmt.Printf("clientsecret : %v\n", *cr.Value)
				log.Debugf("clientId: %v , Client: %v | clientSecret: %v", ClientID, Client, *cr.Value)
				//fmt.Println(*cr.Value)
				ClientSecret = *cr.Value
				fmt.Printf(" ClientSecret %s | Client %s \n", ClientSecret, Client)
			} else {
				log.Error("err in getting client secret ", err)
			}
			//break
		}
		if client.ClientID != nil && *client.ClientID == REALM_MANAGEMENT_CLIENT {
			RealmManagementClientId = *client.ID
			log.Debugf("RealmManagementClientId: %v , Client: %v ", RealmManagementClientId, REALM_MANAGEMENT_CLIENT)
		}
	}
	//
	//return adminToken, nil
	return nil
}

func GetAdminToken(ctx context.Context, username string, password string, realm string) (*gocloak.JWT, error) {
	log.Debug("into keycloak GetAdminToken")

	client := gocloak.NewClient("https://auth.tngrm.io")
	grant := "password"
	clientId := "admin-cli"
	log.Infof("into keycloak GetAdminToken %v", realm)
	adminToken, err := client.GetToken(ctx, realm, gocloak.TokenOptions{
		GrantType: &grant,
		ClientID:  &clientId,
		Username:  &username,
		Password:  &password,
	})
	if err != nil {
		log.Errorf("error during login admin for realm %s, err: %s", realm, err.Error())
		return nil, utils.NewError("AuthKeycloak", err.Error())
	}

	return adminToken, nil
}

func GetToken(ctx context.Context, username string, password string, clientId string, realm string) (*gocloak.JWT, error) {
	log.Debug("into keycloak GetToken")
	grant := "password"
	token, err := GCloakClient.GetToken(ctx, realm, gocloak.TokenOptions{
		GrantType: &grant,
		ClientID:  &clientId, //customername_client
		Username:  &username,
		Password:  &password,
	})
	if err != nil {
		log.Errorf("error during get token admin for create client: %s", err.Error())
		return nil, utils.NewError("AuthKeycloak", err.Error())
	}

	return token, nil
}

// refresh
func RefreshToken(ctx context.Context, refreshToken string, realm string, clientID string) (*gocloak.JWT, error) {
	//
	log.Debug("into keycloak RefreshToken")
	//getting admin token for realm
	tokenRefreshed, err := GCloakClient.RefreshToken(ctx, refreshToken, clientID, ClientSecret, realm)
	if err != nil {
		log.Errorf("error during refresh token admin for realm %s, err: %s", realm, err.Error())
		return nil, utils.NewError("AuthKeycloak", err.Error())
	}

	return tokenRefreshed, nil
}
