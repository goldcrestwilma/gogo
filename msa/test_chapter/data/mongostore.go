package data

import "labix.org/v2/mgo"

type MongoStore struct {
	session *mgo.Session
}

func NewMongoStore(connection string) (*MongoStore, error) {
	session, err := mgo.Dial(connection)
	if err != nil {
		return nil, err
	}

	return &MongoStore{session: session}, nil
}

func (m *MongoStore) Search(name string) []Kitten {
	s := m.session.Clone()
	defer s.Close()

	var results []Kitten
	c := s.DB("kittenserver").C("kittens")
	err := c.Find(Kitten{Name: name}).All(&results)
	if err != nil {
		return nil
	}

	return results
}

func (m *MongoStore) DeleteAllKittens() {
	s := m.session.Clone()
	defer s.Close()

	s.DB("kittenserver").C("kittens").DropCollection()
}

// InsertKittens inserts a slice of kittens into the datastore
func (m *MongoStore) InsertKittens(kittens []Kitten) {

	s := m.session.Clone()
	defer s.Close()

	s.DB("kittenserver").C("kittens").Insert(kittens)
}
