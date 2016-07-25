/**
 * Copyright (C) 2016 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package power

import (
	"path/filepath"
	"pkg.deepin.io/dde/api/powersupply/battery"
	"pkg.deepin.io/lib/dbus"
	"strings"
)

const (
	batteryDBusIFC = dbusIFC + ".Battery"
)

func (bat *Battery) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		Dest:       dbusDest,
		ObjectPath: dbusPath + "/battery_" + getValidName(filepath.Base(bat.SysfsPath)),
		Interface:  batteryDBusIFC,
	}
}

func getValidName(n string) string {
	// dbus objpath 0-9 a-z A-Z _
	n = strings.Replace(n, "-", "_x0", -1)
	n = strings.Replace(n, ".", "_x1", -1)
	n = strings.Replace(n, ":", "_x2", -1)
	return n
}

func (bat *Battery) setPropUpdateTime(updateTime int64) {
	if bat.UpdateTime != updateTime {
		bat.UpdateTime = updateTime
		bat.notifyChange("UpdateTime")
	}
}

func (bat *Battery) setPropEnergyNow(energyNow uint64) {
	if bat.EnergyNow != energyNow {
		bat.EnergyNow = energyNow
		bat.notifyChange("EnergyNow")
	}
}

func (bat *Battery) setPropEnergyFull(energyFull uint64) {
	if bat.EnergyFull != energyFull {
		bat.EnergyFull = energyFull
		bat.notifyChange("EnergyFull")
	}
}

func (bat *Battery) setPropPowerNow(powerNow uint64) {
	if bat.PowerNow != powerNow {
		bat.PowerNow = powerNow
		bat.notifyChange("PowerNow")
	}
}

func (bat *Battery) setPropVoltageNow(voltageNow uint64) {
	if bat.VoltageNow != voltageNow {
		bat.VoltageNow = voltageNow
		bat.notifyChange("VoltageNow")
	}
}

func (bat *Battery) setPropPercentage(percentage float64) {
	if bat.Percentage != percentage {
		bat.Percentage = percentage
		bat.notifyChange("Percentage")
	}
}

func (bat *Battery) setPropCapacity(capacity float64) {
	if bat.Capacity != capacity {
		bat.Capacity = capacity
		bat.notifyChange("Capacity")
	}
}

func (bat *Battery) setPropStatus(val battery.Status) {
	if bat.Status != val {
		bat.Status = val
		bat.notifyChange("Status")
	}
}

func (bat *Battery) setPropTimeToEmpty(timeToEmpty uint64) {
	if bat.TimeToEmpty != timeToEmpty {
		bat.TimeToEmpty = timeToEmpty
		bat.notifyChange("TimeToEmpty")
	}
}

func (bat *Battery) setPropTimeToFull(timeToFull uint64) {
	if bat.TimeToFull != timeToFull {
		bat.TimeToFull = timeToFull
		bat.notifyChange("TimeToFull")
	}
}

func (bat *Battery) setPropName(val string) {
	if bat.Name != val {
		bat.Name = val
		bat.notifyChange("Name")
	}
}

func (bat *Battery) setPropTechnology(val string) {
	if bat.Technology != val {
		bat.Technology = val
		bat.notifyChange("Technology")
	}
}

func (bat *Battery) setPropManufacturer(val string) {
	if bat.Manufacturer != val {
		bat.Manufacturer = val
		bat.notifyChange("Manufacturer")
	}
}

func (bat *Battery) setPropModelName(val string) {
	if bat.ModelName != val {
		bat.ModelName = val
		bat.notifyChange("ModelName")
	}
}

func (bat *Battery) setPropSerialNumber(val string) {
	if bat.SerialNumber != val {
		bat.SerialNumber = val
		bat.notifyChange("SerialNumber")
	}
}

func (bat *Battery) setPropEnergyFullDesign(val uint64) {
	if bat.EnergyFullDesign != val {
		bat.EnergyFullDesign = val
		bat.notifyChange("EnergyFullDesign")
	}
}

func (bat *Battery) setPropVoltageMinDesign(val uint64) {
	if bat.VoltageMinDesign != val {
		bat.VoltageMinDesign = val
		bat.notifyChange("VoltageMinDesign")
	}
}
