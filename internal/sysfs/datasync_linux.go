package sysfs

import (
	"os"
	"syscall"

	"github.com/topxeq/gowasm/experimental/sys"
)

func datasync(f *os.File) sys.Errno {
	return sys.UnwrapOSError(syscall.Fdatasync(int(f.Fd())))
}
