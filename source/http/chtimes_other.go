//go:build !unix

package http

import (
	"os"
	"time"
)

// fchtimes falls back to a path-based update on platforms without a
// descriptor-based utimes. The fd-scoped TOCTOU hardening is unix-only; the
// ownership change in save() (f.Chown) is also a no-op equivalent off unix.
func fchtimes(_ *os.File, fallbackPath string, t time.Time) error {
	return os.Chtimes(fallbackPath, t, t)
}
