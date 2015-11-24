package controllers

import (
  "github.com/astaxie/beego"
  records "edu/Task_06_Beego/task06beego/data"
)

type Page3Controller struct {
  beego.Controller
}

func (this *Page3Controller) Get() {
	this.Data["records"] = records.GetListOfRecords()
  this.TplNames = "page3.html"
}