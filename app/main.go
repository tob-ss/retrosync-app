package main

import (
	"fmt"
	"main/icon"
	"time"

	"github.com/slytomcat/systray"
)

func main() {
	onExit := func() {
		now := time.Now()
		fmt.Println("Exit at", now.String())
	}

	systray.Run(onReady, onExit)
}

func addQuitItem() {
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	mQuit.Enable()
	go func() {
		<-mQuit.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
	systray.AddSeparator()
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Retrosync")
	systray.SetTooltip("Retrosync is active!")
	addQuitItem()

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(icon.Data, icon.Data)
		systray.SetTitle("Retrosync")
		systray.SetTooltip("Retrosync is active!")
		// Sets the icon of a menu item. Only available on Mac.

		startScan()
	}()
}
