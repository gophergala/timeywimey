package main

import (
	"encoding/json"
	"io"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type TimeEntry struct {
	Id           bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Date         time.Time     `json:"date" bson:"date"`
	Duration     int           `json:"duration" bson:"duration"`
	Summary      string        `json:"summary" bson:"summary"`
	Author       string        `json:"author" bson:"author"`
	Created      time.Time     `json:"created" bson:"created"`
	LastModified time.Time     `json:"lastModified" bson:"lastModified"`
}

func (t *TimeEntry) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t)
}

func (t *TimeEntry) ToJson() ([]byte, error) {
	return json.Marshal(t)
}

func (t *TimeEntry) Insert() error {
	return Insert("timesheet", interface{}(t))
}
