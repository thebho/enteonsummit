package data

import (
  uuid "code.google.com/p/go-uuid/uuid"
)

type DbRecord struct {
  Key string
  Name string
  Cost float32
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