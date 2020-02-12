package main
import (
  "fmt"
  "log"
  "net/http"
  "os"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.newRouter()
    port := getPort()

    r.HandleFunc("/", indexHandler).Methods("GET")
      fmt.Printf("Server up and running. Running on Port: %s\n", port)
      err := http.ListenAndServe(port, r)

      if err != nil {
        log.Fatal("Error listening and server:", err)
      }
}


