// tap off and record messages related to a device

package server

import (
	"encoding/json"
	"fmt"
	"github.com/brettlangdon/forge"
	"github.com/mmanjoura/utm-admin-v0.8/service/utilities"
	"log"
	"net/http"
	//"os"

	"time"

	"github.com/codegangsta/negroni"
	"github.com/davecgh/go-spew/spew"
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/mmanjoura/utm-admin-v0.8/service/routes"
)

var logTag string = "UTM-API"
var listenAddress string
var downlinkMessages chan<- AmqpMessage
var ueGuid string
var amqpCount int
var displayRow = &DisplayRow{}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func getLatestState(response http.ResponseWriter, request *http.Request) *utilities.Error {
	// Ensure this is a GET request
	if (request.Method != "GET") ||
		(request.Method == "") {
		log.Printf("%s RECEIVED UNSUPPORTED REST REQUEST %s %s\n", logTag, request.Method, request.URL)
		return utilities.ClientError("UNSUPPORTED METHOD", http.StatusBadRequest)
	}

	// Get the latest state; only one response will come back before the requesting channel is closed
	get := make(chan LatestState)
	StateTableCmds <- &get
	state := <-get

	// Send the requested data
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	err := json.NewEncoder(response).Encode(state)
	if err != nil {
		log.Printf("%s RECEIVED REST REQUEST %s BUT ATTEMPTING TO SERIALISE THE RESULT %s YIELDED ERROR %s\n", logTag, request.URL, spew.Sdump(state), err.Error())
		return utilities.ServerError(err)
	}
	//log.Printf("%s Received rest request %s and responding with %s\n", logTag, request.URL, spew.Sdump(state))

	return nil
}

//DownLink (DL) messages
func setReportingInterval(response http.ResponseWriter, request *http.Request) *utilities.Error {
	// Ensure this is a POST request
	if request.Method != "POST" {
		log.Printf("%s RECEIVED UNSUPPORTED REST REQUEST %s %s\n", logTag, request.Method, request.URL)
		return utilities.ClientError("UNSUPPORTED METHOD", http.StatusBadRequest)
	}

	// Get the minutes interval
	var mins uint32
	err := json.NewDecoder(request.Body).Decode(&mins)
	if err != nil {
		log.Printf("%s UNABLE TO EXTRACT THE REPORTING INTERVAL FROM REQUEST %s: %s\n", logTag, request.URL, err.Error())
		return utilities.ClientError("UNABLE TO DECIPHER REPORTING INTERVAL", http.StatusBadRequest)
	}

	// Encode and enqueue the requested data
	err = encodeAndEnqueueReportingInterval(uint32(mins))
	if err != nil {
		log.Printf("%s UNABLE TO ENCODE AND ENQUEUE REPORTING INTERVAL UPDATE FOR UTM-API %s\n", logTag, request.URL)
		return utilities.ClientError("UANBLE TO ENCODE AND ENQUEUE REPORTING INTERVAL", http.StatusBadRequest)
	}

	// Success
	response.WriteHeader(http.StatusOK)
	return nil
}

func processAmqp(username, amqpAddress string) {

	//ueGuid, err := amqp.GetString("ueguid")

	// create a queue and bind it with relevant routing keys
	// amqpAddress := utilities.EnvStr("AMQP_ADDRESS")
	// username := utilities.EnvStr("UNAME")
	// ueGuid = utilities.EnvStr("UEGUID")

	q, err := OpenQueue(username, amqpAddress)

	failOnError(err, "Queue")
	defer q.Close()
	downlinkMessages = q.Downlink

	StateTableCmds <- &Connection{Status: "Disconnected"}

	connState := "Disconnected"
	connected := connState

	for {
		amqpCount = amqpCount + 1
		log.Println()
		fmt.Printf("\n\n=====================> PROCESSING DATAGRAM NO (%v) IN AMQP CHANNEL: =====================================\n", amqpCount)

		//time.Sleep(time.Second * 10)
		select {
		case <-time.After(30 * time.Minute):
			connected = "Disconnected"

		case msg := <-q.Msgs:
			log.Println(logTag, "DECODED MSG:", msg)

			switch value := msg.(type) {
			case *AmqpReceiveMessage:
				log.Println(logTag, "IS RECEIVE")

			case *AmqpResponseMessage:
				log.Println(logTag, "IS RESPONSE")
				if value.Command == "UART_data" || value.Command == "LOOPBACK_data" {
					// UART data from the UTM-API which needs to be decoded
					// NOTE: some old data extracted from logs is loopback_data. FIXME: remove this.
					decode(value.Data)
					// Get the amount of uplink data and send it on to the data table
					now := time.Now()
					StateTableCmds <- &DataVolume{
						UplinkTimestamp: &now,
						UplinkBytes:     uint64(len(value.Data)),
					}
					Row.UlastMsgReceived = &now
					Row.UTotalBytes = uint64(len(value.Data))
					connected = "CONNECTED"
				}

			case *AmqpErrorMessage:
				log.Println(logTag, "IS ERROR")

			default:
				log.Printf("%s MESSAGE TYPE %+v\n", logTag, msg)
				log.Fatal(logTag, "INVALID MESSAGE TYPE")
			}
		}

		if connState != connected {
			connState = connected
			log.Printf("%s SENDING NEW CONNECTION STATE: %s\n", logTag, connState)
			StateTableCmds <- &Connection{Status: connState}
		}

	}
	amqpCount = 0
}

func Run() {

	settings, err := forge.ParseFile("config.cfg")
	if err != nil {
		panic(err)
	}

	amqp, err := settings.GetSection("amqp")
	username, err := amqp.GetString("uname")
	amqpAddress, err := amqp.GetString("amqp_address")

	host, err := settings.GetSection("host")
	port, err := host.GetString("port")

	// Process Amqp messages
	go processAmqp(username, amqpAddress)

	log.SetFlags(log.LstdFlags | log.Llongfile)

	log.Printf("UTM-API SERVICE (%s) REST INTERFACE LISTENING ON %s\n", logTag, listenAddress)

	store := cookiestore.New([]byte("secretkey789"))
	router := routes.LoadRoutes()

	router.Handle("/latestState", utilities.Handler(getLatestState))
	router.Handle("/reportingInterval", utilities.Handler(setReportingInterval))

	n := negroni.Classic()
	static := negroni.NewStatic(http.Dir("static"))
	static.Prefix = "/static"
	n.Use(static)
	//n.Use(negroni.HandlerFunc(system.MgoMiddleware))
	n.Use(sessions.Sessions("global_session_store", store))
	n.UseHandler(router)
	n.Run(port)

}
