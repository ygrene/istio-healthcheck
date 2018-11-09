package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	ADDR = os.Getenv("HEALTHCHECK_ADDRESS_AND_PORT")
	PORT = os.Getenv("LISTEN_PORT")
)

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		//According to Go, their client library will send nothing as a method if its an HTTP GET
		if r.Method == "GET" || r.Method == "" {
			log.Println("Handling upstream connection to", ADDR)
			resp, err := http.Get(ADDR)
			if err != nil {
				log.Fatalln(err)
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("Got upstream status:", resp.StatusCode)
			w.WriteHeader(resp.StatusCode)
			w.Write(body)
		}
	})
	http.ListenAndServe(PORT, nil)
}
