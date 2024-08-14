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
        log.Fatalf("Error loading .env file: %v", err)
    }
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    _, err := fmt.Fprintf(w, "Welcome to the API")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        log.Printf("homeHandler error: %v", err)
    }
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    _, err := fmt.Fprintf(w, "About the API")
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        log.Printf("aboutHandler error: %v", err)
    }
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
    err := http.ListenAndServe(":"+httpPort, r)
    if err != nil {
        log.Fatalf("Failed to start the server on port %s: %v", httpPort, err)
    }
}