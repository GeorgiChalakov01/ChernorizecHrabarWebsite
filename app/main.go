package main

import (
    "fmt"
    "net/http"

    "github.com/a-h/templ"
    "github.com/GeorgiChalakov01/cherhrab/pages/home"
)

func main() {
    // Serve static files from the images directory
    fs := http.FileServer(http.Dir("./images"))
    http.Handle("/images/", http.StripPrefix("/images/", fs))

    http.Handle("/", http.RedirectHandler("/home", http.StatusSeeOther))
    http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
        templ.Handler(home.Home()).ServeHTTP(w, r)
    })

    fmt.Println("Listening on :8080")
    http.ListenAndServe(":8080", nil)
}
