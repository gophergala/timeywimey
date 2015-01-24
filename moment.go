package main

import (
	"encoding/json"
	"io"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Moment struct {
	Id           bson.ObjectId `json:"-" bson:"_id,omitempty"`
	CalendarId   int           `json:"calendarId" bson:"calendarId"`
	ObjectType   string        `json:"objectType" bson:"objectType"`
	StartDate    time.Time     `json:"startDate" bson:"startDate"`
	EndDate      time.Time     `json:"endDate" bson:"endDate"`
	Repeating    bool          `json:"repeating" bson:"repeating"`
	Summary      string        `json:"summary" bson:"summary"`
	CalendarData string        `json:"calendarData" bson:"calendarData"`
	Uri          string        `json:"uri" bson:"uri"`
	LastModified time.Time     `json:"lastModified" bson:"lastModified"`
}

func (m *Moment) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&m)
}

func (m *Moment) ToJson() ([]byte, error) {
	return json.Marshal(m)
}

func (m *Moment) Insert() error {
	return Insert("moments", interface{}(m))
}
