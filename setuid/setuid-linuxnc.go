// +build linux,!cgo,!go1.16

package setuid

import "fmt"

var errNoSetuid = fmt.Errorf("set*id calls are not supported on Linux when built with cgo disabled unless using Go 1.16 or later")

func setuid(uid int) error {
	return errNoSetuid
}

func setgid(gid int) error {
	return errNoSetuid
}

func setresgid(rgid, egid, sgid int) error {
	return errNoSetuid
}

func setresuid(ruid, euid, suid int) error {
	return errNoSetuid
}
