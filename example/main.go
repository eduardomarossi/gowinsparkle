package main

import (
	"fmt"
	"github.com/tadvi/winc"
	"github.com/eduardomarossi/gowinsparkle"
)

func main() {
	mainWindow := winc.NewForm(nil)
	mainWindow.SetSize(400, 300)  // (width, height)
	mainWindow.SetText("Hello World Demo")
	mainWindow.Center()
	mainWindow.Show()

	err := gowinsparkle.WinSparkleLoad("..\\x86\\")
	if err != nil {
		fmt.Println(err)
	}
	gowinsparkle.WinSparkleSetAppDetails("eduardomarossi", "Hello GoWinSparkle", "1.0.0")
	gowinsparkle.WinSparkleSetAppCastUrl("http://localhost:8000/appcast.xml")

	fmt.Printf("Auto-update %d Check Interval %d Last-check %d\n", gowinsparkle.WinSparkleGetAutomaticCheckForUpdates(),
		gowinsparkle.WinSparkleGetUpdateCheckInterval(),
		gowinsparkle.WinSparkleGetLastCheckTime())
	gowinsparkle.WinSparkleSetAutomaticCheckForUpdates(1)
	gowinsparkle.WinSparkleSetUpdateCheckInterval(10600)

	fmt.Printf("Auto-update %d Check Interval %d Last-check %d\n", gowinsparkle.WinSparkleGetAutomaticCheckForUpdates(),
		gowinsparkle.WinSparkleGetUpdateCheckInterval(),
		gowinsparkle.WinSparkleGetLastCheckTime())

	gowinsparkle.WinSparkleInit()
	gowinsparkle.WinSparkleCheckUpdateWithUI()
	mainWindow.OnClose().Bind(wndOnClose)

	winc.RunMainLoop() // Must call to start event loop.
}

func wndOnClose(arg *winc.Event) {
	winc.Exit()
}
