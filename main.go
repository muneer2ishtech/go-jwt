package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Import MySQL driver

	"go-jwt/config"
	"go-jwt/handlers"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Establish database connection
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUsername, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize HTTP server
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		handlers.SignupHandler(w, r, db)
	})
	http.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		handlers.SigninHandler(w, r, db)
	})

	// Start server
	addr := ":" + cfg.Port
	fmt.Printf("Server is listening on port %s...\n", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
