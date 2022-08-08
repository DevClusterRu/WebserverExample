package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", hnd)
	//http.HandleFunc("/export", csvExport.ExportListener)
	log.Println("Starting webserver...")
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func hnd(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello123123 !")
}
