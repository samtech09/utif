package utif

import (
	"net"
	"net/mail"
)

func IsValidIp(ip string) bool {
	return net.ParseIP(ip) != nil
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
