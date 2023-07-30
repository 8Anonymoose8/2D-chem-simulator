package main

import (
	"fmt"
	"log"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK
	gtk.Init(nil)

	// Create a new top level window
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Sample Application")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Create a new button
	btn, err := gtk.ButtonNewWithLabel("Click me!")
	if err != nil {
		log.Fatal("Unable to create button:", err)
	}
	btn.Connect("clicked", func() {
		fmt.Println("Button clicked!")
	})

	// Add the button to the window
	win.Add(btn)

	// Set the default window size
	win.SetDefaultSize(300, 200)

	// Recursively show all widgets contained in this window
	win.ShowAll()

	// Begin the GTK main loop
	gtk.Main()
}
