package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "encoding/json"
)

type Person struct {
    ID string `json:"id,omitempty"`
    FirstName string `json:"firstname,omitempty"`
    LastName string `json:"lastname,omitempty"`
    Address *Address `json:"address,omitempty"`
}

type Address struct {
    City string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var peoples []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(peoples)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)

    var e Person

    for _, item := range peoples {
        if item.ID == params["id"] {
            e = item
        }
    }

    json.NewEncoder(w).Encode(e)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]

    peoples = append(peoples, person)
    json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for i, p := range peoples {
        if p.ID == params["id"] {
            peoples = append(peoples[:i], peoples[i+1])
            break
        }
    }

    json.NewEncoder(w).Encode(peoples)
}

func main() {
    peoples = append(peoples, Person{
        "1",
        "John",
        "Doe",
        &Address{
            "City X",
            "State X",
        },
    })
    peoples = append(peoples, Person{
        "2",
        "Jane",
        "Doe",
        &Address{
            "City Y",
            "State Y",
        },
    })
    peoples = append(peoples, Person{
        "3",
        "Jane",
        "Doe",
        &Address{
            "City Y",
            "State Y",
        },
    })


    router := mux.NewRouter()
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}
