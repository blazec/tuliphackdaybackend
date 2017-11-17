package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
)


type Users struct {
  Id string `json:"id"`
  SlackName string `json:"slackName"`
  Name string `json:"name"`
  Url string `json:"url"`
}

var name string = "all"

func GetUsers(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  name := params["name"]
  fmt.Println(name)
  jsonData := `
    [{"id": "1",
    "slackName": "Ben Cooke",
    "name": "Ben Cooke",
    "url": "ben.cooke"},
    {"id": "2",
    "slackName": "ben.hackett",
    "name": "Ben Hackett",
    "url": "ben.hackett"}
    ]
  `
  var users []map[string]interface{}
  err := json.Unmarshal([]byte(jsonData), &users)
  if err != nil {
    fmt.Println(err)
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(users)

}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/users/{name}", GetUsers).Methods("GET")
  // router.HandleFunc("/user/{name}", GetUser).Methods("GET")


  log.Fatal(http.ListenAndServe(":8000", router))
}
