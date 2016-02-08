// +build !windows,!plan9

package pidfile

import (
	"fmt"
	"io"
	"os"
	"sync"
	"syscall"
)

var (
	pidFile     *os.File
	pidFileName string
)

type pidfile struct {
	once sync.Once
	f    *os.File
	path string
}

func (p *pidfile) Close() error {
	p.once.Do(func() {
		// Try and remove file, don't care if it fails.
		os.Remove(p.path)

		p.f.Close()
		p.f = nil
	})

	return nil
}

// Opens and locks a file and writes the current PID to it. The file is kept open
// until you close the returned interface, at which point it is deleted. It may also
// be deleted if the program exits without closing the returned interface.
func Open(path string) (io.Closer, error) {
	return OpenWith(path, fmt.Sprintf("%d\n", os.Getpid()))
}

// Opens and locks a file and writes body to it. The file is kept open until
// you close the returned interface, at which point it is deleted. It may also
// be deleted if the program exits without closing the returned interface.
func OpenWith(path, body string) (io.Closer, error) {
	f, err := open(path)
	if err != nil {
		return nil, err
	}

	_, err = f.WriteString(body)
	if err != nil {
		f.Close()
		return nil, err
	}

	return &pidfile{
		f:    f,
		path: path,
	}, nil
}

func open(path string) (*os.File, error) {
	var f *os.File
	var err error

	for {
		f, err = os.OpenFile(path,
			syscall.O_RDWR|syscall.O_CREAT|syscall.O_EXCL, 0644)
		if err != nil {
			if !os.IsExist(err) {
				return nil, err
			}

			f, err = os.OpenFile(path, syscall.O_RDWR, 0644)
			if err != nil {
				if os.IsNotExist(err) {
					continue
				}
				return nil, err
			}
		}

		err = syscall.FcntlFlock(f.Fd(), syscall.F_SETLK, &syscall.Flock_t{
			Type: syscall.F_WRLCK,
		})
		if err != nil {
			f.Close()
			return nil, err
		}

		st1 := syscall.Stat_t{}
		err = syscall.Fstat(int(f.Fd()), &st1) // ffs
		if err != nil {
			f.Close()
			return nil, err
		}

		st2 := syscall.Stat_t{}
		err = syscall.Stat(path, &st2)
		if err != nil {
			f.Close()

			if os.IsNotExist(err) {
				continue
			}

			return nil, err
		}

		if st1.Ino != st2.Ino {
			f.Close()
			continue
		}

		break
	}

	return f, nil
}
