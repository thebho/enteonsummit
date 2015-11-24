package rest

import (
	"github.com/astaxie/beego"
	"encoding/json"
)

type EchoController3 struct {
	beego.Controller
}

type EchoAddress struct {
	Add1,Add2,City,State,Zip string
}

func (this *EchoController3) Post() {
	var echoAddress EchoAddress
	var echoValue EchoValue

	req := this.Ctx.Request
	p := make([]byte, req.ContentLength)

	this.Ctx.Request.Body.Read(p)

	if err := json.Unmarshal(p, &echoAddress); err != nil {
    this.CustomAbort(500, "Unable to unmarshal value")
	} else {

		// Try adding each attribute field to the return string
		concatString(echoAddress.Add1, &echoValue)
		concatString(echoAddress.Add2, &echoValue)
		concatString(echoAddress.City, &echoValue)
		concatString(echoAddress.State, &echoValue)
		concatString(echoAddress.Zip, &echoValue)

		this.Data["json"] = echoValue

		this.ServeFormatted()
	}
}

func concatString(att string, echo *EchoValue) {
	if att != "" {
		echo.Val += string(att[0])
	}
}
