// +build !windows,!plan9

package chroot

import (
	"fmt"
	"path/filepath"
	"strings"
	"syscall"
)

var anchor = "/"

func getAnchor() string {
	return anchor
}

func rel(path string) (chrootRelativePath string, canAddress bool) {
	p, err := filepath.Rel(Anchor(), path)
	if err != nil || p == "" {
		return "", false
	}

	if p == ".." || strings.HasPrefix(p, "../") {
		return "", false
	}

	if p == "." {
		return "/", true
	}

	return "/" + p, true
}

// Chroots into a directory. The package keeps track of the resultant chroot
// path, even after multiple successive chroot operations. The anchor is
// updated if the operation is successful.
func Chroot(path string) error {
	if !filepath.IsAbs(path) {
		return fmt.Errorf("chroot path must be absolute: %v", path)
	}

	err := syscall.Chroot(path)
	if err != nil {
		return err
	}

	anchor = filepath.Join(anchor, path[1:])
	return nil
}
