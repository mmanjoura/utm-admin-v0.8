package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type UtmMsgs struct {
	ID   bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Date time.Time     `bson:"date" json:"date"`
	Uid  string        `bson:"uid" json:"uid"`
	Msg  string        `bson:"Msg" json:"Msg"`
}

type AmqpMessage struct {
	Device_uuid   string `bson:"device_uuid" json:"device_uuid"`
	Endpoint_uuid int    `bson:"endpoint_uuid" json:"endpoint_uuid"`
	Payload       []byte `bson:"payload json:"payload"`
}

type AmqpReceiveMessage struct {
	AmqpMessage
	Device_name string `bson:"device_name" json:"device_name`
	Id          string `bson:"id" json:"id"`
}

type AmqpResponseMessage struct {
	AmqpMessage
	Device_name string `bson:"device_name" json:"device_name`
	Command     string `bson:"command" json:"command"`
	Data        []byte `bson:"data" json:"data"`
}

type AmqpErrorMessage struct {
	AmqpMessage
	Queue   string `bson:"queue" json:"queue"`
	Message string `bson:"message" json:"message"`
	Reason  string `bson:"reason" json:"reason"`
}

func (u *UtmMsgs) Insert(db *mgo.Database, uuid string, msg string) {

	Msg := UtmMsgs{}
	utmColl := db.C("UtmMsgs")
	Msg.ID = bson.NewObjectId()
	Msg.Date = time.Now()
	Msg.Uid = uuid
	Msg.Msg = msg
	utmColl.Insert(&Msg)

}
