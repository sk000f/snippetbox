package main

import (
	"log"
	"net/http"
)

// Define home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/".
	// If it doesn't then use the http.NotFound() function to return a 404 response.
	// We then return from the handler to avoid continuing execution
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Handler to show a specific snippet
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet"))
}

// Handler to create a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

func main() {
	// Use the new http.NewServeMux() function to initialise a new servemux,
	// then register handler functions for each URL pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

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
