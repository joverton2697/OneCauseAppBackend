package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
}

func login(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	token := r.FormValue(`token`)

	t := time.Now()
	tokenHour := t.Hour()
	tokenMinute := t.Minute()
	tokenCompare := "" + strconv.Itoa(tokenHour) + strconv.Itoa(tokenMinute)

	if username != "c137@onecause.com" || password != "#th@nH@rm#y#r!$100%D0p#" || token != tokenCompare {
		resp := make(map[string]string)
		resp["response"] = password + " is not #th@nH@rm#y#r!$100%D0p#"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	} else {
		resp := make(map[string]string)
		resp["response"] = "Success"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
}

func main() {
	http.HandleFunc("/", login)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
