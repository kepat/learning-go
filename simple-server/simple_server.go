package main

import (
    "fmt"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "math/rand"
)

// CREATE A HTTP SERVER

// http.ResponseWriter assembles the servers response and writes to 
// the client
// http.Request is the clients request
func mainPage(w http.ResponseWriter, r *http.Request) {
    var display string
    display += "<html><body style=\"text-align: center\">"
    display += "URL Path: " + fmt.Sprintf("%s\n", r.URL.Path) + "<br><br>"
    display += "<a href=\"pink\">Pink Page</a> <br>"
    display += "<a href=\"url\">URL Page</a> <br>"
    display += "<a href=\"create\">Create Name in Local DB</a> <br>"
    display += "<a href=\"display\">Display Name from Local DB</a> <br>"
    display += "<a href=\"delete\">Clear All Name in Local DB</a> <br>"
    display += "</body></html>"
    fmt.Fprintf(w, display)
}

func pinkPage(w http.ResponseWriter, r *http.Request) {
    var display string
    display += "<html><body style=\"background-color: pink\">"
    display += "<a href=\"..\">Back Page</a>"
    display += "</body></html>"
    fmt.Fprintf(w, display)
}

func getURLPage (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Hello, you've requested: %s\n", r.URL.Path)
}

func createName (w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:@/go")
    if err != nil {
        fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Failed : %s\n", err)
    }
    defer db.Close()

    // Prepare statement for inserting data
    stmtIns, err := db.Prepare("INSERT INTO user (name) VALUES( ? )")
    if err != nil {
        fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Failed : %s\n", err)
    }
    defer stmtIns.Close()

    var name string = randomString(10)

    stmtIns.Exec(name)

    fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Success! Created " + name + ".")
}

func deleteAllName (w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:@/go")
    if err != nil {
        fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Failed : %s\n", err)
    }
    defer db.Close()

    // Prepare statement for inserting data
    stmtIns, err := db.Prepare("DELETE FROM user")
    if err != nil {
        fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Failed : %s\n", err)
    }
    defer stmtIns.Close()

    stmtIns.Exec()

    fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Success")
}

func displayNames (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html>")
    db, err := sql.Open("mysql", "root:@/go")
    if err != nil {
        fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Failed : %s\n", err)
    }
    defer db.Close()

    // Prepare statement for reading data
    rows, err := db.Query("SELECT name FROM user")
    if err != nil {
        fmt.Fprintf(w, "<a href=\"..\">Back Page</a> Failed : %s\n", err)
    }

    var name string;

    for rows.Next() {
        err = rows.Scan(&name)
        if err != nil {
             fmt.Fprintf(w, "Failed : %s\n", err)
        }

        fmt.Fprintf(w, fmt.Sprintf("%s\n", name) + "<br>")
    }

    fmt.Fprintf(w, "<a href=\"..\">Back Page</a>")
    fmt.Fprintf(w, "</html>")
}

func main() {

	// Calls for function handlers output
    http.HandleFunc("/", mainPage)
    http.HandleFunc("/pink", pinkPage)
    http.HandleFunc("/url", getURLPage)
    http.HandleFunc("/create", createName)
    http.HandleFunc("/display", displayNames)
    http.HandleFunc("/delete", deleteAllName)
    
    // Listen to port 8080 and handle requests
    http.ListenAndServe(":8080", nil)
}

func randomString(n int) string {
    var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    b := make([]rune, n)
    for i := range b {
        b[i] = letter[rand.Intn(len(letter))]
    }
    return string(b)
}
