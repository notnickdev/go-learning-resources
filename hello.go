package main

import (
    "fmt"
    "net/http"
    "log"
)

func getTasks(w http.ResponseWriter, r *http.Request) {
    // Checks if the client is on the correct route
    if r.URL.Path != "/tasks" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    switch r.Method {
    case "GET":
        http.ServeFile(w, r, "tasks.html")
    case "POST":
        if err := r.ParseForm(); err != nil {
            fmt.Fprintf(w, "Something went wrong. Error: %v", err)
            return
        }

        task := r.FormValue("task")
        completed := r.FormValue("completed")

        fmt.Fprintf(w, "Task: %s\n", task)
        fmt.Fprintf(w, "Completed: %s\n", completed)
    default:
        fmt.Fprintf(w, "Only GET and POST requests")
    }
}

func main() {
    http.HandleFunc("/tasks", getTasks)

    fmt.Printf("Starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}