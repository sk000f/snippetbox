package main

import (
	"log"
	"net/http"
)

// Define home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Use the new http.NewServeMux() function to initialise a new servemux,
	// then register the home() function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Use the http.ListenAndServe() function to start a new web server.
	// We pass in two parameters:
	//   - the TCP network address to listen on (:4000)
	//   - the servemux we just created
	// If http.ListenAndServe() returns an error we use log.Fatal() to log the error and exit.
	// Any error returned by http.ListenAndServe() is always non-null.
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
