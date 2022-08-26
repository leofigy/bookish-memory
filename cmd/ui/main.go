package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	fs "gihub.com/leofigy/bookish-memory/fs"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Welcome to the bookish-memory")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText(getData())
		}),
	))

	w.ShowAndRun()
}

func getData() (msg string) {
	inf, err := fs.CalcSpace("")
	if err != nil {
		return fmt.Sprintf("error %s", err)
	}

	return fmt.Sprintf("default used space : %d", inf.GetSpace().Used)
}
