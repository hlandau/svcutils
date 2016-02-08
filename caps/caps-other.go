// +build !linux linux,!cgo

package caps

const supported = false

func haveAny() bool {
	return false
}

func drop() error {
	return nil
}
