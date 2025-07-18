package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	eq "github.com/regularblueberry/EquiBook/application"
)

func RunApp(appdata *eq.App){

	a := app.New()

	// Create a new window for the application
	w := a.NewWindow("My Finance App")

	// Create a simple label widget
	greeting := widget.NewLabel("Welcome to My Finance App!")

	// Create a button
	button := widget.NewButton("Get Started", func() {
		greeting.SetText("Let's manage your money!")
	})

	w.SetContent(container.NewVBox(
		greeting,
		button,
		widget.NewLabel("More features coming soon..."),
	))

	// Show the window and run the application event loop
	w.ShowAndRun()
}

func RunMainLoop(appdata *eq.App){


}
