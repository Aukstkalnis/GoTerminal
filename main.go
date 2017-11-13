package main

import (
	"fmt"
	"log"

	sciter "github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/rice"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	fmt.Println("Hello")
	r := sciter.NewRect(100, 500, 800, 800)
	//w, err := window.New(sciter.DefaultWindowCreateFlag, sciter.DefaultRect)
	w, err := window.New(sciter.DefaultWindowCreateFlag, r)
	if err != nil {
		log.Fatal(err)
	}
	rice.HandleDataLoad(w.Sciter)
	ok := w.SetOption(sciter.SCITER_SET_DEBUG_MODE, 1)
	if !ok {
		log.Println("set debug mode failed")
	}
	if err = w.LoadFile("scapp.html"); err != nil {
		log.Fatal(err)
	}
	w.Show()
	w.Run()
}
