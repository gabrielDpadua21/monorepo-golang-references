package database

import "gopkg.in/mgo.v2/bson"

type ODM struct {
}

type FuncMap map[string]interface{}

func NewODM(config ConfigDB) (*ODM, error) {

	ODM := new(ODM)
	err := initDB(config)
	if err != nil {
		return nil, err
	}
	return ODM, nil
}

func (d *ODM) NewSession() *Session {
	session := &Session{}
	session.Init()
	return session
}

func (d *ODM) Collection(col string) *Session {
	session := d.NewSession()
	session.collection = col
	return session
}

func (d *ODM) Fields(fields ...string) *Session {
	session := d.NewSession()
	session.fields = fields
	return session
}

func (d *ODM) Sort(fields ...string) *Session {
	session := d.NewSession()
	session.sort = fields
	return session
}

func (d *ODM) Limit(i int) *Session {
	session := d.NewSession()
	session.limit = i
	return session
}

func (d *ODM) Offset(i int) *Session {
	session := d.NewSession()
	session.offset = i
	return session
}

func (d *ODM) Where(query bson.M) *Session {
	session := d.NewSession()
	session.where = query
	return session
}

func (d *ODM) WhereArray(query []bson.M) *Session {
	session := d.NewSession()
	session.whereArray = query
	return session
}

// Aggregate function to mongo db and return record as map[string]interface{}
func (d *ODM) AggregateOne(i interface{}) error {
	session := d.NewSession()
	return session.AggregateOne(i)
}

// Aggregate function to mongo db and return records as []map[string]interface{}
func (d *ODM) AggregateAll(i interface{}) error {
	session := d.NewSession()
	return session.AggregateAll(i)
}

// Get Query a raw sql and return records as []map[string][]byte
func (d *ODM) Find(i interface{}) error {
	session := d.NewSession()
	return session.Find(i)
}

func (d *ODM) FindOne(i interface{}) error {
	session := d.NewSession()
	return session.FindOne(i)
}

func (d *ODM) FindByID(id string, i interface{}) error {
	session := d.NewSession()
	return session.FindByID(id, i)

}

func (d *ODM) Insert(doc Document) error {
	session := d.NewSession()
	return session.Insert(doc)
}

func (d *ODM) Update(doc Document) error {
	session := d.NewSession()
	return session.Update(doc)
}

func (d *ODM) Delete() error {
	session := d.NewSession()
	return session.Delete()
}

func (d *ODM) DeleteByID(id string) error {
	session := d.NewSession()
	return session.DeleteByID(id)
}

func (d *ODM) Count() (int, error) {
	session := d.NewSession()
	return session.Count()
}

func (d *ODM) CountWhere() (int, error) {
	session := d.NewSession()
	return session.Count()
}

func (d *ODM) Close() {
	mongoSession.Close()
}
