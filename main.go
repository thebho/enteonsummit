package main

import (
	"os"
	"path"

	"github.com/astaxie/beego"
)

func init() {
	beego.TemplateLeft = "<<<"
	beego.TemplateRight = ">>>"
	beego.EnableAdmin = true
}

var httpport = 8080
var adminport = 8082

func initialize_tmpdir(tmpdir string) {
	_ = os.RemoveAll(tmpdir)
	_ = os.Mkdir(tmpdir, 0700)
}

func main() {
	beego.HttpPort = httpport
	beego.EnableAdmin = true
	beego.AdminHttpPort = adminport
	me := path.Base(os.Args[0])
	beego.AppName = me
	tmpdir := "./tmp"
	initialize_tmpdir(tmpdir)

	beego.SessionOn = true
	beego.SessionProvider = "file"
	beego.SessionSavePath = tmpdir
	beego.SessionName = me
	beego.Run()
}
