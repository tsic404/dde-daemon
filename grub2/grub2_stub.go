package main

import (
	"dlib/dbus"
	"fmt"
	"strconv"
)

const (
	_GRUB2_DEST = "com.deepin.daemon.Grub2"
	_GRUB2_PATH = "/com/deepin/daemon/Grub2"
	_GRUB2_IFC  = "com.deepin.daemon.Grub2"
)

func (grub *Grub2) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		_GRUB2_DEST,
		_GRUB2_PATH,
		_GRUB2_IFC,
	}
}

func (grub *Grub2) Init() {
	// TODO
}

func (grub *Grub2) Load() {
	// TODO
	grub.readEntries()
	grub.readSettings()
}

func (grub *Grub2) Save() {
	// TODO
	grub.writeSettings()
}

func (grub *Grub2) GetEntries() []string {
	return grub.entries // TODO
}

func (grub *Grub2) SetDefaultEntry(index uint32) {
	indexStr := strconv.FormatInt(int64(index), 10)
	grub.settings["GRUB_DEFAULT"] = indexStr
}

func (grub *Grub2) GetDefaultEntry() uint32 {
	index, err := strconv.ParseInt(grub.settings["GRUB_DEFAULT"], 10, 32)
	if err != nil {
		logError(fmt.Sprintf(`valid value, settings["GRUB_DEFAULT"]=%s`, grub.settings["GRUB_DEFAULT"])) // TODO
		return 0
	}
	return uint32(index)
}

func (grub *Grub2) SetTimeout(timeout int32) {
	timeoutStr := strconv.FormatInt(int64(timeout), 10)
	grub.settings["GRUB_TIMEOUT"] = timeoutStr
}

func (grub *Grub2) GetTimeout() int32 {
	timeout, err := strconv.ParseInt(grub.settings["GRUB_TIMEOUT"], 10, 32)
	if err != nil {
		logError(fmt.Sprintf(`valid value, settings["GRUB_TIMEOUT"]=%s`, grub.settings["GRUB_TIMEOUT"])) // TODO
		return 5
	}
	return int32(timeout)
}

func (grub *Grub2) SetGfxmode(gfxmode string) {
	grub.settings["GRUB_GFXMODE"] = gfxmode
}

func (grub *Grub2) GetGfxmode() string {
	return grub.settings["GRUB_GFXMODE"]
}

func (grub *Grub2) SetBackground(imageFile string) {
	grub.settings["GRUB_BACKGROUND"] = imageFile
}

func (grub *Grub2) GetBackground() string {
	return grub.settings["GRUB_BACKGROUND"]
}

func (grub *Grub2) SetTheme(themeFile string) {
	grub.settings["GRUB_THEME"] = themeFile
}

func (grub *Grub2) GetTheme() string {
	return grub.settings["GRUB_THEME"]
}
