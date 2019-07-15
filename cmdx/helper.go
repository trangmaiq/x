package cmdx

import (
  "fmt"
  "os"
)

// Fatalf write the message to os.Stderr and exit with code 1
func Fatalf(message string, args ...interface{}) {
	if len(args) > 0 {
	  _, _ = fmt.Fprintf(os.Stderr, message+"\n", args...)
	} else {
	  _, _ = fmt.Fprintf(os.Stderr, message)
	}
	os.Exit(1)
}