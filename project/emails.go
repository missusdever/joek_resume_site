package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-mail/mail"
)

var theClientID string
var theClientSecret string
var theAccessToken string
var theRefreshToken string
var myemailAddress string
var myemailBKPAddress string
var myEmailPassword string

// Declare DataType from Ajax 11
type UserEmail struct {
	FName     string `json:"FName"`
	LName     string `json:"LName"`
	Email     string `json:"Email"`
	Subject   string `json:"Subject"`
	Message   string `json:"Message"`
	Website   string `json:"Website"`
	AreaCode  string `json:"AreaCode"`
	PhoneNum1 string `json:"PhoneNum1"`
	PhoneNum2 string `json:"PhoneNum2"`
	PhoneNum3 string `json:"PhoneNum3"`
}

// Used to handle emails submitted to us
func emailSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		fmt.Println("DEBUG: We are in emailSubmit.")
		//Collect JSON from Postman or wherever
		//Get the byte slice from the request body ajax
		bs, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}

		//Unmarshal Data
		var emailPosted UserEmail
		json.Unmarshal(bs, &emailPosted)

		m := mail.NewMessage()

		m.SetHeader("From", emailPosted.Email)

		m.SetHeader("To", myemailAddress, myemailBKPAddress)

		m.SetHeader("Subject", emailPosted.Subject)

		m.SetBody("text/html", emailPosted.Message)

		//m.Attach("lolcat.jpg")

		d := mail.NewDialer("smtp.gmail.com", 587, myemailAddress, myEmailPassword)

		// Send the email to Kate, Noah and Oliver.

		if err := d.DialAndSend(m); err != nil {
			type successMSG struct {
				Message     string `json:"Message"`
				SuccessNum  int    `json:"SuccessNum"`
				RedirectURL string `json:"RedirectURL"`
			}
			msgSuccess := successMSG{
				Message:     "Email issue! Please review information sent\n" + err.Error(),
				SuccessNum:  1,
				RedirectURL: "http://" + "localhost:" + port,
			}

			theJSONMessage, err := json.Marshal(msgSuccess)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Fprint(w, string(theJSONMessage))
		} else {
			type successMSG struct {
				Message     string `json:"Message"`
				SuccessNum  int    `json:"SuccessNum"`
				RedirectURL string `json:"RedirectURL"`
			}
			msgSuccess := successMSG{
				Message:     "Email Sent! I hope to get back to you soon!",
				SuccessNum:  0,
				RedirectURL: "http://" + "localhost:" + port,
			}

			theJSONMessage, err := json.Marshal(msgSuccess)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Email successfully submitted.")
			fmt.Fprint(w, string(theJSONMessage))
		}

	}
}
