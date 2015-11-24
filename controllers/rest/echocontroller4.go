package rest

import (
	db "edu/Task_07_Angular/task07angular/data"
	"github.com/astaxie/beego"
	"encoding/json"
	"fmt"
)

type EchoController4 struct {
	beego.Controller
}

type EchoArray struct{
	Val []db.DbRecord
}


func (this *EchoController4) Get() {

	var echoValue string

	fmt.Println("Get called")

	echoValue = this.Ctx.Input.Params[":val"]

	var couchStruct db.CouchStruct

	couchStruct.ConnectToCouchbase()
	defer couchStruct.CloseBucket()

	if echoValue == "" {
		records, _ := couchStruct.GetAllRecordsFromView("items")

		var echoArray EchoArray
		echoArray.Val = records
		this.Data["json"] = echoArray
	} else {
		// this will get more detail on a specific record
	}

	this.ServeFormatted()
}

// Deletes all or one record from the Couchbase db
func (this *EchoController4) Delete() {
	fmt.Println("Delete called")

	var (
		val string
		couchStruct db.CouchStruct
		echoError EchoError
		err error
	)

	couchStruct.ConnectToCouchbase()
	defer couchStruct.CloseBucket()

	val = this.Ctx.Input.Params[":val"]

	// if val is nil, delete all records, else delete only the provided key
	if val == "" {
		fmt.Println("Deleting all")
		err = couchStruct.DeleteAllRecords()
	} else {
		fmt.Printf("Deleting: %v\n",val)
		err = couchStruct.DeleteRecord(val)
	}

	if err != nil {
		echoError.Err = "delete record error"
	}

		// Returns error or "" if the delete record was successful
		this.Data["json"] = echoError

		this.ServeFormatted()
}


// Adds new record to couchbase db
func (this *EchoController4) Post() {
	fmt.Println("called")

	var newDbRecord db.DbRecord
	var couchStruct db.CouchStruct

	couchStruct.ConnectToCouchbase()
	defer couchStruct.CloseBucket()

	req := this.Ctx.Request
	p := make([]byte, req.ContentLength)

	this.Ctx.Request.Body.Read(p)

	if err := json.Unmarshal(p, &newDbRecord); err != nil {
    this.CustomAbort(500, "Unable to unmarshal value")
	} else {

		fmt.Printf("New Record: %v,%v,%v\n",newDbRecord.Key,newDbRecord.Name,newDbRecord.Cost)

		couchStruct.AddRecord(&newDbRecord)


		this.Data["json"] = newDbRecord

		this.ServeFormatted()
	}
}
