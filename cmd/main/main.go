package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/yale8848/xlsx2db/internal/ui"
)

func main() {
	sciter.SetOption(sciter.SCITER_SET_SCRIPT_RUNTIME_FEATURES, sciter.ALLOW_SYSINFO)

	ui.NewWin().Create("index.html")
}
