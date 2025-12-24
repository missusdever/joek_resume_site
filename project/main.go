package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/gorilla/mux"
)

var port string //Port String

/* TEMPLATE DEFINITION BEGINNING */
var template1 *template.Template

type Config struct {
	Email struct {
		AW_CLIENTID string `yaml:"AW_CLIENTID"`
		AW_CLIENT_SECRET string `yaml:"AW_CLIENT_SECRET"`
		AW_ACCESS_TOKEN string `yaml:"AW_ACCESS_TOKEN"`
		AW_REFRESHTOKEN string `yaml:"AW_REFRESHTOKEN"`
		MY_EMAIL_ADDRESS string `yaml:"MY_EMAIL_ADDRESS"`
		MY_BACKUP_EMAIL string `yaml:"MY_BACKUP_EMAIL"`
		MY_EMAIL_PASSWORD string `yaml:"MY_EMAIL_PASSWORD"`
	} `yaml:"Email"`
}

func init() {
	//Initialize Port
	port = "3000"
	//Initialize template
	template1 = template.Must(template.ParseGlob("./static/templates/*"))
	//Initalize Emails with Credentials
	getCredsYaml() //Get Creds from a Yaml file
}

//Handles the Index page
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving the index page.\n")

	err1 := template1.ExecuteTemplate(w, "index.gohtml", "nil")
	HandleError(w, err1)
}

//Handles the contact page
func contact(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving the contact page.\n")
	//Here is our ViewData struct
	type ViewData struct {
		Host string `json:"Host"`
	}
	vd := ViewData{
		Host: port,
	}

	err1 := template1.ExecuteTemplate(w, "contact.gohtml", vd)
	HandleError(w, err1)
}

//Handles the gamedesign page
func gamedesign(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving the gamedesign page.\n")

	err1 := template1.ExecuteTemplate(w, "gamedesign.gohtml", "nil")
	HandleError(w, err1)
}

//Handles the level design page
func leveldesign(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving the leveldesign page.\n")

	err1 := template1.ExecuteTemplate(w, "leveldesign.gohtml", "nil")
	HandleError(w, err1)
}

//Handles the softwaredesign page
func softwaredesign(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving the softwaredesign page.\n")

	err1 := template1.ExecuteTemplate(w, "softwaredesign.gohtml", "nil")
	HandleError(w, err1)
}

//Handles the test page
func test(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving the test page.\n")
	fmt.Printf("URL Path: %v\n", r.URL.Path)

	err1 := template1.ExecuteTemplate(w, "test.gohtml", "nil")
	HandleError(w, err1)
}

//Handles the writing page
func writing(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving the writing page.\n")

	err1 := template1.ExecuteTemplate(w, "writing.gohtml", "nil")
	HandleError(w, err1)
}

// Handle Errors passing templates
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

func handleRequests() {
	fmt.Printf("DEBUG: Handling Request...\n")

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Handle("/favicon.ico", http.NotFoundHandler()) //For missing FavIcon
	/* Handle passed templates */
	myRouter.HandleFunc("/test", test)
	myRouter.HandleFunc("/", index)
	myRouter.HandleFunc("/contact", contact)
	myRouter.HandleFunc("/gamedesign", gamedesign)
	myRouter.HandleFunc("/leveldesign", leveldesign)
	myRouter.HandleFunc("/softwaredesign", softwaredesign)
	myRouter.HandleFunc("/writing", writing)
	/* HANDLE EMAIL SENDING */
	myRouter.HandleFunc("/emailSubmit", emailSubmit).Methods("POST")

	log.Printf("listenting on port %s\n", port)

	//Serve our static files
	myRouter.Handle("/", http.FileServer(http.Dir("./static")))
	myRouter.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	log.Fatal(http.ListenAndServe(":"+port, myRouter))
}

func main() {
	handleRequests()
}

/* This gets our Creds from a yaml file */
func getCredsYaml() { 
	fmt.Printf("We will be getting creds with yaml rn...\n")

	// Open the YAML file
	file, err := os.Open("security/creds.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the YAML file
	decoder := yaml.NewDecoder(file)
	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}

	//Mongo and Email Creds
	theClientID = config.Email.AW_CLIENTID
	theClientSecret = config.Email.AW_CLIENT_SECRET
	theAccessToken = config.Email.AW_ACCESS_TOKEN
	theRefreshToken = config.Email.AW_REFRESHTOKEN
	myemailAddress = config.Email.MY_EMAIL_ADDRESS
	myemailBKPAddress = config.Email.MY_BACKUP_EMAIL
	myEmailPassword = config.Email.MY_EMAIL_PASSWORD

}
