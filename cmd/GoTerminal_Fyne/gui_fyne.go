package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

// func makeSerialListDropDown(_ fyne.Window) fyne.CanvasObject {
// 	portList := terminal.GetPortList()
// 	selectEntry := widget.NewSelectEntry(pwoerList)
// 	selectEntry.PlaceHolder = "Type or select"
// 	disabledCheck := widget.NewCheck("Disabled check", func(bool) {})
// 	disabledCheck.Disable()
// 	radio := widget.NewRadioGroup([]string{"Radio Item 1", "Radio Item 2"}, func(s string) { fmt.Println("selected", s) })
// 	radio.Horizontal = true
// 	disabledRadio := widget.NewRadioGroup([]string{"Disabled radio"}, func(string) {})
// 	disabledRadio.Disable()

// 	return container.NewVBox(
// 		widget.NewSelect([]string{"Option 1", "Option 2", "Option 3"}, func(s string) { fmt.Println("selected", s) }),
// 		selectEntry,
// 		widget.NewCheck("Check", func(on bool) { fmt.Println("checked", on) }),
// 		disabledCheck,
// 		radio,
// 		disabledRadio,
// 		widget.NewSlider(0, 100),
// 	)
// }

func makeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("About"),
	)
}

// func new_sidebar() fyne.CanvasObject {
// 	var obj fyne.CanvasObject
// 	return obj
// }

// func new_body() fyne.CanvasObject {
// 	var obj fyne.CanvasObject
// 	return obj
// }

// func new_statusbar() fyne.CanvasObject {
// 	var obj fyne.CanvasObject
// 	return obj
// }
/*
-------------------------------------------------------
[...] | COM11 x | COM9 x |
-------------------------------------------------------
[...] | [Open] [BaudRate v] [DataBits v] [Parity v] [StopBits v] [FlowControl v] [logToFile]
[...] | [Send] [Data To Send                              ]
[...] |------------------------------------------------
[...] |
      |
      |
      |
-------------------------------------------------------
[DTR] [RTS] [CTS] [DSR] [RI] [CD]
-------------------------------------------------------
*/
func run_gui_app() {
	a := app.New()
	w := a.NewWindow("GO Terminal")

	// myWindow.SetMainMenu(makeMenu(myApp, myWindow))

	w.SetContent(container.New(layout.NewVBoxLayout()))

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text3 := canvas.NewText("(right)", color.White)
	content := container.New(layout.NewHBoxLayout(), text1, text2, layout.NewSpacer(), text3)

	text4 := canvas.NewText("centered", color.White)
	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), text4, layout.NewSpacer())
	w.SetContent(container.New(layout.NewVBoxLayout(), content, centered))
	w.ShowAndRun()
}
