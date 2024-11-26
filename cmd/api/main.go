package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
}

const webPort = "8080"

func main() {

	log.Println("Starting broker service")

	app := &Config{}

	// define our http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	log.Println("starting broker on port: ", webPort)
	//start, serve and listen
	err := srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}

}
