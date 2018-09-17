package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

func getMessage() string {
    host, _ := os.Hostname()
    return fmt.Sprintf("It works on %s!\n", host)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<html><body><h3>%s</h3><img style=\"width:50%%\" src=\"static/demo1.png\" /></body></html>\n", getMessage())
}

func textHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, getMessage())
}

func handler(w http.ResponseWriter, r *http.Request) {
    ua := r.Header.Get("User-agent")
    if strings.HasPrefix(ua, "curl") {
        textHandler(w, r)
    } else {
        htmlHandler(w, r)
    }
}

func main() {
    http.HandleFunc("/", handler)
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    log.Fatal(http.ListenAndServe(":80", nil))
}


