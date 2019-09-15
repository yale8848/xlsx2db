package ui

import (
	"fmt"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"log"
	"path/filepath"
	"strings"
)

type Win interface {
	AddDefineFunction(w *window.Window)
	Create(htmlPath string)
}
type Winer struct {
}

func NewWin() Win {
	return &Winer{}
}
func (this *Winer) AddDefineFunction(w *window.Window) {
	w.DefineFunction("setDBSource", func(args ...*sciter.Value) *sciter.Value {
		driver = args[0].String()
		source = args[1].String()

		fmt.Println(driver, source)
		return nil
	})
	w.DefineFunction("CreateWindow", func(args ...*sciter.Value) *sciter.Value {
		ht := args[0].String()
		ht = ht[0:strings.Index(ht, ".")]
		var win Win
		switch ht {
		case "table":
			win = &Table{}
			break
		}
		win.Create(args[0].String())
		return nil
	})
}
func (this *Winer) Create(htmlPath string) {
	w, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_RESIZEABLE|
		sciter.SW_POPUP|
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
