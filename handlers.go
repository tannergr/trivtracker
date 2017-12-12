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
    "io/ioutil"
  	"html/template"
  	"golang.org/x/net/context"
)

// create a new item
func CreatePlace(w http.ResponseWriter, r *http.Request) {
    //params := mux.Vars(r)
    user := getUser(w,r)
    if (user!=nil && checkDBadmin(user.Sub)){
      var place Place
      _ = json.NewDecoder(r.Body).Decode(&place)
      insertPlace(place, w)
      w.Write("entered production")
    } else {
      var place Place
      _ = json.NewDecoder(r.Body).Decode(&place)
      insertSuggestedPlace(place, w)
      w.Write("Suggested")
    }
}
func GetPlace(w http.ResponseWriter, r *http.Request){
  params := mux.Vars(r)
  id, err := strconv.Atoi(params["id"])
  if err != nil {
    log.Fatal(err)
  }
  templace := getPlaceDB(id)
  json.NewEncoder(w).Encode(templace)
}

// Display all from the people var
func GetPlaces(w http.ResponseWriter, r *http.Request){
    templaces := getAllPlaces(r.FormValue("lat"),r.FormValue("long"),
                              r.FormValue("day"),r.FormValue("type"))
    json.NewEncoder(w).Encode(templaces)
}

func DeletePlace(w http.ResponseWriter, r *http.Request){
  user := getUser(w,r)
  if (user != nil && checkDBadmin(user.Sub)){
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
      log.Fatal(err)
    }
    res, _ := deleteDB(id).RowsAffected()
    fmt.Print(res)

    json.NewEncoder(w).Encode(res)
  } else {
    http.Error(w, "Permission Denied", 403)
  }
}

func loginHandler(w http.ResponseWriter, r *http.Request){
  fmt.Println("Login Handler OK")
  state = randToken()
	store, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	store.Values["state"] = state
  fmt.Println(store.Values["state"])
	err = store.Save(r, w);
  if err != nil{
    fmt.Println(err.Error())
  }
  loginUrl := struct{ Url string }{}
	loginUrl.Url = getLoginURL(state)
	templ := template.Must(template.ParseFiles("./static/login.html"))
	templ.Execute(w, loginUrl)
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	store, err := store.Get(r, "session")
	if err != nil {
		httpError(w, err, "Getting store")
		return
	}
	query := r.URL.Query()
  fmt.Println(query)
	retrievedState := store.Values["state"]
	if retrievedState != query.Get("state") {
		httpError(w, err, "Getting state from store")
		return
	}

	token, err := conf.Exchange(context.Background(), query.Get("code"))
	if err != nil {
		httpError(w, err, "token bit")
		return
	}
	client := conf.Client(context.Background(), token)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
  fmt.Println(email)
	if err != nil {
		httpError(w, err, "email bit")
		return
	}
	defer email.Body.Close()
	data, _ := ioutil.ReadAll(email.Body)
	user := User{}
	json.Unmarshal(data, &user)
	store.Values["user"] = user
	store.Save(r, w)
	fmt.Println("Email body: ", string(data))
	http.Redirect(w, r, "/user", 301)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
  user := getUser(w,r)
  if checkDBadmin(user.Sub){
    fmt.Print("succeeedsd")
    templ := template.Must(template.ParseFiles("./static/add.html"))
  	templ.Execute(w, user)
  } else {
    fmt.Print("Access Denied")
    templ := template.Must(template.ParseFiles("./static/userDenied.html"))
  	templ.Execute(w, user)
  }
}
func getLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}
func httpError(w http.ResponseWriter, err error, reason string) {
	fmt.Println(reason)
	http.Error(w, err.Error(), http.StatusInternalServerError)
	return
}

func getUser(w http.ResponseWriter, r *http.Request)(*User){
  session, err := store.Get(r, "session")
  if err != nil {
    httpError(w, err, "getting session")
    panic(" ")
  }

  val := session.Values["user"]
  //changed from panicing when no user cookies to return nil
  if _, ok := val.(*User); !ok {
    return nil
  }
  return val.(*User)
}
