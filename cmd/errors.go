package cmd

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

var (
	NetworkError = errors.New(toBoldRed("NETWORK ERROR: Make sure you have internet connection."))
	InputError   = errors.New(toBoldRed("INPUT ERROR: Make sure you entered a correct type argument."))
)

func toBoldRed(text string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", text)
}

func isNetworkError(err error) bool {
	if _, ok := err.(net.Error); ok {
		return true
	}
	return false
}

func isStrconvError(err error) bool {
	if _, ok := err.(*strconv.NumError); ok {
		return true
	}
	return false
}

func filterError(err error) error {
	switch {
	case isNetworkError(err):
		return NetworkError
	case isStrconvError(err):
		return InputError
	default:
		return err
	}
}
