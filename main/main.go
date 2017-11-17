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
  name = params["name"]
  fmt.Println(name)
  jsonData := `
    [{
      "id": "1",
      "slackName": "Ben Cooke",
      "name": "Ben Cooke",
      "url": "ben.cooke"},
      {"id": "2",
      "slackName": "ben.hackett",
      "name": "Ben Hackett",
      "url": "ben.hackett",
    }]
  `
  var users []map[string]interface{}
  err := json.Unmarshal([]byte(jsonData), &users)
  if err != nil {
    log.Fatal(err)
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  name = params["name"]

  jsonData := `
    {
      "id": "1",
      "slackName": "Ben Cooke",
      "name": "Ben Cooke",
      "url": "ben.cooke",
      "github": "https://github.com/bencookie95",
      "gitlab": "",
      "languages": ["Java"],
      "repos": [
        {
          "name": "AndroidRestService",
          "link": "https://github.com/BenCookie95/AndroidRestService",
          "languages": ["Java"]
        },
        {
          "name": "Blackjack",
          "link": "https://github.com/BenCookie95/Blackjack",
          "languages": ["Java"]
        }
      ]
    }
  `
  var user map[string]interface{}
  err := json.Unmarshal([]byte(jsonData), &user)
  if err != nil {
    log.Fatal(err)
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(user)

}


func main() {
  router := mux.NewRouter()
  router.HandleFunc("/users/{name}", GetUsers).Methods("GET")
  router.HandleFunc("/user/{name}", GetUser).Methods("GET")
  log.Println("Listening on port 8000")
  log.Fatal(http.ListenAndServe(":8000", router))
}
