package main

import (
	"log"
	_ "fiber-app/docs"   

	"fiber-app/internal/app"
)


// @title Fiber Production API
// @version 1.0
// @description Production ready Fiber backend
// @termsOfService https://example.com/terms

// @contact.name Aman
// @contact.email aman@example.com

// @host localhost:3010
// @BasePath /api/v1

// 🔐 Auth
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization



//fior exaple 


frwnjhwfv
func main() {
	app := app.New()
    log
	log.Fatal(app.Listen(":3010"))
}
