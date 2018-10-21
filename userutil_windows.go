package userutil

import (
	"os"
)

// IsRoot returns a nil error if the current process is running at the
// Administrator level.
func IsRoot() error {
	f, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		return UserError{
			reason:  "The current user is not running as Administrator",
			notRoot: true,
		}
	}

	f.Close()

	return nil
}
