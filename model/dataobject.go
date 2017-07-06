package model

import (
	r "gopkg.in/gorethink/gorethink.v3"
)

type Person struct {
	ID       string `json:"id",gorethink:"id"`
	Login    string `json:"login",gorethink:"login"`
	Password string `json:"Password",gorethink:"Password"`
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

func NewPerson(p Person) (Person, error) {
	res, err := r.UUID().Run(session)
	if err != nil {
		return p, err
	}

	var UUID string
	err = res.One(&UUID)
	if err != nil {
		return p, err
	}

	p.ID = UUID

	res, err = r.DB("Persons").Table("Persons").Insert(p).Run(session)
	if err != nil {
		return p, err
	}

	return p, nil
}
