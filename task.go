package main

import (
	"encoding/json"
	"io"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id                bson.ObjectId `json:"-" bson:"_id,omitempty"`
	RelatedTo         int           `json:"relatedTo" bson:"relatedTo"`
	Created           time.Time     `json:"created" bson:"created"`
	Author            string        `json:"author" bson:"author"`
	OriginalEstimate  int           `json:"originalEstimate" bson:"originalEstimate"`
	RemainingEstimate int           `json:"remainingEstimate" bson:"remainingEstimate"`
	CompletedTime     int           `json:"completedTime" bson:"completedTime"`
	Type              string        `json:"type" bson:"type"`
}

func (t *Task) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t)
}

func (t *Task) ToJson() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Task) Insert() error {
	return Insert("tasks", interface{}(t))
}
