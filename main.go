package main

import (
    "encoding/base64"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
    b := make([]byte, length)
    for i := range b {
        b[i] = charset[rand.Intn(len(charset))]
    }
    return string(b)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    now := time.Now().Format(time.RFC3339)
    token := r.URL.Query().Get("token")

    if token == "" {
        token = generateRandomString(16)
        token = base64.StdEncoding.EncodeToString([]byte(token))
        w.Header().Set("AUTH-TOKEN", token)
        w.Header().Set("AUTH0-TOKEN", token)
        w.Header().Set("AUTH1_TOKEN", token)
        fmt.Fprintf(w, "Hello, World! Generated random token at: %s", now)
    } else {
        encodedToken := base64.StdEncoding.EncodeToString([]byte(token))
        w.Header().Set("AUTH_TOKEN", encodedToken)
        fmt.Fprintf(w, "Hello, World! Time: %s", now)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet, http.MethodPost:
            handleRequest(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    fmt.Println("Server listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
