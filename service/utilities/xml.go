package utilities

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UtmXml struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Date    time.Time     `bson:"date" json:"date"`
	Uid     string        `bson:"uid" json:"uid"`
	XmlData string        `bson:"XmlData" json:"XmlData"`
}

func (u *UtmXml) Insert(db *mgo.Database, xmlData string, uuid string) {

	var utmXml = UtmXml{}

	xml := db.C("UtmsXml")
	u.ID = bson.NewObjectId()
	u.Date = time.Now()
	u.Uid = uuid
	u.XmlData = xmlData
	xml.Insert(&utmXml)

}

func (u *UtmXml) Get(db *mgo.Database, uuid string) error {

	return db.C("UtmsXml").FindId(bson.ObjectIdHex(uuid)).One(&u)

}
