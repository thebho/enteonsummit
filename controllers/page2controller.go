package controllers

import (
  "github.com/astaxie/beego"
)

type Page2Controller struct {
  beego.Controller
}

func (this *Page2Controller) Get() {
  this.Data["names"] = []string{"Bob","Mary"}
  this.TplNames = "page2.html"
}