package ui

import (
	"fmt"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
	"path/filepath"
)

type Table struct {
	Winer
}

var driver string
var source string

func (this *Table) AddDefineFunction(w *window.Window) {
	this.Winer.AddDefineFunction(w)

	w.DefineFunction("getDBName", func(args ...*sciter.Value) *sciter.Value {
		return sciter.NewValue(args[0].Int() + 1)
	})
	w.DefineFunction("getTableName", func(args ...*sciter.Value) *sciter.Value {
		return sciter.NewValue(args[0].Int() + 1)
	})
	w.DefineFunction("setDBSource", func(args ...*sciter.Value) *sciter.Value {
		driver = args[0].String()
		source = args[1].String()

		fmt.Println(driver, source)
		return nil
	})
}
func (this *Table) Create(htmlPath string) {
	w, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_RESIZEABLE|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_ENABLE_DEBUG,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	this.AddDefineFunction(w)
	htmlPath, _ = filepath.Abs(htmlPath)
	w.LoadFile(htmlPath)
	w.Show()
	w.Run()
}
