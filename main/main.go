package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  // "io/ioutil"
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
    [
    {
      "id": "1",
      "slackName": "Ben Cooke",
      "name": "Ben Cooke",
      "url": "ben.cooke"
    },
    {
      "id": "2",
      "slackName": "ben.hackett",
      "name": "Ben Hackett",
      "url": "ben.hackett"
    },
    {
      "id": "3",
      "slackName": "blaise",
      "name": "Blaise Calaycay",
      "url": "blaise.calaycay"
    },
    {
      "id": "4",
      "slackName": "jamesn",
      "name": "James Nicolas",
      "url": "james.nicolas"
    }
    ]
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

// func GetUserRepos(user string)  {
//   // "repos": [
//   //   {
//   //     "name": "AndroidRestService",
//   //     "link": "https://github.com/BenCookie95/AndroidRestService",
//   //     "languages": ["Java"]
//   //   },
//   //   {
//   //     "name": "Blackjack",
//   //     "link": "https://github.com/BenCookie95/Blackjack",
//   //     "languages": ["Java"]
//   //   }
//   // ]
//   var client http.Client
//   var repos []map[string]interface{}
//   resp, err := client.Get("https://api.github.com/users/bencookie95/repos")
//   if err != nil {
//
//   }
//   defer resp.Body.Close()
//
//   jsonData, err := ioutil.ReadAll(resp.Body)
//   json.Unmarshal(jsonData, &repos)
//
//   // repos now has repos data from github for the specified user
//
//   var newRepos []map[string]interface{}
//
//   for i := range repos {
//       repoData := make(map[string]interface{})
//       repo := repos[i]
//       repoData["name"] = repo["name"]
//       repoData["link"] = "https://github.com/" + repo["full_name"].(string)
//
//       // populate languages
//       var languages []map[string]interface{}
//       resp, err := client.Get(repoData["languages_url"])
//       defer resp.Body.Close()
//       jsonLanguageData, err := ioutil.ReadAll(resp.Body)
//       json.Unmarshal(jsonLanguageData, &languages)
//   }
//   // return repos
//   fmt.Println(newRepos)
//
// }

func main() {
  router := mux.NewRouter()
  // fmt.Println(GetUserRepos("ben")[0])
  // GetUserRepos("ben")
  router.HandleFunc("/users/{name}", GetUsers).Methods("GET")
  router.HandleFunc("/user/{name}", GetUser).Methods("GET")
  log.Println("Listening on port 8000")
  log.Fatal(http.ListenAndServe(":8000", router))
}
