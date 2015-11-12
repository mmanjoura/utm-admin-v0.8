// This needs to get factored out into a separate package.

package server

import (
	"encoding/json"
	"fmt"
	"log"

	"time"

	"github.com/streadway/amqp"

	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

var UuidMap = make(map[string]*DisplayRow)
var UuidSlice = make([]*DisplayRow, len(UuidMap))

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

type Queue struct {
	conn     *amqp.Connection
	quit     chan interface{}
	Msgs     chan interface{}
	Downlink chan AmqpMessage
}

var TotalMsgs uint64
var TotalBytes uint64

var Row = &DisplayRow{}
var RowsList = make([]*DisplayRow, 1)

func consumeQueue(channel *amqp.Channel, chanName string) (<-chan amqp.Delivery, error) {
	msgChan, err := channel.Consume(
		chanName, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	if err != nil {
		return nil, fmt.Errorf("CONSUMING '%s' CHANNEL: %s", chanName, err.Error())
	}
	return msgChan, nil
}

func OpenQueue(username, amqpAddress string) (*Queue, error) {
	q := Queue{}
	q.quit = make(chan interface{})
	q.Msgs = make(chan interface{})
	q.Downlink = make(chan AmqpMessage)

	conn, err := amqp.Dial(amqpAddress)
	if err != nil {
		return nil, fmt.Errorf("Connecting to RabbitMQ: %s", err.Error)
	}
	q.conn = conn

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	receiveChan, err := consumeQueue(channel, username+".receive")
	if err != nil {
		return nil, err
	}
	responseChan, err := consumeQueue(channel, username+".response")
	if err != nil {
		return nil, err
	}
	errorChan, err := consumeQueue(channel, username+".error")
	if err != nil {
		return nil, err
	}

	//Open DB
	//amqpSession, err := mgo.Dial("127.0.0.1:27017")

	// rec_Collection := amqpSession.DB("utm-db").C("receivemessages")
	// res_Collection := amqpSession.DB("utm-db").C("responsemessages")
	// err_Collection := amqpSession.DB("utm-db").C("errormessages")

	if err != nil {
		panic(err)
	}

	//var m AmqpMessage
	go func() {
		defer func() {
			// Close the downlinkMessages channel when
			//the AMQP handler is closed down and set it to nil
			downlinkMessages = nil
			close(q.Downlink)
			//amqpSession.Close()
		}()
		// Continually process AMQP messages until commanded to quit
		for {
			var msg amqp.Delivery
			receivedMsg := false
			select {
			case <-q.quit:
				return
			case msg = <-receiveChan:
				receivedMsg = true
				m := AmqpReceiveMessage{}
				err = json.Unmarshal(msg.Body, &m)
				log.Printf("MAY BE THERE IS AN ERROR HERE %+v\n", &m)
				if err == nil {
					//rec_Collection.Insert(&m)
					log.Printf("SENDING THE FOLLOWING DOWNLINQ MESSAGE %+v\n")
					q.Msgs <- &m
				}
			case msg = <-responseChan:
				receivedMsg = true
				m := AmqpResponseMessage{}
				err = json.Unmarshal(msg.Body, &m)
				if err == nil {
					q.Msgs <- &m
					log.Printf("%s UTM UUID IS %+v\n", logTag, m.Device_uuid)
					log.Printf("%s UTM UUID IS %+v\n", logTag, m.Device_name)
					Row.Uuid = m.Device_uuid
					Row.UnitName = m.Device_name
					//res_Collection.Insert(&m)
				}
			case msg = <-errorChan:
				receivedMsg = true
				m := AmqpErrorMessage{}
				err = json.Unmarshal(msg.Body, &m)
				if err == nil {
					//err_Collection.Insert(&m)
					q.Msgs <- &m
				}
			case dlMsg := <-q.Downlink:
				serialisedData, err := json.Marshal(dlMsg)
				if err != nil {
					log.Printf("%s ATTEMPTING TO JSONIFY AMQP MESSAGE %+v RESULTS IN ERROR: %s\n", logTag, dlMsg, err.Error())
				} else {
					publishedMsg := amqp.Publishing{
						DeliveryMode: amqp.Persistent,
						Timestamp:    time.Now(),
						ContentType:  "application/json",
						Body:         serialisedData,
					}
					err = channel.Publish(username, "send", false, false, publishedMsg)
					if err != nil {
						log.Printf("%s UNABLE TO PUBLISH DOWNLINK MESSAGE %+v DUE TO ERROR: %s\n", logTag, publishedMsg, err.Error())
					} else {
						log.Printf("%s PUBLISHED DOWNLINK MESSAGE %+v %s\n", logTag, publishedMsg, string(serialisedData))
					}
				}
			}
			if receivedMsg {
				if err == nil {
					log.Printf("%s GOT %+v\n", logTag, string(msg.Body))
				} else {
					log.Printf("%s RECEIVED %+v WHICH IS UNDECODABLE: %s\n", logTag, string(msg.Body), err.Error())
				}
			}

		}
	}()
	return &q, nil
}

func (q *Queue) Close() {
	// FIXME: maybe dont need a quit chan... can get EOF info from AMQP chans somehow.
	q.conn.Close()
	close(q.quit)
}
