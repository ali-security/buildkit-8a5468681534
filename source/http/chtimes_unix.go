//go:build unix

package http

import (
	"os"
	"time"

	"golang.org/x/sys/unix"
)

// fchtimes sets the access and modification times on the open descriptor of f
// (futimes) rather than re-resolving fallbackPath, closing the symlink-swap
// TOCTOU that os.Chtimes on a path would leave open. mTime is second- or
// HTTP-date-granularity, so futimes' microsecond resolution is sufficient.
// fallbackPath is unused on unix.
func fchtimes(f *os.File, _ string, t time.Time) error {
	tv := unix.NsecToTimeval(t.UnixNano())
	return unix.Futimes(int(f.Fd()), []unix.Timeval{tv, tv})
}
