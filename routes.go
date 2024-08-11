package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
)

func loadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the API")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "About the API")
}

func routes() *mux.Router {
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler).Methods("GET")
    r.HandleFunc("/about", aboutHandler).Methods("GET")
    return r
}

func main() {
    loadEnv()

    r := routes()

    httpPort := os.Getenv("HTTP_PORT")
    if httpPort == "" {
        log.Fatal("HTTP_PORT environment variable not set")
    }

    fmt.Printf("Starting server on port %s\n", httpPort)
    log.Fatal(http.ListenAndServe(":"+httpPort, r))
}