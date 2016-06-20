package exit

import (
	"fmt"
	"os"
)

var (
	// UsePanic indicates whether to use the Panic method to process the Exit call
	UsePanic = false
)

// Error is used by panic
type Error int

// Error returns a string in format "exit status %d"
func (err Error) Error() string {
	return fmt.Sprintf("exit status %d", err)
}

// Exit the program with status code
func Exit(code int) {
	if UsePanic {
		panic(Error(code))
	}
	os.Exit(code)
}
