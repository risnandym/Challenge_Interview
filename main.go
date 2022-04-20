package main

import (
	"interview/config"
	"interview/docs"
	"interview/routes"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	//programmatically set swagger info
	docs.SwaggerInfo.Title = "Interview Challenge"
	docs.SwaggerInfo.Description = "This is a time calculation and sample server Movie."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	db := config.ConnectDataBase()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	server := routes.SetupRouter(db)
	server.Run()
}
