// +build !windows

package userutil

import (
	"strings"
	"runtime"
	"os/user"
)

// IsRoot returns a nil error if the current user is root.
func IsRoot() error {
	currentUser, err := user.Current()
	if err != nil {
		// For whatever reason, 'user.Current()' throws a "not implemented
		// error" when running as a launch daemon on macOS.
		if runtime.GOOS == "darwin" && strings.Contains(err.Error(), "Current not implemented on") {
			return nil
		}
		return UserError{
			reason:        "Failed to check if current user is root - " + err.Error(),
			unableToCheck: true,
		}
	}

	if currentUser.Username != "root" {
		return UserError{
			reason:  "The current user is not 'root'",
			notRoot: true,
		}
	}

	return nil
}
