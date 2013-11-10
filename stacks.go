package main

import (
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Stack struct {
    Title string      `json:"title"`
    Penalty string    `json:"penalty,omitempty"`
    Passphrase string `json:"passphrase"`
}

func CreateStack (w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    // From the JSON request body:
    newStack := &Stack{}

    dec := json.NewDecoder(r.Body)
    dec.Decode(newStack)

    newStack.Passphrase = "happy red hen"

    log.Println(newStack.Title)
    log.Println(newStack.Penalty)
    log.Println(newStack.Passphrase)

    // From the url string using gorilla:
    vars := mux.Vars(r)
    title := vars["title"]
    penalty := vars["penalty"]

    stack := &Stack{Title: title, Penalty: penalty}

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(stack)

    if err != nil {
        log.Fatal("couldn't decode res: ", err)
        return
    }
}

func main () {
    r := mux.NewRouter()

    r.HandleFunc("/stack/{title}", CreateStack).Methods("POST")

    http.Handle("/", r)
    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
