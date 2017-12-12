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

	dbmanager, err := db.CreateDynamoManager()
	if err != nil {
		log.Fatal(err)
	}

	app := app.InitApp(dbmanager)
	server := server.Init(app)
	http.Handle("/", http.FileServer(http.Dir("./views")))
	err = http.ListenAndServe(addr, server)
	log.Fatalf("error listening: %s", err)
}
