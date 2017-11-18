package main

import (
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "fmt"
  "io/ioutil"
  // "reflect"
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

  repos := GetUserRepos(name)
  languages := GetLanguages(repos)

  userData := make(map[string]interface{})
  userData["id"] = 1
  userData["slackName"] = name
  userData["name"] = name
  userData["url"] = name
  userData["github"] = "https://github.com/" + name
  userData["gitlab"] = ""
  userData["languages"] = languages
  userData["repos"] = repos

  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(userData)

}

func GetLanguages(repos []map[string]interface{}) map[string]int {
    allLanguages := make(map[string]int)

    for iRepo:= range repos {
      repo := repos[iRepo]
      languages := repo["languages"].([]string)
      for iLanguage := 0; iLanguage < len(languages); iLanguage++ {
        language := languages[iLanguage]
        if val, ok := allLanguages[language]; ok {
          allLanguages[language] = val + 1
        } else {
          allLanguages[language] = 1
        }
      }
    }

    return allLanguages
}


func GetUserRepos(user string)  []map[string]interface{} {
  /*
  Format of repos

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
  */
  var client http.Client

  var repos []map[string]interface{}

  resp, err := client.Get("https://api.github.com/users/" + user + "/repos")
  if err != nil {
    log.Println(err)
  }
  defer resp.Body.Close()
  jsonData, err := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(jsonData), &repos)

  // repos now has repos data from github for the specified user
  newRepos := make([]map[string]interface{}, 0)

  for i := range repos {
      repoData := make(map[string]interface{})
      repo := repos[i]

      repoData["name"] = repo["name"]
      repoData["link"] = "https://github.com/" + repo["full_name"].(string)

      /* populate languages */
      var client http.Client
      var languageSizeMap map[string]int
      // languageSizeMap := map[string]int{"Java": 123, "PHP": 132, "Swift": 10}
      languages := make([]string, 0)
      resp, err := client.Get(repo["languages_url"].(string))
      if err != nil {
        log.Println(err)
      }
      defer resp.Body.Close()
      jsonLanguageData, err := ioutil.ReadAll(resp.Body)
      json.Unmarshal(jsonLanguageData, &languageSizeMap)
      for language, _ := range languageSizeMap {
        languages = append(languages, language)
      }
      repoData["languages"] = languages

      newRepos = append(newRepos, repoData)
  }
  return newRepos

}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/users/{name}", GetUsers).Methods("GET")
  router.HandleFunc("/user/{name}", GetUser).Methods("GET")
  log.Println("Listening on port 8000")
  log.Fatal(http.ListenAndServe(":8000", router))
}
