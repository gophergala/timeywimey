package main

import (
	"encoding/json"
	"io"
	"math/big"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Invoice struct {
	Id           bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Entries      []TimeEntry   `json:"entries" bson:"entries"`
	StartDate    time.Time     `json:"startDate" bson:"startDate"`
	EndDate      time.Time     `json:"endDate" bson:"endDate"`
	Taxes        *big.Rat      `json:"taxes" bson:"taxes"`
	Author       string        `json:"author" bson:"author"`
	Created      time.Time     `json:"created" bson:"created"`
	LastModified time.Time     `json:"lastModified" bson:"lastModified"`
}

func (i *Invoice) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&i)
}

func (i *Invoice) ToJson() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Invoice) Insert() error {
	return Insert("invoices", interface{}(i))
}
