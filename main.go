package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT env var must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, "Error")
			log.Print("Error: " + err.Error())
		}
		log.Print("Received: " + string(data))
		fmt.Fprint(w, "Logged")
	})

	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
