package ui

import (
	"encoding/json"
	"fmt"
	tool "github.com/GeertJohan/go.rice"
	"github.com/go-vgo/robotgo"
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/rice"
	"github.com/sciter-sdk/go-sciter/window"
	"github.com/yale8848/xlsx2db/internal/db"
	"github.com/yale8848/xlsx2db/internal/file"
	"log"
	"time"
)

type Win interface {
	AddDefineFunction(w *window.Window)
	Create(htmlPath string)
}
type mainWin struct {
	driver    string
	source    string
	w         *window.Window
	tableInfo []db.TableInfo
}

type ret struct {
	Err  int         `json:"err"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

type TabFileData struct {
	Index    int    `json:"index"`
	Field    string `json:"field"`
	DefValue string `json:"defValue"`
}

func (this *ret) Fail(msg string) {
	this.Err = 1
	this.Msg = msg
}
func (this *ret) Ok(d interface{}) {
	this.Data = d
}
func (this *ret) String() string {
	b, _ := json.Marshal(&this)
	return string(b)
}
func NewMain() Win {
	return &mainWin{}
}
func (this *mainWin) call(fn string, args string) {
	this.w.Call(fn, sciter.NewValue(args))
}
func (this *mainWin) msg(msg string) {
	this.w.Call("msg", sciter.NewValue(msg))
}
func (this *mainWin) AddDefineFunction(w *window.Window) {

	w.DefineFunction("importData", func(args ...*sciter.Value) *sciter.Value {

		p := args[0].String()
		dbName := args[1].String()
		tbName := args[2].String()
		data := args[3].String()
		num := args[4].Int()

		value := make([]TabFileData, 0)
		_ = json.Unmarshal([]byte(data), &value)

		this.call("start", "")

		go func() {

			var err error

			defer func() {
				if err != nil {
					this.msg(err.Error())
				}
				this.call("finish", "")
			}()
			count := 0
			xf := file.NewFile()

			mfd := make(map[int]string)
			tbfInfo := make([]file.TableFileInfo, 0)

			getFieldType := func(field string) string {

				for _, v := range this.tableInfo {
					if v.ColumnName == field {
						return v.ColumnType
					}
				}
				return ""
			}
			for _, v := range value {
				mfd[v.Index] = v.Field
				tbfInfo = append(tbfInfo, file.TableFileInfo{Index: v.Index, FieldType: getFieldType(v.Field), DefValue: v.DefValue})
			}

			xd := db.NewDB()
			err = xd.Connect(this.driver, this.source)
			if err != nil {
				return
			}
			dmap := make([]map[string]interface{}, 0)
			startTime := time.Now()
			msgCount := func() {
				count += len(dmap)
				this.msg(fmt.Sprintf("已导入%d,用时%s", count, time.Now().Sub(startTime).String()))
			}
			err = xf.GetDataByIndex(p, tbfInfo, func(i int, strings map[int]string) error {

				mp := make(map[string]interface{})
				for k, v := range strings {
					mp[mfd[k]] = v
				}
				dmap = append(dmap, mp)
				if len(dmap) == num {
					err := xd.Insert(dbName, tbName, dmap)
					if err != nil {
						return err
					}
					msgCount()

					dmap = make([]map[string]interface{}, 0)
				}

				return nil
			})

			if err != nil {
				return
			}
			if len(dmap) > 0 {
				err = xd.Insert(dbName, tbName, dmap)
				if err != nil {
					return
				}
				msgCount()
			}

		}()

		return sciter.NewValue("")
	})
	w.DefineFunction("log", func(args ...*sciter.Value) *sciter.Value {
		fmt.Println(args[0].String())
		return sciter.NewValue("")
	})
	w.DefineFunction("getFileTitles", func(args ...*sciter.Value) *sciter.Value {
		p := args[0].String()
		index := args[1].Int()
		f := file.NewFile()
		t, e := f.GetFileTitlesBySheetIndex(p, index)
		r := ret{}
		if e != nil {
			r.Fail(e.Error())
		} else {
			r.Ok(t)
		}
		return sciter.NewValue(r.String())
	})

	w.DefineFunction("getSheetNames", func(args ...*sciter.Value) *sciter.Value {
		p := args[0].String()
		f := file.NewFile()
		t, e := f.GetSheetNames(p)
		r := ret{}
		if e != nil {
			r.Fail(e.Error())
		} else {
			r.Ok(t)
		}
		return sciter.NewValue(r.String())
	})

	w.DefineFunction("getDBName", func(args ...*sciter.Value) *sciter.Value {

		r := ret{}
		d := db.NewDB()
		err := d.Connect(this.driver, this.source)
		if err != nil {
			r.Fail(err.Error())
		} else {
			ts, err := d.GetDBNames()
			if err != nil {
				r.Fail(err.Error())
			} else {
				r.Ok(ts)
			}
		}
		b, _ := json.Marshal(&r)
		return sciter.NewValue(string(b))
	})
	w.DefineFunction("getTableName", func(args ...*sciter.Value) *sciter.Value {

		r := ret{}
		d := db.NewDB()
		err := d.Connect(this.driver, this.source)
		if err != nil {
			r.Fail(err.Error())
		} else {
			ts, err := d.GetTableNames(args[0].String())
			if err != nil {
				r.Fail(err.Error())
			} else {
				r.Ok(ts)
			}
		}
		return sciter.NewValue(r.String())

	})
	w.DefineFunction("setDBSource", func(args ...*sciter.Value) *sciter.Value {

		this.driver = args[0].String()
		this.source = args[1].String()
		return sciter.NewValue("")
	})
	w.DefineFunction("getTableInfos", func(args ...*sciter.Value) *sciter.Value {

		r := ret{}
		d := db.NewDB()
		err := d.Connect(this.driver, this.source)
		if err != nil {
			r.Fail(err.Error())
		} else {
			ts, err := d.GetTabScheme(args[0].String(), args[1].String())
			if err != nil {
				r.Fail(err.Error())
			} else {
				this.tableInfo = ts
				r.Ok(ts)
			}
		}
		return sciter.NewValue(r.String())
	})
}
func (this *mainWin) Create(htmlPath string) {

	tool.MustFindBox("assets")

	sw, sh := robotgo.GetScreenSize()

	ww := 500
	wh := 600

	left := int32((sw - ww) / 2)
	top := int32((sh - wh) / 2)

	w, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_RESIZEABLE|
		sciter.SW_MAIN|
		sciter.SW_ENABLE_DEBUG,
		&sciter.Rect{left, top, left + int32(ww), top + int32(wh)})
	if err != nil {
		log.Fatal(err)
	}
	rice.HandleDataLoad(w.Sciter)

	this.w = w
	this.AddDefineFunction(w)

	w.LoadFile("rice://assets/" + htmlPath)
	w.Show()
	w.Run()
}
