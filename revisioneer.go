package main

import (
  "io"
  "time"
  "log"
  "fmt"
  "github.com/gorilla/mux"
  "net/http"
  "encoding/json"
)

type Deploy struct {
  DeployedAt time.Time
  Sha string
}

var revisions []Deploy = make([]Deploy, 0)

func ListRevisions(w http.ResponseWriter, req *http.Request) {
  b, err := json.Marshal(revisions)
  if err == nil {
    io.WriteString(w, string(b))
  } else {
    io.WriteString(w, "[]")
  }
}

func CreateRevision(w http.ResponseWriter, req *http.Request) {
  dec := json.NewDecoder(req.Body)

  var deploy Deploy
  err := dec.Decode(&deploy)
  if err != nil && err != io.EOF {
    log.Fatal(err)
  } else {
    deploy.DeployedAt = time.Now()
  }

  revisions = append(revisions, deploy);
  io.WriteString(w, "")
}

func init() {
  r := mux.NewRouter()
  r.HandleFunc("/revisions", ListRevisions).
    Methods("GET")
  r.HandleFunc("/revisions", CreateRevision).
    Methods("POST")
  http.Handle("/", r)
}

func main() {
  fmt.Printf("Server listening on port %d\n", 8080)
  http.ListenAndServe(":8080", nil)
}