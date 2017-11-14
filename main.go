package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	addr := ":" + os.Getenv("PORT")
	http.Handle("/", http.FileServer(http.Dir("./views")))
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("error listening: %s", err)
	}
}
