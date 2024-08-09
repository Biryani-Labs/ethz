package config

import (
	"path"
)

var (
	HOME_DIR = ""
	ASK_PASS = false
)

func LocateInHomePath(location string) string {
	return path.Join(HOME_DIR, location)
}

func SetHomevar(home string) {
	HOME_DIR = home
}

func SetAskpass(askpass bool) {
	ASK_PASS = askpass
}
