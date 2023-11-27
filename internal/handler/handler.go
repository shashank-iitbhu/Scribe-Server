package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, I'm Scribe!")
}

// HandleRequests handles incoming HTTP requests.
func HandleRequests() {

	// Setup root handler.
	http.HandleFunc("/", hello)

	// Setup /files handler.
	fileSystem := viper.GetString("fileSystem")
	http.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(fileSystem))))

	// Start serving requests.
	hostPort := fmt.Sprintf(":%s", viper.GetString("hostPort"))
	log.Printf("listening on port %s", hostPort)
	log.Fatal(http.ListenAndServe(hostPort, nil))
}
