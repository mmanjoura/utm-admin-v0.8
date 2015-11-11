package server

import (
	"log"

	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type LatestState struct {
	Connection                         Connection
	LatestRssi                         *Rssi                         `json:"LatestRssi,omitempty"`
	LatestRssiDisplay                  *RssiDisplay                  `json:"LatestRssiDisplay,omitempty"`
	LatestPowerState                   *PowerState                   `json:"LatestPowerState,omitempty"`
	LatestPowerStateDisplay            *PowerStateDisplay            `json:"LatestPowerStateDisplay,omitempty"`
	LatestInitIndUlMsg                 *InitIndUlMsg                 `json:"LatestInitIndUlMsg,omitempty"`
	LatestDataVolume                   *DataVolume                   `json:"LatestDataVolume,omitempty"`
	LatestReportingIntervalSetCnfUlMsg *ReportingIntervalSetCnfUlMsg `json:"LatestReportingIntervalSetCnfUlMsg,omitempty"`
	LatestPollIndUlMsg                 *PollIndUlMsg                 `json:"LastestPollIndUlMsg,omitempty"`
	LatestTrafficReportIndUlMsg        *TrafficReportIndUlMsg        `json:"LatestTrafficReportIndUlMsg,omitempty"`
	LatestDisplayRow                   *DisplayRow                   `json:"LatestDisplayRow,omitempty"`
	LatestMeasurementsIndUlMsg         *MeasurementsIndUlMsg         `json:"LatestMeasurementsIndUlMsg,omitempty"`
}

// To update a latest value send a copy of one of the following structs into this channel: *GpsPosition,
// *LclPosition, *SoundLevel, *Luminosity, *Temperature, *Rssi, *PowerState, *InitIndUlMsg; an independant
// copy of their contents will be stored

// To get the latest state send a '*chan LatestState' into this channel; a LatestState struct containing
// an independant copy of all quantities (the memory contents pointed to will never be changed by the
// state table will be put onto the channel and then the channel closed To terminate execution simply close chan
var StateTableCmds chan<- interface{}

func operateStateTable() {
	cmds := make(chan interface{})
	StateTableCmds = cmds
	log.Printf("%s DATA TABLE COMMAND CHANNEL CREATED AND NOW BEING SERVICED\n", logTag)
	go func() {
		rssi := &RSSIData{}
		state := LatestState{}
		for msg := range cmds {
			log.Printf("%s DATA TABLE COMMAND: %+v\n", logTag, msg)

			switch value := msg.(type) {

			case *Connection:
				state.Connection = *value
				log.Printf("%s SETTING CONNECTION STATE IN DATA TABLE: %+v\n", logTag, value)

			case *Rssi:
				state.LatestRssi = value.DeepCopy()
				if value != nil {
					state.LatestRssiDisplay = rssi.makeRssiDisplay(value)
				}

			case *PowerState:
				state.LatestPowerState = value.DeepCopy()
				state.LatestPowerStateDisplay = updatePowerStateDisplay(state, value)

			case *InitIndUlMsg:
				// Wipe out all stored state because
				// the UTM-API has been power cycled
				state = LatestState{}
				state.Connection = Connection{Status: "Connected"}
				state.LatestInitIndUlMsg = value.DeepCopy()

			case *DataVolume:
				state.LatestDataVolume = updateDataVolume(state, value)
				fmt.Printf("\n%s -----## DATA VOLUME STATE ##----- \n\n %s\n\n", logTag, spew.Sdump(state))
				fmt.Printf("\n%s -----## DATA VOLUME VALUE ##----- \n\n %s\n\n", logTag, spew.Sdump(value))

			case *ReportingIntervalSetCnfUlMsg:
				state.LatestReportingIntervalSetCnfUlMsg = value.DeepCopy()

			case *MeasurementsIndUlMsg:
				state.LatestMeasurementsIndUlMsg = value.DeepCopy()

			case *PollIndUlMsg:
				state.LatestPollIndUlMsg = value.DeepCopy()

			case *DisplayRow:
				state.LatestDisplayRow = value.DeepCopy()

			case *TrafficReportIndUlMsg:
				state.LatestTrafficReportIndUlMsg = value.DeepCopy()

			case *chan LatestState:

				// Duplicate the memory pointed to into a new LatestState struct,
				//post it and close the channel
				log.Printf("%s STATE TABLE FETCHING LATEST STATE\n", logTag)
				// time.Sleep(time.Second * 1)
				latest := LatestState{}
				latest.Connection = state.Connection
				latest.LatestRssi = state.LatestRssi.DeepCopy()
				latest.LatestRssiDisplay = state.LatestRssiDisplay.DeepCopy()
				latest.LatestPowerState = state.LatestPowerState.DeepCopy()
				latest.LatestPowerStateDisplay = state.LatestPowerStateDisplay.DeepCopy()
				latest.LatestInitIndUlMsg = state.LatestInitIndUlMsg.DeepCopy()
				latest.LatestDataVolume = state.LatestDataVolume.DeepCopy()
				latest.LatestReportingIntervalSetCnfUlMsg = state.LatestReportingIntervalSetCnfUlMsg.DeepCopy()
				latest.LatestPollIndUlMsg = state.LatestPollIndUlMsg.DeepCopy()
				latest.LatestTrafficReportIndUlMsg = state.LatestTrafficReportIndUlMsg.DeepCopy()
				latest.LatestDisplayRow = state.LatestDisplayRow.DeepCopy()
				*value <- latest
				close(*value)
				log.Print("%s DATA TABLE PROVIDED LATEST STATE AND CLOSED RETURN CHANNEL\n", logTag)

			default:
				log.Printf("%s UNRECOGNISED DATA TABLE COMMAND; IGNORING: %s\n", logTag, spew.Sdump(msg))
			}
		}

		log.Printf("%s DATA TABLE COMMAND CHANNEL CLOSED; STOPPING\n", logTag)
	}()

}

func init() {
	operateStateTable()
}
