package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	// App port
	HTTPPort := ":8086"
	if len(os.Getenv("PORT")) > 0 {
		HTTPPort = ":" + os.Getenv("PORT")
	}

	log.Printf("listening on %s\n", HTTPPort)

	err := http.ListenAndServe(HTTPPort, stdlibSrv{})
	if err != nil {
		log.Fatal(err)
	}
}

type stdlibSrv struct{}

func (s stdlibSrv) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(`{ "greeting" : "Hello world!" }`))
}
