package store

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pkg/errors"
)

type ClientType struct {
	DB DBType
}

type DBType struct {
	Collection CollectionType `json:"collection"`
	Session    SessionType    `json:"session,omitempty"`
}

type SessionType struct {
	Username string `json:"username"`
}

type CollectionType struct {
	Meetings Meetings `json:"meetings"`
}

// GetClient Get a Client instance
func GetClient() (*ClientType, error) {
	client := ClientType{}
	buff, err := ioutil.ReadFile("/var/agenda/data.json")
	if err != nil {
		return nil, errors.New("Opening data file failed")
	}
	if err := json.Unmarshal(buff, &client); err != nil {
		return nil, err
	}
	return &client, nil
}

// Dump the db data stores in Client instance
func (cl *ClientType) Dump() error {
	buff, err := json.Marshal(cl.DB)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("/var/agenda/data.json", buff, 0); err != nil {
		return err
	}
	return nil
}

func (cl *ClientType) Commit() error {
	_, err := json.Marshal(cl.DB)
	if err != nil {
		return err
	}
	return nil
}
