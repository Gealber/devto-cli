package cmd

import (
	"errors"
	"fmt"
	"net"
	"strconv"
)

var (
	NetworkTimeoutError   = errors.New(toBoldRed("TIMEOUT NETWORK ERROR: Make sure you have internet connection."))
	NetworkTemporaryError = errors.New(toBoldRed("TEMPORARY NETWORK ERROR: Make sure you have internet connection."))
	InputError            = errors.New(toBoldRed("INPUT ERROR: Make sure you entered a correct type argument."))
	ApiKeyMissing         = errors.New(toBoldRed("API_KEY is empty, provide one with the auth command"))
)

func toBoldRed(text string) string {
	return fmt.Sprintf("\033[1;31m%s\033[0m", text)
}

func isNetworkTimeoutError(err error) bool {
	if nErr, ok := err.(net.Error); ok && nErr.Timeout() {
		return true
	}
	return false
}

func isNetworkTemporaryError(err error) bool {
	if nErr, ok := err.(net.Error); ok && nErr.Temporary() {
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
	case isNetworkTemporaryError(err):
		return NetworkTemporaryError
	case isNetworkTimeoutError(err):
		return NetworkTimeoutError
	case isStrconvError(err):
		return InputError
	default:
		return err
	}
}
