package utils

import (
	"fmt"
	"os"
)

func ErrExit(err error, prepend string) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", prepend, err)
	os.Exit(1)
}
