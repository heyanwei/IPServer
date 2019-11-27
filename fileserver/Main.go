package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Start file service....")
	http.Handle("/", http.FileServer(http.Dir("file/.")))
	err := http.ListenAndServe(":12334", nil)
	if err != nil {
		log.Println("Start Failed...")
		return
	}
	log.Println("Start file service success")
}
