package main

import (
    "encoding/json"
    "net/http"
    //"math"
    _ "github.com/lib/pq"
    _ "database/sql"
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "strconv"
)

// create a new item
func CreatePlace(w http.ResponseWriter, r *http.Request) {
    //params := mux.Vars(r)
    var place Place
    _ = json.NewDecoder(r.Body).Decode(&place)
    insertPlace(place, w)
    json.NewEncoder(w).Encode(place)
}

// Display all from the people var
func GetPlaces(w http.ResponseWriter, r *http.Request){
    fmt.Print()
    templaces := getAllPlaces(r.FormValue("lat"),r.FormValue("long"),
                              r.FormValue("day"),r.FormValue("type"))
    json.NewEncoder(w).Encode(templaces)
}

func DeletePlace(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  id, err := strconv.Atoi(params["id"])
  if err != nil {
    log.Fatal(err)
  }
  res, _ := deleteDB(id).RowsAffected()
  fmt.Print(res)

  json.NewEncoder(w).Encode(res)
}
