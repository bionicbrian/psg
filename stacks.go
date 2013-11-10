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

func makePassphrase() string {
    return "happy red hen"
}

func CreateStack (w http.ResponseWriter, r *http.Request) {
    // For vars from the url string using gorilla:
    // vars := mux.Vars(r)
    // title := vars["title"]

    defer r.Body.Close()

    newStack := &Stack{}

    dec := json.NewDecoder(r.Body)
    dec.Decode(newStack)
    newStack.Passphrase = makePassphrase()

    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(newStack)

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
