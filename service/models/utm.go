package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Uuid struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Uid     string        `bson:"uid" json:"uid"`
	RSSI    string        `bson:"rssi" json:"rssi"`
	Battery string        `bson:"battery" json:"battery"`
	Company string        `bson:"company" json:"company"`
}

type AmqpMessage struct {
	Device_uuid   string `bson:"device_uuid" json:"device_uuid"`
	Endpoint_uuid int    `bson:"endpoint_uuid" json:"endpoint_uuid"`
	Payload       []int  `bson:"payload json:"payload"`
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

var v = [...]Uuid{
	{Uid: "861f9e8c-5b8d-11e5-885d-feff819cdc9a", RSSI: "20", Battery: "6%", Company: "u-blox"},
	{Uid: "861fa274-5b8d-11e5-885d-feff819cdc9b", RSSI: "27", Battery: "60%", Company: "u-blox"},
	{Uid: "861fa3fa-5b8d-11e5-885d-feff819cdc9c", RSSI: "32", Battery: "67%", Company: "u-blox"},
	{Uid: "861fa53a-5b8d-11e5-885d-feff819cdc9d", RSSI: "62", Battery: "87%", Company: "vodafone"},
	{Uid: "861fa670-5b8d-11e5-885d-feff819cdc9e", RSSI: "72", Battery: "17%", Company: "vodafone"},
	{Uid: "861fa79c-5b8d-11e5-885d-feff819cdc9f", RSSI: "52", Battery: "78%", Company: "vodafone"},
}

func (u *Uuid) Insert(db *mgo.Database) {

	c := db.C("uuids")
	for _, e := range v {
		e.ID = bson.NewObjectId()
		c.Insert(&e)
	}

}
