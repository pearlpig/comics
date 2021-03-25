package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func service() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	fmt.Println("Listening port 1234 ...")

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("view/index.html")
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	w.Write(content)
}
