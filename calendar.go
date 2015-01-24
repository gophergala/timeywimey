package main

import (
	"encoding/json"
	"io"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Calendar struct {
	Id           bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Name         string        `json:"name" bson:"name"`
	Members      []string      `json:"members" bson:"members"`
	Moments      []Moment      `json:"moments" bson:"moments"`
	Created      time.Time     `json:"created" bson:"created"`
	LastModified time.Time     `json:"lastModified" bson:"lastModified"`
}

func (c *Calendar) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&c)
}

func (c *Calendar) ToJson() ([]byte, error) {
	return json.Marshal(c)
}

func (c *Calendar) Insert() error {
	return Insert("calendars", interface{}(c))
}
