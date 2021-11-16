package database

import (
	"fmt"

	valid "gopkg.in/asaskevich/govalidator.v4"
	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	collection string
	limit      int
	offset     int
	where      bson.M
	whereArray []bson.M
	order      string
	sort       []string
	params     []interface{}
	fields     []string
	pk         string
	join       string
	groupBy    string
}

func (s *Session) Init() {
	s.limit = 0
	s.offset = 0
	s.order = ""
}

func (s *Session) Where(query bson.M) *Session {
	s.where = query
	return s
}

func (s *Session) WhereArray(query []bson.M) *Session {
	s.whereArray = query
	return s
}

func (s *Session) Fields(fields ...string) *Session {
	s.fields = fields
	return s
}

func (s *Session) Sort(fields ...string) *Session {
	s.sort = fields
	return s
}

func (s *Session) Limit(i int) *Session {
	s.limit = i
	return s
}

func (s *Session) Offset(i int) *Session {
	s.offset = i
	return s
}

func (s *Session) Find( i interface{}) error {
	if s.collection == "" {
		return fmt.Errorf("need to set a collection")
	}
	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)
	mongoQuery := c.Find(s.where)

	if s.limit > 0 {
		mongoQuery = mongoQuery.Limit(s.limit)
	}

	if s.offset > 0 {
		mongoQuery = mongoQuery.Skip(s.offset)
	}

	if len(s.fields) > 0 {
		mapFields := make(map[string]interface{})
		for _, s := range s.fields {
			mapFields[s] = 1
		}
		mongoQuery = mongoQuery.Select(mapFields)
	}

	if len(s.sort) > 0 {
		mongoQuery = mongoQuery.Sort(s.sort...)
	}
	return mongoQuery.All(i)

}

func (s *Session) FindOne( i interface{}) error {
	if s.collection == "" {
		return fmt.Errorf("need to set a collection")
	}
	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)
	mongoQuery := c.Find(s.where)

	if s.limit > 0 {
		mongoQuery = mongoQuery.Limit(s.limit)
	}

	if s.offset > 0 {
		mongoQuery = mongoQuery.Skip(s.offset)
	}

	if len(s.fields) > 0 {
		mapFields := make(map[string]interface{})
		for _, s := range s.fields {
			mapFields[s] = 1
		}
		mongoQuery = mongoQuery.Select(mapFields)
	}

	if len(s.sort) > 0 {
		mongoQuery = mongoQuery.Sort(s.sort...)
	}
	return mongoQuery.One(i)

}

func (s *Session) AggregateOne(i interface{}) error {
	if s.collection == "" {
		return fmt.Errorf("need to set a tablename")
	}

	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)

	return c.Pipe(s.whereArray).One(i)

}

func (s *Session) AggregateAll(i interface{}) error {
	if s.collection == "" {
		return fmt.Errorf("need to set a collection")
	}

	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)

	return c.Pipe(s.whereArray).All(i)
}

func (s *Session) FindByID(id string, i interface{}) error {
	if s.collection == "" {
		return fmt.Errorf("need to set a tablename")
	}

	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)

	if !bson.IsObjectIdHex(id) {
		return fmt.Errorf("Mongo ObjectID is invalid")
	}
	err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(i)
	return err
}

func (s *Session) Insert(i interface{}) error {
	if s.collection == "" {
		return fmt.Errorf("need to set a tablename")
	}

	result, err := valid.ValidateStruct(i)
	if result == false {
		return err
	}

	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)
	//doc.SetId(bson.NewObjectId())
	//now := time.Now()
	//doc.SetCreated(now)
	//doc.SetUpdated(now)
	return c.Insert(i)
}

func (s *Session) Update(doc interface{}) error {
	if s.collection == "" {
		return fmt.Errorf("need to set a tablename")
	}
	result, err := valid.ValidateStruct(doc)
	if result == false {
		return err
	}

	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)
	//oID := doc.GetId()
	//if oID.Valid() == false {
	//	return fmt.Errorf("Invalid oID")
	//}
	//doc.SetUpdated(time.Now())

	if len(s.where)  == 0 {
		return fmt.Errorf("update query could not be empty")
	}

	return c.Update(s.where, doc)
}

func (s *Session) Delete() error {
	if s.collection == "" {
		return fmt.Errorf("need to set a tablename")
	}
	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)

	if len(s.where)  == 0 {
		return fmt.Errorf("update query could not be empty")
	}
	return c.Remove(s.where)
}

func (s *Session) DeleteByID(id string) error {
	if s.collection == "" {
		return fmt.Errorf("need to set a tablename")
	}
	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)
	oID := bson.ObjectIdHex(id)
	if oID.Valid() == false {
		return fmt.Errorf("Invalid oID")
	}
	return c.RemoveId(oID)
}

func (s *Session) Count() (int, error) {
	if s.collection == "" {
		return 0, fmt.Errorf("need to set a collection")
	}
	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)
	return c.Count()
}

func (s *Session) CountWhere() (int, error) {
	if s.collection == "" {
		return 0, fmt.Errorf("need to set a collection")
	}
	ss := mongoSession.Copy()
	defer ss.Close()
	c := ss.DB(dbName).C(s.collection)
	mongoQuery := c.Find(s.where)

	return mongoQuery.Count()
}
