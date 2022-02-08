package main
/*
* Copyright © 2022 Allan Nava <>
* Created 08/02/2022
* Updated 08/02/2022
*
 */

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
}