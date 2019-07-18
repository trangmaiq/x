package cmdx

import (
	"fmt"
	"os"
)

// Must fatals with the optional message if err is not nil
func Must(err error, message string, args ...interface{}) {
	if err == nil {
		return
	}

	_, _ = fmt.Fprintf(os.Stderr, message+"\n", args)
	os.Exit(1)
}

// Fatalf write the message to os.Stderr and exit with code 1
func Fatalf(message string, args ...interface{}) {
	if len(args) > 0 {
		_, _ = fmt.Fprintf(os.Stderr, message+"\n", args...)
	} else {
		_, _ = fmt.Fprintf(os.Stderr, message)
	}
	os.Exit(1)
}
