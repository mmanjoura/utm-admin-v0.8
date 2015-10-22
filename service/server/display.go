package server

import (
	"time"
)

type TemperatureDisplay struct {
	Timestamp       time.Time
	TemperatureDegC float32
}

type RssiDisplay struct {
	Timestamp         time.Time
	RssiDbm           float32
	MCL               float32
	HighestMCL        float32
	InsideGsmCoverage bool
}

type PowerStateDisplay struct {
	Timestamp        time.Time
	ChargeStateText  string
	BatteryVoltageV  float32
	EnergyConsumedWh float32
}

type MCLValue struct {
	MCL       float32
	Timestamp time.Time
}

type RSSIData struct {
	History []MCLValue
}

func (value *TemperatureDisplay) DeepCopy() *TemperatureDisplay {
	if value == nil {
		return nil
	}
	result := &TemperatureDisplay{
		Timestamp:       value.Timestamp,
		TemperatureDegC: value.TemperatureDegC,
	}
	return result
}

func (value *RssiDisplay) DeepCopy() *RssiDisplay {
	if value == nil {
		return nil
	}
	result := &RssiDisplay{
		Timestamp:         value.Timestamp,
		RssiDbm:           value.RssiDbm,
		MCL:               value.MCL,
		HighestMCL:        value.HighestMCL,
		InsideGsmCoverage: value.InsideGsmCoverage,
	}
	return result
}

func (value *PowerStateDisplay) DeepCopy() *PowerStateDisplay {
	if value == nil {
		return nil
	}
	result := &PowerStateDisplay{
		Timestamp:        value.Timestamp,
		ChargeStateText:  value.ChargeStateText,
		BatteryVoltageV:  value.BatteryVoltageV,
		EnergyConsumedWh: value.EnergyConsumedWh,
	}
	return result
}

func makeTemperatureDisplay(newData *Temperature) *TemperatureDisplay {
	display := TemperatureDisplay{
		Timestamp:       newData.Timestamp,
		TemperatureDegC: float32(newData.Temperature),
	}
	return &display
}

func (rd *RSSIData) makeRssiDisplay(newData *Rssi) *RssiDisplay {
	display := RssiDisplay{
		Timestamp: newData.Timestamp,
	}
	// This RSSI calculation and GSM coverage comparison is hokum because a) the data values are not accurate b) the reported data is not appropriate for making this comparison.  See Peter Smith for details
	//display.RssiDbm = float32(newData.Rssi)*3.0 - 207.0

	// Rob Meades conversion formula: tedISampleApp/CommsForm.cs
	var csqAtNeg80dBm float32 = 37
	display.RssiDbm = -80.0 + (float32(newData.Rssi)-csqAtNeg80dBm)*3.0

	display.MCL = 32 - display.RssiDbm

	display.InsideGsmCoverage = (display.RssiDbm >= -106.0)

	mv := MCLValue{MCL: display.MCL, Timestamp: display.Timestamp}
	rd.History = append(rd.History, mv)

	if len(rd.History) > 10 {
		rd.History = rd.History[1:]
	}

	display.HighestMCL = display.MCL
	for _, v := range rd.History {
		if display.Timestamp.Sub(v.Timestamp).Minutes() > 60 {
			break
		}
		if display.HighestMCL < v.MCL {
			display.HighestMCL = v.MCL
		}
	}

	return &display
}

func updateDataVolume(currentState LatestState, newData *DataVolume) *DataVolume {
	// Change nothing if the new data is nil
	if newData == nil {
		return currentState.LatestDataVolume
	}
	// Overwrite if there is no stored data
	if currentState.LatestDataVolume == nil {
		return newData.DeepCopy()
	}
	// Update accumulated values
	result := &DataVolume{
		UplinkBytes:   currentState.LatestDataVolume.UplinkBytes + newData.UplinkBytes,
		DownlinkBytes: currentState.LatestDataVolume.DownlinkBytes + newData.DownlinkBytes,
	}
	if newData.UplinkTimestamp != nil {
		result.UplinkTimestamp = newData.UplinkTimestamp
	}
	if newData.DownlinkTimestamp != nil {
		result.DownlinkTimestamp = newData.DownlinkTimestamp
	}
	return result
}

func updatePowerStateDisplay(currentState LatestState, newData *PowerState) *PowerStateDisplay {
	// Change nothing if the new data is nil
	if newData == nil {
		return currentState.LatestPowerStateDisplay
	}
	// Overwrite if there is no stored data
	display := &PowerStateDisplay{
		Timestamp:        newData.Timestamp,
		ChargeStateText:  ChargeStateEnumDisplay[newData.ChargeState],
		BatteryVoltageV:  float32(newData.BatteryMV) / 1000.0,
		EnergyConsumedWh: float32(newData.EnergyUWH) / 1000000.0,
	}
	if currentState.LatestPowerStateDisplay == nil {
		return display
	}
	// Update accumulated values
	display.EnergyConsumedWh = display.EnergyConsumedWh + currentState.LatestPowerStateDisplay.EnergyConsumedWh
	return display
}
