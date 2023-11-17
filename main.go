package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/justjais/go-crm-basic/database"
	"github.com/justjais/go-crm-basic/lead"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database!!")
	}
	fmt.Println("Connection opened to Database!!")
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated!!")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
}
