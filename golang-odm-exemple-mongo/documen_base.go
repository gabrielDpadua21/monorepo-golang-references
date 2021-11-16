package database

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)
type Document interface {
	GetId() bson.ObjectId
	SetId(bson.ObjectId)
	GetCreated() time.Time
	SetCreated(time.Time)
	GetUpdated() time.Time
	SetUpdated(time.Time)
}

//DocumentBase ...
type DocumentBase struct {
	ID      bson.ObjectId `json:"_id" bson:"_id"`
	Created time.Time     `json:"Created" bson:"Created" validate:"required"`
	Updated time.Time     `json:"Updated" bson:"Updated" validate:"required"`
}

type DocumentBase2 struct {
	ID      bson.ObjectId `json:"id" bson:"_id"`
	Created time.Time     `json:"Created" bson:"Created" validate:"required"`
	Updated time.Time     `json:"Updated" bson:"Updated" validate:"required"`
}

func (d *DocumentBase) GetId()(bson.ObjectId){
	return d.ID
}

func (d *DocumentBase) SetId(id bson.ObjectId){
	d.ID = id
}

func (d *DocumentBase) GetCreated()(time.Time){
	return d.Created
}

func (d *DocumentBase) SetCreated(created time.Time){
	d.Created = created
}

func (d *DocumentBase) GetUpdated()(time.Time){
	return d.Updated
}

func (d *DocumentBase) SetUpdated(updated time.Time){
	d.Updated = updated
}
