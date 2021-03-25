package main

import (
	"log"
	"os"

	"github.com/zserge/lorca"
)

var ui lorca.UI

func app() {
	var err error
	ui, err = lorca.New("", "", 960, 720)
	if err != nil {
		log.Fatalln("Error: ", err)
		return
	}
	defer ui.Close()
	ui.Load("http://localhost:1234")
	sigc := make(chan os.Signal)
	select {
	case <-sigc:
	case <-ui.Done():
	}
	log.Println("exiting...")
}
