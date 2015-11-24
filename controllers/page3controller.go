package controllers

import (
  "github.com/astaxie/beego"
)

type Page3Controller struct {
  beego.Controller
}

func (this *Page3Controller) Get() {
  this.TplNames = "page3.html"
}