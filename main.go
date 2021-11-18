package main

import (
    "fmt"
    "gee"
    "net/http"
)

func main() {
    app := gee.New()

    app.GET("/", func(w http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(w, "URL.Path = %s\n", req.URL.Path)
    })

    app.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
        for k, v := range req.Header {
            fmt.Fprintf(w, "Header[%q] = %q", k, v)
        }
    })
    app.Run(":9999")
}