// +build !windows,!plan9

package dupfd

// Duplicates an FD to a target FD. See dup2(2).
func Dup2(sourceFD, targetFD int) error {
	return dup2(sourceFD, targetFD)
}
