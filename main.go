package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", hnd)
	log.Println("Starting webserver...")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func hnd(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hi!")

}
