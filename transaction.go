package main

import (
	"encoding/json"
	"io"
	"math/big"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Transaction struct {
	Id            bson.ObjectId `json:"-" bson:"_id,omitempty"`
	Date          time.Time     `json:"date" bson:"date"`
	DebitAccount  string        `json:"debitAccount" bson:"debitAccount"`
	DebitAmount   *big.Rat      `json:"debitAmount" bson:"debitAmount"`
	CreditAccount string        `json:"creditAccount" bson:"creditAccount"`
	CreditAmount  *big.Rat      `json:"creditAmount" bson:"creditAmount"`
}

func (t *Transaction) FromJson(r io.Reader) error {
	return json.NewDecoder(r).Decode(&t)
}

func (t *Transaction) ToJson() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Transaction) Insert() error {
	return Insert("transactions", interface{}(t))
}
