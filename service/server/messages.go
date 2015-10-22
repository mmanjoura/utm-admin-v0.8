package server

import (
	"time"
)

type DisplayRow struct {
	TotalMsgs        uint64     `json:"TotalMsgs,omitempty"`
	Uuid             string     `json:"Uuid,omitempty"`
	UnitName         string     `json:"UnitName, omitempty"`
	Mode             string     `json:"Mode, omitempty"`
	UTotalMsgs       uint64     `json:"UTotalMsgs, omitempty"`
	UTotalBytes      uint64     `json:"UTotalBytes, omitempty"`
	UlastMsgReceived *time.Time `json:"UlastMsgReceived, omitempty"`
	DtotalMsgs       uint64     `json:"DtotalMsgs, omitempty"`
	DTotalBytes      uint64     `json:"DTotalBytes, omitempty"`
	DlastMsgReceived *time.Time `json:"DlastMsgReceived, omitempty"`
	RSRP             int32      `json:"RSRP, omitempty"`
	BatteryLevel     string     `json:"BatteryLevel, omitempty"`
	DiskSpaceLeft    string     `json:"DiskSpaceLeft, omitempty"`
}

type Connection struct {
	Status string
}

type Luminosity struct {
	Timestamp  time.Time
	Luminosity uint16
}

type Temperature struct {
	Timestamp   time.Time
	Temperature int8
}

type Rssi struct {
	Timestamp time.Time
	Rssi      uint8
}

var ChargeStateEnumDisplay = map[ChargeStateEnum]string{
// 	CHARGING_UNKNOWN: "Unknown",
// 	CHARGING_OFF:     "No",
// 	CHARGING_ON:      "Charging",
// 	CHARGING_FAULT:   "Fault",
}

type PowerState struct {
	Timestamp   time.Time
	ChargeState ChargeStateEnum
	BatteryMV   uint16
	EnergyUWH   uint32
}

type InitIndUlMsg struct {
	Timestamp         time.Time
	WakeUpCode        WakeUpEnum
	RevisionLevel     uint16
	sdCardNotRequired bool
}
type IntervalsGetCnfUlMsg struct {
	reportingInterval  int32
	heartbeatSeconds   int32
	heartbeatSnapToRtc bool
}

type ReportingIntervalSetCnfUlMsg struct {
	Timestamp                time.Time
	ReportingIntervalMinutes uint32
}

type TrafficReportIndUlMsg struct {
	numDatagramsUl int32
	numBytesUl     int32
	numDatagramsDl int32
	numBytesDl     int32
}
type PollIndUlMsg struct {
	Mode          string
	EnergyLeft    string
	DiskSpaceLeft string
}

type DataVolume struct {
	UplinkTimestamp   *time.Time
	DownlinkTimestamp *time.Time
	UplinkBytes       uint64
	DownlinkBytes     uint64
}

// type Measurements struct {
// 	time                int32 //!< Time in UTC seconds.
// 	gnssPositionPresent bool
// 	gnssPosition        GnssPosition
// 	cellIdPresent       bool
// 	cellId              uint16
// 	rsrpPresent         bool
// 	rsrp                Rsrp
// 	rssiPresent         bool
// 	rssi                int32
// 	temperaturePresent  bool
// 	temperature         int8
// 	powerStatePresent   bool
// 	powerState          PowerState
// }

// type PowerState struct {
// 	//ChargerState ChargeStateEnumDisplay
// 	batteryMV uint16 //!< battery voltage in mV, max 10,000 mV, will be coded to a resolution of 159 mV (i.e. in 5 bits).
// 	energyMWh uint32 //!< mWh left in the battery, only up to the first 24 bits are valid (i.e. max value is 16777215).
// }

// type GnssPosition struct {
// 	latitude  int32 //!< In thousandths of a minute of arc (so divide by 60,000 to get degrees)
// 	longitude int32 //!< In thousandths of a minute of arc (so divide by 60,000 to get degrees)
// 	elevation int32 //!< In metres
// }

// type Rsrp struct {
// 	value            int16 //!< The RSRP value in 10ths of a dBm.
// 	isSyncedWithRssi bool  //!< If true then the RSRP value was
// 	//! taken at the same time as the RSSI
// 	//! value in a report.  In this case
// 	//! SNR can be computed.
// }

func (value *DataVolume) DeepCopy() *DataVolume {
	if value == nil {
		return nil
	}
	result := &DataVolume{
		UplinkTimestamp:   value.UplinkTimestamp,
		DownlinkTimestamp: value.DownlinkTimestamp,
		UplinkBytes:       value.UplinkBytes,
		DownlinkBytes:     value.DownlinkBytes,
	}
	return result
}

func (value *Temperature) DeepCopy() *Temperature {
	if value == nil {
		return nil
	}
	result := &Temperature{
		Timestamp:   value.Timestamp,
		Temperature: value.Temperature,
	}
	return result
}

func (value *Rssi) DeepCopy() *Rssi {
	if value == nil {
		return nil
	}
	result := &Rssi{
		Timestamp: value.Timestamp,
		Rssi:      value.Rssi,
	}
	return result
}

func (value *PowerState) DeepCopy() *PowerState {
	if value == nil {
		return nil
	}
	result := &PowerState{
		Timestamp:   value.Timestamp,
		ChargeState: value.ChargeState,
		BatteryMV:   value.BatteryMV,
		EnergyUWH:   value.EnergyUWH,
	}
	return result
}

func (value *InitIndUlMsg) DeepCopy() *InitIndUlMsg {
	if value == nil {
		return nil
	}
	result := &InitIndUlMsg{
		Timestamp:         value.Timestamp,
		WakeUpCode:        value.WakeUpCode,
		RevisionLevel:     value.RevisionLevel,
		sdCardNotRequired: value.sdCardNotRequired,
	}
	return result
}

func (value *TrafficReportIndUlMsg) DeepCopy() *TrafficReportIndUlMsg {
	if value == nil {
		return nil
	}
	result := &TrafficReportIndUlMsg{

		numDatagramsUl: value.numDatagramsUl,
		numBytesUl:     value.numBytesUl,
		numDatagramsDl: value.numDatagramsDl,
		numBytesDl:     value.numBytesDl,
	}
	return result
}

func (value *PollIndUlMsg) DeepCopy() *PollIndUlMsg {
	if value == nil {
		return nil
	}
	result := &PollIndUlMsg{
		Mode:          value.Mode,
		EnergyLeft:    value.EnergyLeft,
		DiskSpaceLeft: value.DiskSpaceLeft,
	}
	return result
}

func (value *ReportingIntervalSetCnfUlMsg) DeepCopy() *ReportingIntervalSetCnfUlMsg {
	if value == nil {
		return nil
	}
	result := &ReportingIntervalSetCnfUlMsg{
		Timestamp:                value.Timestamp,
		ReportingIntervalMinutes: value.ReportingIntervalMinutes,
	}
	return result
}

func (value *DisplayRow) DeepCopy() *DisplayRow {
	if value == nil {
		return nil
	}
	result := &DisplayRow{
		TotalMsgs:        value.TotalMsgs,
		Uuid:             value.Uuid,
		UnitName:         value.UnitName,
		Mode:             value.Mode,
		UTotalMsgs:       value.UTotalMsgs,
		UTotalBytes:      value.UTotalBytes,
		UlastMsgReceived: value.UlastMsgReceived,
		DtotalMsgs:       value.DtotalMsgs,
		DTotalBytes:      value.DTotalBytes,
		DlastMsgReceived: value.DlastMsgReceived,
		RSRP:             value.RSRP,
		BatteryLevel:     value.BatteryLevel,
		DiskSpaceLeft:    value.DiskSpaceLeft,
	}
	return result
}
