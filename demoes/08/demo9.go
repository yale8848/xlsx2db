package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func defFunc(w *window.Window) {
	//注册dump函数方便在tis脚本中打印数据
	w.DefineFunction("dump", func(args ...*sciter.Value) *sciter.Value {
		for _, v := range args {
			fmt.Print(v.String() + " ")
		}
		fmt.Println()
		return sciter.NullValue()
	})
	//处理文件移动
	w.DefineFunction("moveFile", func(args ...*sciter.Value) *sciter.Value {
		file := args[0].String()
		folder := args[1].String()

		//去掉路径左边的file://
		file = strings.TrimLeft(file, "file://")
		//获取文件名
		fileName := filepath.Base(file)
		//读取文件数据
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return sciter.NewValue(err.Error())
		}
		//写入文件数据
		err = ioutil.WriteFile(filepath.Join(folder, fileName), data, os.ModePerm)
		if err == nil {
			//删除原文件
			os.Remove(file)
			return sciter.NewValue("移动成功")
		} else {
			return sciter.NewValue("移动失败")
		}
	})
}

func main() {

	sciter.SetOption(sciter.SCITER_SET_SCRIPT_RUNTIME_FEATURES, sciter.ALLOW_SYSINFO)
	sciter.SetOption(sciter.SCITER_SET_DEBUG_MODE, 1)
	w, err := window.New(sciter.SW_TITLEBAR|
		sciter.SW_RESIZEABLE|
		sciter.SW_CONTROLS|
		sciter.SW_MAIN|
		sciter.SW_ENABLE_DEBUG,
		nil)
	if err != nil {
		log.Fatal(err)
	}
	p, _ := filepath.Abs("demo9.html")
	w.LoadFile(p)
	w.SetTitle("view对象学习")
	defFunc(w)
	w.Show()
	w.Run()
}
