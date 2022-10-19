package main

import (
	"Webserver/requests"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	if len(os.Args) > 1 {
		if os.Args[1] == "-n" {
			for {
				print(".")
				time.Sleep(1 * time.Second)
			}
		}
	}

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
