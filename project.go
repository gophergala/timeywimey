package main

import (
	"encoding/json"
	"io"

	"gopkg.in/mgo.v2"
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

func (p *Project) Get(projectName, username string) error {
	mongoUri := connectionString()
	sess, err := mgo.Dial(mongoUri)
	if err != nil {
		return err
	}
	defer sess.Close()

	sess.SetSafe(&mgo.Safe{})
	c := sess.DB("timeywimey").C("projects")
	err = c.Find(Project{Name: projectName, Members: []string{username}}).One(&p)
	if err != nil {
		return err
	}

	return nil
}

func GetByUsername(username string) ([]Project, error) {
	mongoUri := connectionString()
	sess, err := mgo.Dial(mongoUri)
	if err != nil {
		return nil, err
	}
	defer sess.Close()

	var projects []Project
	sess.SetSafe(&mgo.Safe{})
	c := sess.DB("timeywimey").C("projects")
	err = c.Find(Project{Members: []string{username}}).All(&projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (p *Project) Insert() error {
	return Insert("projects", interface{}(p))
}
