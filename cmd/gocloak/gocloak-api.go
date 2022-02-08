package main

/*
* Copyright © 2022 Allan Nava <>
* Created 08/02/2022
* Updated 08/02/2022
*
 */
import (
	"fmt"

	"github.com/Allan-Nava/gocloak-ddd/config"
	"github.com/Allan-Nava/gocloak-ddd/utils"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

// @title GoCloak DDD API Gateway
// @version 1.0
// @description customers and related data op
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	//
	f := fiber.New()
	f.Use(logger.New())
	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type, Authorization",
		AllowMethods: "GET, HEAD, OPTIONS, PUT, PATCH, POST, DELETE",
	}))
	f.Use(recover.New())

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	//
	utils.SetupEnv() //dotenv
	config.SetEnvConfig()

	//swagger setup
	f.Get("/swagger/*", swagger.HandlerDefault) // default

	f.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	fmt.Println("\nGOCloak DDD API")
	fmt.Println("\n2022 ∆ Allan Nava™")
	fmt.Printf("\nENV: %s", config.CONFIGURATION.AppEnv)
	fmt.Printf("\nRUNNING MODE: %s", config.CONFIGURATION.RunningMode)
	fmt.Printf("\nPORT HTTP: %s", config.CONFIGURATION.Port)
	f.Listen("0.0.0.0:8080")
}
