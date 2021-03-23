# Utilities for writing services in Go

[![godocs.io](https://godocs.io/gopkg.in/hlandau/svcutils.v2?status.svg)](https://godocs.io/gopkg.in/hlandau/svcutils.v2) [![No modules](https://www.devever.net/~hl/f/no-modules2.svg) 100% modules-free.](https://www.devever.net/~hl/gomod)

The following packages are contained in this repository:

  * chroot, a package for chrooting and then determining whether absolute paths
    can be addressed within that chroot and, if so, converting them
    appropriately;

  * exepath, a package for determining the absolute path of the executable as
    invoked, but without resolving symlinks, which can be useful in some
    circumstances;

  * passwd, a package for determining user and group information on \*NIX
    systems beyond that available in `os/user`;

  * pidfile, a package for creating and locking PID files on \*NIX;

  * setuid, a package for changing UID and GID on \*NIX systems,
    including workarounds for the unfortunate absurdities underlying
    Linux's implementation of setuid (which means that `syscall.Setuid`
    does not work on Linux);

  * systemd, a package for detecting whether systemd is in use and sending
    status messages to it, in a way that works in a chroot;

  * caps, a package for detecting and dropping capabilities on Linux;

  * dupfd, a package for duplicating file descriptors to a target
    file descriptor number, which irons out some differences between
    different \*NIX platforms.

## Licence

    © 2015—2016 Hugo Landau <hlandau@devever.net>

[Licenced under the licence with SHA256 hash
`fd80a26fbb3f644af1fa994134446702932968519797227e07a1368dea80f0bc`, a copy of
which can be found
here.](https://github.com/hlandau/rilts/blob/master/licences/COPYING.MIT)
