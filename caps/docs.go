// Package caps provides functions for controlling capabilities on Linux. On
// other OSes, the functions are no-ops.
package caps

// This constant will be true iff the target platform supports capabilities.
const Supported = supported

// Returns true iff there are any capabilities available to the program.
// Returns false on non-Linux OSes.
func HaveAny() bool {
	return haveAny()
}

// Attempt to drop all capabilities. Does nothing on non-Linux OSes.
func Drop() error {
	return drop()
}
