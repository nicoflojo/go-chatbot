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

func indexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Got my server up and running.")
}

// gets port for cloud deployment
func getPort() string {
  port := os.Getenv("PORT")
  if port == "" {
    port = ":4000"
    fmt.Printf("Port not defined, using Port %s running   port\n", port)
  }
 return port
}
