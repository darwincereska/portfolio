package backend

import (
	"net/http"
	"log"
	"encoding/json"
	"strconv"
)

func ParseRoutes() {
	fs := http.FileServer(http.Dir("./dist"))
    http.Handle("/", fs) // Displays astro content quickly
    
    http.HandleFunc("/api/hello/", apiHandler)
}

type Response struct {
 Message string `json:"message"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
 name := r.URL.Path[len("/api/hello/"):]
 if name == "" {
  name = "World"
 }
 response := Response{Message: "Hello " + name}
 w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(response)
}

func Serve(host string, port int) {
	portStr := strconv.Itoa(port)
	log.Printf("Listening at %s on port %s\n", host, portStr)
    err := http.ListenAndServe(host+":"+portStr, nil)
    if err != nil {
        log.Fatal(err)
    }
}