package main

import (
    "fmt"
    "net/http"
)

// CREATE A HTTP SERVER

// http.ResponseWriter assembles the servers response and writes to 
// the client
// http.Request is the clients request
func mainPage(w http.ResponseWriter, r *http.Request) {
    var display string
    display += "<html><body>"
    display += "URL Path: " + fmt.Sprintf("%s\n", r.URL.Path) + "<br>"
    display += "<a href=\"red\">Red Page</a> <br>"
    display += "<a href=\"url\">URL Page</a> <br>"
    display += "</body></html>"
    fmt.Fprintf(w, display)
}

func redPage(w http.ResponseWriter, r *http.Request) {
    var display string
    display += "<html><body style=\"background-color:red\">"
    display += "<a href=\"..\">Back Page</a>"
    display += "</body></html>"
    fmt.Fprintf(w, display)
}

func getURLPage (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Hello, you've requested: %s\n", r.URL.Path)
}

func main() {

	// Calls for function handlers output
    http.HandleFunc("/", mainPage)
    http.HandleFunc("/red", redPage)
    http.HandleFunc("/url", getURLPage)
    
    // Listen to port 8080 and handle requests
    http.ListenAndServe(":8080", nil)
}
