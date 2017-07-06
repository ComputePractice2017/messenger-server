package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

type Person struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"Password"`
}

var session *r.Session

func InitSesson() error {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address: "localhost",
	})
	return err
}

func GetPersons() ([]Person, error) {
	res, err := r.DB("Persons").Table("Persons").Run(session)
	if err != nil {
		return nil, err
	}

	var response []Person
	err = res.All(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
