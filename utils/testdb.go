// package main

// import (
// 	"log"
// 	"my-graphql-server/config"
// 	"my-graphql-server/utils"
// )

// func main() {
// 	// Load configuration
// 	cfg, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatal("Error loading config:", err)
// 	}

// 	// Connect to database
// 	db, err := utils.ConnectDB(cfg)
// 	if err != nil {
// 		log.Fatal("Error connecting to database:", err)
// 	}
// 	defer db.Close()

// 	log.Println("Database connection test successful!")
// }
