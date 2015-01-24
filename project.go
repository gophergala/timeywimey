package main

import (
	"encoding/json"
	"io"

	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	Id        bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Members   []string      `json:"members" bson:"members"`
	Timelines []Calendar    `json:"timelines" bson:"timelines"`
}

func (p *Project) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&p)
}

func (p *Project) ToJson() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Project) Insert() error {
	return Insert("projects", interface{}(p))
}
