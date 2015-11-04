package server

/*
#cgo CFLAGS: -I .
#include <stdbool.h>
#include <stdint.h>
#include <utm_api.h>

uint32_t pointerSub(const char *a, const char *b)
{
	return (uint32_t) (a - b);
}

PollIndUlMsg_t getPollIndUlMsg(UlMsgUnion_t in)
{
	return in.pollIndUlMsg;
}

InitIndUlMsg_t getInitIndUlMsg(UlMsgUnion_t in)
{
	return in.initIndUlMsg;
}

IntervalsGetCnfUlMsg_t getIntervalsGetCnfUlMsg(UlMsgUnion_t in)
{
	return in.intervalsGetCnfUlMsg;
}

ReportingIntervalSetCnfUlMsg_t getReportingIntervalSetCnfUlMsg(UlMsgUnion_t in)
{
	return in.reportingIntervalSetCnfUlMsg;
}

HeartbeatSetCnfUlMsg_t getHeartbeatSetCnfUlMsg(UlMsgUnion_t in)
{
	return in.heartbeatSetCnfUlMsg;
}

DebugIndUlMsg_t getDebugIndUlMsg(UlMsgUnion_t in)
{
	return in.debugIndUlMsg;
}

TrafficReportGetCnfUlMsg_t getTrafficReportGetCnfUlMsg(UlMsgUnion_t in)
{
	return in.trafficReportGetCnfUlMsg;
}

MeasurementsIndUlMsg_t getMeasurementsIndUlMsg(UlMsgUnion_t in)
{
	return in.measurementsIndUlMsg;
}
TrafficReportIndUlMsg_t getTrafficReportIndUlMsg(UlMsgUnion_t in)
{
	return in.trafficReportIndUlMsg;
}
*/
import "C"
import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
	"time"
	"unsafe"
)

var ulDecodeTypeDisplay map[int]string = map[int]string{
	C.DECODE_RESULT_TRANSPARENT_UL_DATAGRAM:                      "DECODE_RESULT_TRANSPARENT_UL_DATAGRAM",
	C.DECODE_RESULT_INIT_IND_UL_MSG:                              "DECODE_RESULT_INIT_IND_UL_MSG",
	C.DECODE_RESULT_POLL_IND_UL_MSG:                              "DECODE_RESULT_POLL_IND_UL_MSG",
	C.DECODE_RESULT_DATE_TIME_IND_UL_MSG:                         "DECODE_RESULT_DATE_TIME_IND_UL_MSG",
	C.DECODE_RESULT_DATE_TIME_SET_CNF_UL_MSG:                     "DECODE_RESULT_DATE_TIME_SET_CNF_UL_MSG",
	C.DECODE_RESULT_DATE_TIME_GET_CNF_UL_MSG:                     "DECODE_RESULT_DATE_TIME_GET_CNF_UL_MSG",
	C.DECODE_RESULT_MODE_SET_CNF_UL_MSG:                          "DECODE_RESULT_MODE_SET_CNF_UL_MSG",
	C.DECODE_RESULT_MODE_GET_CNF_UL_MSG:                          "DECODE_RESULT_MODE_GET_CNF_UL_MSG",
	C.DECODE_RESULT_HEARTBEAT_SET_CNF_UL_MSG:                     "DECODE_RESULT_HEARTBEAT_SET_CNF_UL_MSG",
	C.DECODE_RESULT_REPORTING_INTERVAL_SET_CNF_UL_MSG:            "DECODE_RESULT_REPORTING_INTERVAL_SET_CNF_UL_MSG",
	C.DECODE_RESULT_INTERVALS_GET_CNF_UL_MSG:                     "DECODE_RESULT_INTERVALS_GET_CNF_UL_MSG",
	C.DECODE_RESULT_MEASUREMENTS_GET_CNF_UL_MSG:                  "DECODE_RESULT_MEASUREMENTS_GET_CNF_UL_MSG",
	C.DECODE_RESULT_MEASUREMENTS_IND_UL_MSG:                      "DECODE_RESULT_MEASUREMENTS_IND_UIntervalL_MSG",
	C.DECODE_RESULT_MEASUREMENT_CONTROL_SET_CNF_UL_MSG:           "DECODE_RESULT_MEASUREMENT_CONTROL_SET_CNF_UL_MSG",
	C.DECODE_RESULT_MEASUREMENTS_CONTROL_GET_CNF_UL_MSG:          "DECODE_RESULT_MEASUREMENTS_CONTROL_GET_CNF_UL_MSG",
	C.DECODE_RESULT_MEASUREMENTS_CONTROL_IND_UL_MSG:              "DECODE_RESULT_MEASUREMENTS_CONTROL_IND_UL_MSG",
	C.DECODE_RESULT_MEASUREMENTS_CONTROL_DEFAULTS_SET_CNF_UL_MSG: "DECODE_RESULT_MEASUREMENTS_CONTROL_DEFAULTS_SET_CNF_UL_MSG",
	C.DECODE_RESULT_TRAFFIC_REPORT_IND_UL_MSG:                    "DECODE_RESULT_TRAFFIC_REPORT_IND_UL_MSG",
	C.DECODE_RESULT_TRAFFIC_REPORT_GET_CNF_UL_MSG:                "DECODE_RESULT_TRAFFIC_REPORT_GET_CNF_UL_MSG",
	C.DECODE_RESULT_TRAFFIC_TEST_MODE_PARAMETERS_SET_CNF_UL_MSG:  "DECODE_RESULT_TRAFFIC_TEST_MODE_PARAMETERS_SET_CNF_UL_MSG",
	C.DECODE_RESULT_TRAFFIC_TEST_MODE_PARAMETERS_GET_CNF_UL_MSG:  "DECODE_RESULT_TRAFFIC_TEST_MODE_PARAMETERS_GET_CNF_UL_MSG",
	C.DECODE_RESULT_TRAFFIC_TEST_MODE_RULE_BREAKER_UL_DATAGRAM:   "DECODE_RESULT_TRAFFIC_TEST_MODE_RULE_BREAKER_UL_DATAGRAM",
	C.DECODE_RESULT_TRAFFIC_TEST_MODE_REPORT_IND_UL_MSG:          "DECODE_RESULT_TRAFFIC_TEST_MODE_REPORT_IND_UL_MSG",
	C.DECODE_RESULT_TRAFFIC_TEST_MODE_REPORT_GET_CNF_UL_MSG:      "DECODE_RESULT_TRAFFIC_TEST_MODE_REPORT_GET_CNF_UL_MSG",
	C.DECODE_RESULT_DEBUG_IND_UL_MSG:                             "DECODE_RESULT_DEBUG_IND_UL_MSG",
}

type DiskSpaceLeft uint32

const (
	DISK_SPACE_LEFT_LESS_THAN_1_GB DiskSpaceLeft = C.DISK_SPACE_LEFT_LESS_THAN_1_GB
	DISK_SPACE_LEFT_MORE_THAN_1_GB DiskSpaceLeft = C.DISK_SPACE_LEFT_MORE_THAN_1_GB
	DISK_SPACE_LEFT_MORE_THAN_2_GB DiskSpaceLeft = C.DISK_SPACE_LEFT_MORE_THAN_2_GB
	DISK_SPACE_LEFT_MORE_THAN_4_GB DiskSpaceLeft = C.DISK_SPACE_LEFT_MORE_THAN_4_GB
	MAX_NUM_DISK_SPACE_LEFT        DiskSpaceLeft = C.MAX_NUM_DISK_SPACE_LEFT
)

type EnergyLeftEnum uint32

const (
	ENERGY_LEFT_LESS_THAN_5_PERCENT  EnergyLeftEnum = C.ENERGY_LEFT_LESS_THAN_5_PERCENT
	ENERGY_LEFT_LESS_THAN_10_PERCENT EnergyLeftEnum = C.ENERGY_LEFT_LESS_THAN_10_PERCENT
	ENERGY_LEFT_MORE_THAN_10_PERCENT EnergyLeftEnum = C.ENERGY_LEFT_MORE_THAN_10_PERCENT
	ENERGY_LEFT_MORE_THAN_30_PERCENT EnergyLeftEnum = C.ENERGY_LEFT_MORE_THAN_30_PERCENT
	ENERGY_LEFT_MORE_THAN_50_PERCENT EnergyLeftEnum = C.ENERGY_LEFT_MORE_THAN_50_PERCENT
	ENERGY_LEFT_MORE_THAN_70_PERCENT EnergyLeftEnum = C.ENERGY_LEFT_MORE_THAN_70_PERCENT
	ENERGY_LEFT_MORE_THAN_90_PERCENT EnergyLeftEnum = C.ENERGY_LEFT_MORE_THAN_90_PERCENT
	MAX_NUM_ENERGY_LEFT              EnergyLeftEnum = C.MAX_NUM_ENERGY_LEFT
)

type ModeEnum uint32

const (
	MODE_NULL          ModeEnum = C.MODE_NULL
	MODE_SELF_TEST     ModeEnum = C.MODE_SELF_TEST
	MODE_COMMISSIONING ModeEnum = C.MODE_COMMISSIONING
	MODE_STANDARD_TRX  ModeEnum = C.MODE_STANDARD_TRX
	MODE_TRAFFIC_TEST  ModeEnum = C.MODE_TRAFFIC_TEST
	MAX_NUM_MODES      ModeEnum = C.MAX_NUM_MODES
)

type WakeUpEnum uint32

const (
	WAKE_UP_CODE_OK       WakeUpEnum = C.WAKE_UP_CODE_OK
	WAKE_UP_CODE_WATCHDOG WakeUpEnum = C.WAKE_UP_CODE_WATCHDOG
	// WAKE_UP_CODE_AT_COMMAND_PROBLEM   WakeUpEnum = C.WAKE_UP_CODE_AT_COMMAND_PROBLEM
	// WAKE_UP_CODE_NETWORK_SEND_PROBLEM WakeUpEnum = C.WAKE_UP_CODE_NETWORK_SEND_PROBLEM
	WAKE_UP_CODE_MEMORY_ALLOC_PROBLEM WakeUpEnum = C.WAKE_UP_CODE_MEMORY_ALLOC_PROBLEM
	WAKE_UP_CODE_PROTOCOL_PROBLEM     WakeUpEnum = C.WAKE_UP_CODE_PROTOCOL_PROBLEM
	WAKE_UP_CODE_GENERIC_FAILURE      WakeUpEnum = C.WAKE_UP_CODE_GENERIC_FAILURE
	WAKE_UP_CODE_REBOOT               WakeUpEnum = C.WAKE_UP_CODE_REBOOT
)

type OrientationEnum uint32

const (
// ORIENTATION_UNCERTAIN     OrientationEnum = C.ORIENTATION_UNCERTAIN
// ORIENTATION_FACE_UP       OrientationEnum = C.ORIENTATION_FACE_UP
// ORIENTATION_FACE_DOWN     OrientationEnum = C.ORIENTATION_FACE_DOWN
// ORIENTATION_UPRIGHT       OrientationEnum = C.ORIENTATION_UPRIGHT
// ORIENTATION_UPSIDE_DOWN   OrientationEnum = C.ORIENTATION_UPSIDE_DOWN
// ORIENTATION_ON_LEFT_SIDE  OrientationEnum = C.ORIENTATION_ON_LEFT_SIDE
// ORIENTATION_ON_RIGHT_SIDE OrientationEnum = C.ORIENTATION_ON_RIGHT_SIDE
)

var OrientationEnumDisplay = map[OrientationEnum]string{
// ORIENTATION_UNCERTAIN:     "Unknown orientation",
// ORIENTATION_FACE_UP:       "Face up",
// ORIENTATION_FACE_DOWN:     "Face down",
// ORIENTATION_UPRIGHT:       "Upright",
// ORIENTATION_UPSIDE_DOWN:   "Upside down",
// ORIENTATION_ON_LEFT_SIDE:  "Lying on left side",
// ORIENTATION_ON_RIGHT_SIDE: "Lying on right side",
}

type ChargeStateEnum uint32

const (
// CHARGING_UNKNOWN ChargeStateEnum = C.CHARGING_UNKNOWN
// CHARGING_OFF     ChargeStateEnum = C.CHARGING_OFF
// CHARGING_ON      ChargeStateEnum = C.CHARGING_ON
// CHARGING_FAULT   ChargeStateEnum = C.CHARGING_FAULT
)

func decode(data []byte) {

	// Holder for the extracted message
	var inputBuffer C.union_UlMsgUnionTag_t
	inputPointer := (*C.UlMsgUnion_t)(unsafe.Pointer(&inputBuffer))

	// Loop over the messages in the datagram
	startPointer := (*C.char)(unsafe.Pointer(&data[0]))
	nextPointer := (*C.char)(unsafe.Pointer(&data[0]))
	nextPointerPointer := (**C.char)(unsafe.Pointer(&nextPointer))

	hexBuffer := hex.Dump(data)
	fmt.Printf("\n\n===>The Whole Input Butter (%s) \n", hexBuffer)

	var decoderCount int

	for {
		decoderCount = decoderCount + 1
		TotalMsgs = TotalMsgs + 1
		fmt.Println()
		fmt.Printf("\n\n===>DECODING MESSAGE NO (%v)  OF AMQP DATAGRAM NO (%v) \n", decoderCount, amqpCount)

		fmt.Printf("\n%s -----## SHOW BUFFER DATA ##----- \n\n %s\n\n", logTag, spew.Sdump(inputBuffer))
		used := C.pointerSub(nextPointer, startPointer)
		fmt.Printf("%s MESSAGE DATA USED %s\n", logTag, spew.Sdump(used))
		remaining := C.uint32_t(len(data)) - used
		fmt.Printf("%s MESSAGE DATA REMAINING %s\n", logTag, spew.Sdump(remaining))
		if remaining == 0 {
			break
		}

		result := C.decodeUlMsg(nextPointerPointer, remaining, inputPointer, nil, nil)
		fmt.Printf("%s DECODE RECEIVED UPLINK MESSAGE %+v \n\n -----## %s ##----- \n\n", logTag, result, ulDecodeTypeDisplay[int(result)])

		// Extract any data to be recorded; the C symbols are not available outside
		// the package so convert into concrete go types
		var data interface{} = nil
		var rawData interface{} = nil
		var multipleRecords []interface{} = nil

		switch int(result) {
		case C.DECODE_RESULT_INIT_IND_UL_MSG:
			value := C.getInitIndUlMsg(inputBuffer)
			rawData = value
			data = &InitIndUlMsg{
				Timestamp:         time.Now(),
				WakeUpCode:        WakeUpEnum(value.wakeUpCode),
				RevisionLevel:     uint16(value.revisionLevel),
				sdCardNotRequired: bool(value.sdCardNotRequired),
			}

		case C.DECODE_RESULT_INTERVALS_GET_CNF_UL_MSG:
			value := C.getIntervalsGetCnfUlMsg(inputBuffer)
			rawData = value
			data = &IntervalsGetCnfUlMsg{
				ReportingInterval:  uint32(value.reportingInterval),
				HeartbeatSeconds:   uint32(value.heartbeatSeconds),
				HeartbeatSnapToRtc: bool(value.heartbeatSnapToRtc),
			}

			Row.ReportingInterval = uint32(value.reportingInterval)
			Row.HeartbeatSeconds = uint32(value.heartbeatSeconds)
			Row.HeartbeatSnapToRtc = bool(value.heartbeatSnapToRtc)
			multipleRecords = append(multipleRecords, Row)

			fmt.Printf("%s RECEIVED INTERVAL REQUEST CONFIRM %s\n", logTag, spew.Sdump(Row))

		case C.DECODE_RESULT_REPORTING_INTERVAL_SET_CNF_UL_MSG:
			value := C.getReportingIntervalSetCnfUlMsg(inputBuffer)
			rawData = value
			data = &ReportingIntervalSetCnfUlMsg{
				Timestamp: time.Now(),
				//ReportingIntervalMinutes: uint32(value.reportingInterval),
			}

		case C.DECODE_RESULT_HEARTBEAT_SET_CNF_UL_MSG:
			rawData = C.getHeartbeatSetCnfUlMsg(inputBuffer)

		case C.DECODE_RESULT_POLL_IND_UL_MSG:
			value := C.getPollIndUlMsg(inputBuffer)
			rawData = value
			data = &PollIndUlMsg{
				Mode:          string(ModeEnum(value.mode)),
				EnergyLeft:    string(EnergyLeftEnum(value.energyLeft)),
				DiskSpaceLeft: string(DiskSpaceLeft(value.diskSpaceLeft)),
			}

			Row.TotalMsgs = TotalMsgs
			Row.Mode = ModeLookUp[string(ModeEnum(value.mode))]
			Row.BatteryLevel = EnergyLeftLookUP[int(EnergyLeftEnum(value.energyLeft))]
			Row.DiskSpaceLeft = DiskSpaceLookUP[int(DiskSpaceLeft(value.diskSpaceLeft))]

			multipleRecords = append(multipleRecords, Row)

			//encodeAndEnqueueIntervalGetReq(Row.Uuid)

		case C.DECODE_RESULT_DEBUG_IND_UL_MSG:
			rawData = C.getDebugIndUlMsg(inputBuffer)

		case C.DECODE_RESULT_TRAFFIC_REPORT_GET_CNF_UL_MSG:
			rawData = C.getTrafficReportGetCnfUlMsg(inputBuffer)

		case C.DECODE_RESULT_TRAFFIC_REPORT_IND_UL_MSG:
			value := C.getTrafficReportIndUlMsg(inputBuffer)
			rawData = value
			data = &TrafficReportIndUlMsg{
				numDatagramsUl: int32(value.numDatagramsUl),
				numBytesUl:     int32(value.numBytesUl),
				numDatagramsDl: int32(value.numDatagramsDl),
				numBytesDl:     int32(value.numBytesDl),
			}

		case C.DECODE_RESULT_TRANSPARENT_UL_DATAGRAM:
			rawData = "value"
		case C.DECODE_RESULT_DATE_TIME_IND_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_DATE_TIME_SET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_DATE_TIME_GET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_MODE_SET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_MODE_GET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_MEASUREMENTS_GET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_MEASUREMENTS_IND_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_MEASUREMENT_CONTROL_SET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_MEASUREMENTS_CONTROL_GET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_MEASUREMENTS_CONTROL_IND_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_MEASUREMENTS_CONTROL_DEFAULTS_SET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_TRAFFIC_TEST_MODE_PARAMETERS_SET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_TRAFFIC_TEST_MODE_PARAMETERS_GET_CNF_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_TRAFFIC_TEST_MODE_RULE_BREAKER_UL_DATAGRAM:
			rawData = "value"
		case C.DECODE_RESULT_TRAFFIC_TEST_MODE_REPORT_IND_UL_MSG:
			rawData = "value"
		case C.DECODE_RESULT_TRAFFIC_TEST_MODE_REPORT_GET_CNF_UL_MSG:
			rawData = "value"

		default:
			// Can't decode the uplink; this case is here to avoid a panic,
			// the code below will handle this situation fine
		}

		// Send any data to be recorded
		if multipleRecords != nil {
			// Iterate once to get a description of the objects to be recorded and display it
			txt := ""
			for _, record := range multipleRecords {
				txt = fmt.Sprintf("%s%+v ", txt, record)
			}
			fmt.Printf("%s MULTIPLE DECODED ENTRIES BEING RECORDED = %s\n", logTag, txt)
			// Then record the objects
			for _, record := range multipleRecords {
				StateTableCmds <- record
			}
		} else if data != nil {
			fmt.Printf("%s DECODED DATA BEING RECORDED = %+v\n", logTag, data)
			StateTableCmds <- data
		} else if rawData != nil {
			fmt.Printf("%s DATA NOT BEING RECORDED = %+v\n", logTag, rawData)
		} else {
			fmt.Printf("%s ERROR: UNDECODABLE MESSGE RECEIVED %s\n", logTag, spew.Sdump(inputBuffer))
		}

	}

	decoderCount = 0
}

func encodeAndEnqueueReportingInterval(mins uint32) error {
	if downlinkMessages != nil {

		// Create a buffer that is big enough to store all
		// the encoded data and take a pointer to it's first element
		var outputBuffer [C.MAX_DATAGRAM_SIZE_RAW]byte
		outputPointer := (*C.char)(unsafe.Pointer(&outputBuffer[0]))

		// Populate the data struture to be encoded
		var data C.ReportingIntervalSetReqDlMsg_t

		//		data.reportingIntervalMinutes = C.uint32_t(mins)
		var dataPointer = (*C.ReportingIntervalSetReqDlMsg_t)(unsafe.Pointer(&data))

		// Encode the data structure into the output buffer
		cbytes := C.encodeReportingIntervalSetReqDlMsg(outputPointer, dataPointer, nil, nil)

		// Send the populated output buffer bytes to NeulNet
		payload := outputBuffer[:cbytes]
		msg := AmqpMessage{
			Device_uuid:   ueGuid,
			Endpoint_uuid: 4,
			//Payload:       payload,
		}
		for _, v := range payload {
			msg.Payload = append(msg.Payload, int(v))
		}
		log.Printf("%s ENCODED A REPORTING INTERVAL UPDATE OF %d MINUTES AS %v USING AMQP MESSAGE %+v\n", logTag, mins, payload, msg)

		downlinkMessages <- msg
		now := time.Now()

		// Record the downlink data volume
		StateTableCmds <- &DataVolume{
			DownlinkTimestamp: &now,
			DownlinkBytes:     uint64(len(payload)),
		}

		return nil
	}

	return errors.New("NO DOWNLINKMESSAGES CHANNEL AVAILABLE TO ENQUEUE THE ENCODED MESSAGE")
}

func encodeAndEnqueueIntervalGetReq(ueGuid string) error {
	if downlinkMessages != nil {

		// Create a buffer that is big enough to store all
		// the encoded data and take a pointer to it's first element
		var outputBuffer [C.MAX_DATAGRAM_SIZE_RAW]byte
		outputPointer := (*C.char)(unsafe.Pointer(&outputBuffer[0]))

		// Encode the data structure into the output buffer
		cbytes := C.encodeIntervalsGetReqDlMsg(outputPointer, nil, nil)

		// Send the populated output buffer bytes to NeulNet
		payload := outputBuffer[:cbytes]

		msg := AmqpMessage{
			Device_uuid:   ueGuid,
			Endpoint_uuid: 4,
			//Payload:       payload,
		}
		for _, v := range payload {
			msg.Payload = append(msg.Payload, int(v))
		}
		log.Printf("%s ENCODED A REPORTING INTERVAL UPDATE OF %d MINUTES AS USING AMQP MESSAGE %+v\n", logTag, payload, msg)

		downlinkMessages <- msg
		now := time.Now()

		// Record the downlink data volume
		StateTableCmds <- &DataVolume{
			DownlinkTimestamp: &now,
			DownlinkBytes:     uint64(len(payload)),
		}

		//Record in the DisplayRow
		Row.DtotalMsgs = Row.DtotalMsgs + 1
		Row.DTotalBytes = Row.DTotalBytes + uint64(len(payload))
		Row.DlastMsgReceived = &now

		return nil
	}

	return errors.New("NO DOWNLINKMESSAGES CHANNEL AVAILABLE TO ENQUEUE THE ENCODED MESSAGE")
}
