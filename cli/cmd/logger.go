package cmd

import (
	"fmt"
	"os"
)

// Failf .
func Failf(msg string, args ...interface{}) {
	Exitf(1, msg, args...)
}

// Warnf .
func Warnf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
}

// IfErrf .
func IfErrf(err error) {
	if err != nil {
		Warnf("Error %v", err)
	}
}

// Infof .
func Infof(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stdout, msg+"\n", args...)
}

// Exitf .
func Exitf(code int, msg string, args ...interface{}) {
	if code == 0 {
		fmt.Fprintf(os.Stdout, msg+"\n", args...)
	} else {
		fmt.Fprintf(os.Stderr, msg+"\n", args...)
	}
	os.Exit(code)
}
