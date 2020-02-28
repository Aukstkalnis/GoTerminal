package main

import (
	"fmt"
	"log"
	"runtime"

	sciter "github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/rice"
	"github.com/sciter-sdk/go-sciter/window"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func RunGUIApp() {
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
	if err = w.LoadFile("res/scapp.html"); err != nil {
		log.Fatal(err)
	}
	root, _ := w.GetRootElement()
	setCallBacks(root)
	w.Show()
	w.Run()
}

func setCallBacks(root *sciter.Element) {
	// \el, err := root.SelectById("port-list")
	// if err != nil {
	// 	return
	// }
	// el.OnClick(func() {
	// 	pop, err := el.Select("popup")

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	str := terminal.GetPortList()
	// 	pop[0].Clear()
	// 	for _, v := range str {
	// 		newEl, err := sciter.CreateElement("option", v)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		pop[0].Append(newEl)

	// 	}
	// 	fmt.Println("OnClick Event!")
	// })

	// output, err := root.SelectById("terminal-output")
	// if err != nil {
	// 	return
	// }
	// e, err = root.SelectById()
}
