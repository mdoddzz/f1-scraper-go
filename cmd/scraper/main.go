package main

import (
	"fmt"

	"github.com/mdoddzz/f1-scraper-go/pkg/scraper"
	"github.com/mdoddzz/f1-scraper-go/pkg/storage/mongo"
)

func main() {

	fmt.Println("Starting F1 Scraper")

	// Get database connection and initialise storage
	db := mongo.NewDB("mongodb://localhost:27017/?readPreference=primary&ssl=false&directConnection=true", "formula-custom")
	storage := mongo.NewStorage(db)
	fmt.Println("Storage Linked")

	// Initialise service
	service := scraper.NewWithMongo(storage)
	fmt.Println("Scraper Service Created")

	// Start service
	fmt.Println("Scraper Starting...")
	service.Start()
}
