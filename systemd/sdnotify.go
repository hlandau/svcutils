// +build !windows,!plan9

package systemd

import (
	"errors"
	"net"
	"os"
	"sync"
)

// Error returned if no systemd notify protocol socket can be found.
//
// This is an indication that the service is not running under systemd or
// Type=notify is not set in the systemd unit file.
var ErrNoSocket = errors.New("No socket")

// sdNotifySocket
var sdNotifyMutex sync.Mutex
var sdNotifySocket *net.UnixConn
var sdNotifyInited bool

// Send sends a message to the init daemon. It is common to ignore the error.
//
// This function differs from that in github.com/coreos/go-systemd/daemon in
// that that code closes the socket after each call, and so won't work in a
// chroot. This function keeps the socket open; so long as it is called at
// least once before chrooting, it can continue to be used.
//
// May return ErrNoSocket.
func NotifySend(state string) error {
	sdNotifyMutex.Lock()
	defer sdNotifyMutex.Unlock()

	if !sdNotifyInited {
		sdNotifyInited = true

		socketAddr := &net.UnixAddr{
			Name: os.Getenv("NOTIFY_SOCKET"),
			Net:  "unixgram",
		}

		if socketAddr.Name == "" {
			return ErrNoSocket
		}

		conn, err := net.DialUnix(socketAddr.Net, nil, socketAddr)
		if err != nil {
			return err
		}

		sdNotifySocket = conn
	}

	if sdNotifySocket == nil {
		return ErrNoSocket
	}

	_, err := sdNotifySocket.Write([]byte(state))
	return err
}

// © 2015 CoreOS, Inc.    Apache 2.0 License
