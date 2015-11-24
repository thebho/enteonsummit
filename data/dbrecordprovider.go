package data

import (
  uuid "code.google.com/p/go-uuid/uuid"
  "github.com/couchbaselabs/go-couchbase"
  "fmt"
)

type DbRecord struct {
  Key string
  Name string
  Cost float32
}

type CouchStruct struct {
  Client *couchbase.Client
  Pool *couchbase.Pool
  Bucket *couchbase.Bucket
  RecordKeys []string
}

func GetListOfRecords() []DbRecord {
  listOfRecords := make([]DbRecord, 10)
  listOfRecords[0] = DbRecord{uuid.New(), "Shovel", 32.00}
  listOfRecords[1] = DbRecord{uuid.New(), "Lamp", 12.97}
  listOfRecords[2] = DbRecord{uuid.New(), "Food - Bread", 2.00}
  listOfRecords[3] = DbRecord{uuid.New(), "Backback", 25.02}
  listOfRecords[4] = DbRecord{uuid.New(), "Food - Water", 2.00}
  listOfRecords[5] = DbRecord{uuid.New(), "Rope", 10.00}
  listOfRecords[6] = DbRecord{uuid.New(), "Spoon", 1.00}
  listOfRecords[7] = DbRecord{uuid.New(), "Food - Meat", 8.00}
  listOfRecords[8] = DbRecord{uuid.New(), "Tent", 117.00}
  listOfRecords[9] = DbRecord{uuid.New(), "Food - Apple", 1.50}
  return listOfRecords
}

// connects a CouchStruct object to CouchBase db
func (cS *CouchStruct) ConnectToCouchbase() {
  cS.connectClient()
  cS.connectPool()
  cS.connectTask07Bucket()
}

func (cS * CouchStruct) CloseBucket() {
  cS.Bucket.Close()
}

// connects Couchstruct.client to client
func (c *CouchStruct) connectClient() {
  client,err := couchbase.Connect("http://localhost:8091")
  if err != nil {
    fmt.Printf("%v\n",err)
  } else {
    c.Client = &client
  }
}

// connects Couchstruct.client to pool
func (c *CouchStruct) connectPool() {
  pool, err := c.Client.GetPool("default")
  if err != nil {
    fmt.Printf("Error getting pool: %v\n", err)
  } else {
    c.Pool = &pool
  }
}

// connects Couchstruct.client to pool
func (c *CouchStruct) connectTask07Bucket() {
  bucket, err := c.Pool.GetBucket("task07")
  if err != nil {
    fmt.Printf("Error getting bucket: %v\n", err)
  } else {
    c.Bucket = bucket
  }
}

// adds a new DBRecord object to the DB and returns the string key of the new record
func (c *CouchStruct) AddDBRecordToDatabase(newRecord *DbRecord) {

  newRecord.Key = uuid.New()

  c.Bucket.Set(newRecord.Key,0,newRecord)
  c.RecordKeys = append(c.RecordKeys, newRecord.Key)
}

// deletes all the records in CouchbaseDB by keys in c's Keys struct
func (c *CouchStruct) DeleteAllRecords() error {
  opts := map[string]interface{}{"stale" : false }

  results, err := c.Bucket.View("dev_items", "items", opts)
  if err != nil {
    fmt.Printf("error: %v\n",err)
  }
  for _,v := range results.Rows {
    err = c.Bucket.Delete(v.ID)
    if err != nil {
      fmt.Printf("error: %v\n",err)
    }
  }

  return err
}

// deletes a record based on the given key
func (c *CouchStruct) DeleteRecord(key string) error {
  err := c.Bucket.Delete(key)
  if err != nil {
    fmt.Printf("%v\n",err)
  }

  return err
}

// adds a new record to the db
func (c *CouchStruct) AddRecord(dbRecord *DbRecord) error{
  dbRecord.Key = uuid.New()
  err := c.Bucket.Set(dbRecord.Key, 0, dbRecord)
  if err != nil {
    fmt.Printf("%v\n",err)
  }

  return err
}

// Pulls the viewName view from couchbucket and returns a slice of DbRecords from that view
func (c *CouchStruct)GetAllRecordsFromView(viewName string) ([]DbRecord, error) {
	var dbRecord DbRecord

	opts := map[string]interface{}{"stale" : false }

	results, err := c.Bucket.View("dev_items", viewName, opts)
	if err != nil {
		return nil, nil
	}
  dbRecords := make([]DbRecord, results.TotalRows)
  for i,v := range results.Rows {
    err = c.Bucket.Get(v.ID ,&dbRecord)
    if err != nil {
      return nil, err
    } else {
      //fmt.Println(dbRecord)
      dbRecords[i] = dbRecord
    }
  }
  return dbRecords,nil
}
