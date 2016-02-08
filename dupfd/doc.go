// Package dupfd makes the Dup2 system call uniformly available on *NIX
// platforms.
//
// The call is not available on Solaris, but an alternate mechanism is
// available to obtain the same effect. The call has been replaced with
// dup3 on Linux, and dup2 is not available in the legacy-free arm64
// port of Linux.
//
// This package provides a uniform Dup2 function for all of these platforms.
// *NIX only.
package dupfd
