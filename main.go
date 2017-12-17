package main

import (
	"log"
	"net/http"
	"os"

	"github.com/hesidoryn/receiveDASHsmsDOTonline/app"
	"github.com/hesidoryn/receiveDASHsmsDOTonline/db"
	"github.com/hesidoryn/receiveDASHsmsDOTonline/server"
)

func main() {
	addr := ":" + os.Getenv("PORT")

	dbmanager, err := db.CreateDynamoManager(os.Getenv("TABLE_NAME"))
	if err != nil {
		log.Fatal(err)
	}

	app := app.Init(dbmanager)
	server := server.Init(app)
	err = http.ListenAndServe(addr, server)
	log.Fatalf("error listening: %s", err)
}
