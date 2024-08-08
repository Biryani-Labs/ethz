package utils

import (
	"strings"
	"syscall"

	"github.com/Biryani-Labs/ezeth/common/logs"
	"github.com/pterm/pterm"
	"golang.org/x/term"
)

func ReadPassword() string {
	pterm.Print("Enter your password : ")
	password, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		logs.Error(err, "Unable to read password")
		return ""
	}

	return strings.TrimSpace(string(password))
}
