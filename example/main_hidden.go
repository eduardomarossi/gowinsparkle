package main

import (
	"fmt"
	"github.com/tadvi/winc"
	"github.com/eduardomarossi/gowinsparkle"
)

func main() {
	err := gowinsparkle.WinSparkleLoad("..\\x64\\")
	if err != nil {
		fmt.Println(err)
	}
	gowinsparkle.WinSparkleSetAppDetails("eduardomarossi", "Hello GoWinSparkle", "1.0.0")
	gowinsparkle.WinSparkleSetAppCastUrl("http://localhost:8000/appcast.xml")
        gowinsparkle.WinSparkleSetHideUpdaterWindows(1)

	fmt.Printf("Auto-update %d Check Interval %d Last-check %d\n", gowinsparkle.WinSparkleGetAutomaticCheckForUpdates(),
		gowinsparkle.WinSparkleGetUpdateCheckInterval(),
		gowinsparkle.WinSparkleGetLastCheckTime())
	gowinsparkle.WinSparkleSetAutomaticCheckForUpdates(1)
	gowinsparkle.WinSparkleSetUpdateCheckInterval(10600)

	fmt.Printf("Auto-update %d Check Interval %d Last-check %d\n", gowinsparkle.WinSparkleGetAutomaticCheckForUpdates(),
		gowinsparkle.WinSparkleGetUpdateCheckInterval(),
		gowinsparkle.WinSparkleGetLastCheckTime())

	gowinsparkle.WinSparkleInit()
	gowinsparkle.WinSparkleCheckUpdateWithUIAndInstall()
        gowinsparkle.WinSparkleSetHideUpdaterWindows(1)

	winc.RunMainLoop() // Must call to start event loop.
}

