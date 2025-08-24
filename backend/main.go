package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	config "github.com/triyaambak/CRM/config"
	routes "github.com/triyaambak/CRM/routes"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Something went wrong while reading .env file :", err)
	}

	var port string = os.Getenv("API_PORT")
	adr := fmt.Sprintf(":%s", port)

	//Connecting to db
	dbCfg := config.NewDBConfig()
	defer dbCfg.Close()

	router := chi.NewMux()

	routes.SetUpRoutes(router, dbCfg)

	server := http.Server{
		Addr:    adr,
		Handler: router,
	}

	fmt.Println("Starting server on port " + adr)
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println("Server crashed")
		log.Fatal(err)
	}

	fmt.Println("Server closed on port " + adr)
}
