package main
import (
  "log"
  "net/http"
  //"math"
  _ "github.com/lib/pq"
  "database/sql"
  "fmt"
)

func initDB(){
  var err error
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
  db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    log.Fatal(err)
  }

  if err = db.Ping(); err != nil {
    log.Fatal(err)
  }
}

func insertPlace(place Place, w http.ResponseWriter){
  _, err := db.Exec("INSERT INTO places (mapsid, barname, etype, dayofweek, comments, lat, long) VALUES($1, $2, $3, $4, $5, $6, $7)",
   place.MapsID, place.Barname, place.EType, place.DayOfWeek, place.Comments, place.Lat, place.Long)
  if err != nil {
    fmt.Print(err)
    http.Error(w, http.StatusText(500), 500)
    return
  }
}
func getAllPlaces(lat string, long string, day string, etype string)([]*Place){
  if(day=="" || day=="All") {day = "%";}
  if(etype=="" || etype=="All") {etype = "%";}
  fmt.Println(etype)
  fmt.Println(day)

  rows, err := db.Query("SELECT ID, MapsID, barname, etype, DayOfWeek, comments, lat, long FROM places where etype like $1 and dayofweek like $2",
                        etype, day)

  if err != nil {
    log.Fatal(err)
  }
  defer rows.Close()

  templaces := make([]*Place,0)
  for rows.Next() {
    plc := new(Place)
    err := rows.Scan(&plc.ID, &plc.MapsID, &plc.Barname, &plc.EType, &plc.DayOfWeek, &plc.Comments, &plc.Lat, &plc.Long)
    if err != nil {
      log.Fatal(err)
    }
    templaces = append(templaces, plc)
  }
  return templaces
}
func deleteDB(id int)(sql.Result){
  res, err := db.Exec("Delete from places where id=$1", id)
  if err != nil {
    log.Fatal(err)
  }
  return res
}
