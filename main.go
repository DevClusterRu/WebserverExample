package main

import (
	"Webserver/requests"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/getToken", requests.GetToken)
	http.HandleFunc("/getRates", requests.GetRates)
	http.HandleFunc("/getBalance", requests.GetBalance)
	http.HandleFunc("/transaction", requests.SetTransaction)

	log.Println("Starting webserver...")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}

}
