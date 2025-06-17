package main

import (
    "io"
    "log"
    "net/http"
    "strings"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
    r := chi.NewRouter()
    r.Use(middleware.Logger)

    r.Post("/", func(w http.ResponseWriter, r *http.Request) {
        defer r.Body.Close()
        bodyBytes, err := io.ReadAll(r.Body)
        if err != nil {
            http.Error(w, "Error reading request body", http.StatusInternalServerError)
            return
        }
        if len(bodyBytes) == 0 {
            http.Error(w, "Empty request body", http.StatusBadRequest)
            return
        }
        w.Header().Set("Content-Type", "text/plain; charset=utf-8")
        w.Write([]byte(strings.ToUpper(string(bodyBytes))))
    })

    if err := http.ListenAndServe(":3009", r); err != nil {
        log.Fatal(err)
    }
}
