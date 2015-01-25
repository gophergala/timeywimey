package main

import (
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
