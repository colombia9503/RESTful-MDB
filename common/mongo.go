package common

import (
	"gopkg.in/mgo.v2"
)

type mongo struct {
	Tasks *mgo.Collection
}

var DB *mongo

//Iniciar session con mongoDB
func connectDB() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	DB = &mongo{session.DB("taskdb").C("tasks")}
}
