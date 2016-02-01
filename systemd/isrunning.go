package systemd // import "gopkg.in/hlandau/svcutils.v1/systemd"

import "os"

// IsRunningSystemd checks whether the host was booted with systemd as its init
// system. This functions similar to systemd's `sd_booted(3)`: internally, it
// checks whether /run/systemd/system/ exists and is a directory.
// http://www.freedesktop.org/software/systemd/man/sd_booted.html
//
// This comes from github.com/coreos/go-systemd/util, but that package has
// become more complicated and dependent on cgo, so it is duplicated here.
func IsRunning() bool {
	fi, err := os.Lstat("/run/systemd/system")
	if err != nil {
		return false
	}
	return fi.IsDir()
}

// Deprecated.
func IsRunningSystemd() bool {
	return IsRunning()
}

// Technically this is 'is running under systemd with the notify service type'.
// The purpose of this is to allow a daemon to determine whether it should
// behave in a systemd-like way, though, so that's fine.
func IsRunningUnder() bool {
	return IsRunning() && os.Getenv("NOTIFY_SOCKET") != ""
}

// © 2015 CoreOS, Inc.    Apache 2.0 License
