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
	// set Content Type header and return JSON response
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"Response":"Display a snippet"}`))
}

// Handler to create a new snippet
func createSnippet(w http.ResponseWriter, r *http.Request) {
	// Check the request method to block anything that isn't a POST
	if r.Method != http.MethodPost {
		// if request is not a POST then update the response header to include 405 status
		// and write appropriate message in the response body.

		// add cache-control header
		w.Header().Add("Cache-Control", "public")
		w.Header().Add("Cache-Control", "max-age=31536000")
		// Also set an Allow header to inform client what is accepted
		w.Header().Set("Allow", http.MethodPost)
		// use http.Error to automatically call w.WriteHeader and w.Write for us
		http.Error(w, "Method Not Allowed", 405)
		return
	}
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
