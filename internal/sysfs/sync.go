//go:build !windows

package sysfs

import (
	"os"

	"github.com/topxeq/gowasm/experimental/sys"
)

func fsync(f *os.File) sys.Errno {
	return sys.UnwrapOSError(f.Sync())
}
