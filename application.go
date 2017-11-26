package main

import (
    "github.com/gorilla/mux"
  	"github.com/gorilla/sessions"
    "log"
    "net/http"
    //"fmt"
    	"crypto/rand"
    	"encoding/base64"
    	"encoding/gob"
    	//"encoding/json"
    	"golang.org/x/oauth2"
    	"golang.org/x/oauth2/google"
    	//"io/ioutil"
      "os"
    _ "github.com/lib/pq"
    "database/sql"
)
type Credentials struct {
	Cid     string `json:"cid"`
	Csecret string `json:"cs"`
}

var cred Credentials
var conf *oauth2.Config
var state string
var store = sessions.NewCookieStore([]byte("something-very-secret"))

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
type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Gender        string `json:"gender"`
}


var db *sql.DB


func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func init() {
  initDB();
	store.Options = &sessions.Options{
		Domain:   "127.0.0.1",
		Path:     "/",
		MaxAge:   3600 * 8, // 8 hours
		HttpOnly: true,
	}

	gob.Register(&User{})
	// file, err := ioutil.ReadFile("./creds.json")
	// if err != nil {
	// 	fmt.Printf("File Error: %v\n", err)
	// 	os.Exit(1)
	// }
  //
	// json.Unmarshal(file, &cred)
  cred.Cid = os.Getenv("Cid")
  cred.Csecret = os.Getenv("Csecret")

	conf = &oauth2.Config{
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		RedirectURL:  "http://127.0.0.1:8000/auth",
		Scopes: []string{
			// scopes allow you to selectively choose the permissions you need access to
			// for simple login you can just use userinfo.email
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/places", GetPlaces).Methods("GET")
    router.HandleFunc("/places", CreatePlace).Methods("PUT")
    router.HandleFunc("/places/{id}", DeletePlace).Methods("DELETE")
    //router.HandleFunc("/edit", adminHome)
    router.HandleFunc("/login", loginHandler)
  	router.HandleFunc("/auth", authHandler)
  	router.HandleFunc("/user", userHandler)

    router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
