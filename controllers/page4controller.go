package controllers

import (
  "github.com/astaxie/beego"
)

type Page4Controller struct {
  beego.Controller
}

func (this *Page4Controller) Get() {
  this.TplNames = "page4.html"
}
