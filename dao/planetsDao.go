package dao

import (
	"log"

	. "github.com/star-wars-go/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//PlanetsDAO to access planets in database
type PlanetsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "planets"
)

//Connect database function
func (m *PlanetsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//GetAll planets from DB
func (m *PlanetsDAO) GetAll() ([]Planet, error) {
	var planets []Planet
	err := db.C(COLLECTION).Find(bson.M{}).All(&planets)
	return planets, err
}

//GetByID planet from DB by id
func (m *PlanetsDAO) GetByID(id string) (Planet, error) {
	var planet Planet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&planet)
	return planet, err
}

//Create a planet in DB
func (m *PlanetsDAO) Create(planet Planet) (Planet, error) {
	err := db.C(COLLECTION).Insert(&planet)
	return planet, err
}

//Delete a planet from DB
func (m *PlanetsDAO) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

//Update a planet in DB
func (m *PlanetsDAO) Update(id string, planet Planet) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &planet)
	return err
}
