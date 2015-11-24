package rest

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type EchoController struct {
	beego.Controller
}

type EchoValue struct {
	Val string
}

type EchoError struct {
	Err string
}


func (this *EchoController) Get() {

	var echoValue EchoValue

	echoValue.Val = this.Ctx.Input.Params[":val"]

	this.Data["json"] = echoValue

	this.ServeFormatted()
}

func (this *EchoController) Post() {

	var echoValue EchoValue

	req := this.Ctx.Request
	p := make([]byte, req.ContentLength)

	this.Ctx.Request.Body.Read(p)

	if err := json.Unmarshal(p, &echoValue); err != nil {
    this.CustomAbort(500, "Unable to unmarshal value")
	} else {

		this.Data["json"] = echoValue

		this.ServeFormatted()
	}
}
