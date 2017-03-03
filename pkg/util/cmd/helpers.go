package cmd

import (
	"fmt"
	"os"
)

// exitWithError will terminate execution with an error result
// It prints the error to stderr and exits with a non-zero exit code
func ExitWithError(err error) {
	fmt.Fprintf(os.Stderr, "\n%v\n", err)
	os.Exit(1)
}
