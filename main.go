package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
    //"math"
    _ "github.com/lib/pq"
    "database/sql"
)

// The person Type (more like an object)
type Place struct {
    ID          string      `json:"id,omitempty"`
    MapsID      string      `json:"mapsid,omitempty"`
    Barname     string      `json:"barname,omitempty"`
    EType       string      `json:"etype,omitempty"`
    DayOfWeek   string      `json:"dayofweek,omitempty"`
    Comments   string       `json:"comments,omitempty"`
    Lat         float64     `json:"lat,omitempty"`
    Long        float64     `json:"long,omitempty"`
}

var db *sql.DB

func init() {
  initDB()
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/places", GetPlaces).Methods("GET")
    router.HandleFunc("/places", CreatePlace).Methods("PUT")
    router.HandleFunc("/places/{id}", DeletePlace).Methods("DELETE")
    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    log.Fatal(http.ListenAndServe(":3000", router))
}
