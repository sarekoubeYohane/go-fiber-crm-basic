package main

import (
	"fmt"
	"go-fiber-crm-basic/database"
	"go-fiber-crm-basic/lead"

	"log"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/GetLeads", lead.GetLeads)
	app.Get("/api/v1/GetLeads/:id", lead.GetLead)
	app.Post("/api/v1/GetLeads", lead.NewLead)
	app.Delete("/api/v1/GetLeads/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBconn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		log.Panic("failed to connect database")
	}
	database.DBconn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBconn.Close()
}
