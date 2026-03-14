package sysfs_test

import (
	"io/fs"
	"testing/fstest"

	"github.com/topxeq/gowasm"
	"github.com/topxeq/gowasm/experimental/sysfs"
)

var moduleConfig gowasm.ModuleConfig

// This example shows how to adapt a fs.FS to a sys.FS
func ExampleAdaptFS() {
	m := fstest.MapFS{
		"a/b.txt": &fstest.MapFile{Mode: 0o666},
		".":       &fstest.MapFile{Mode: 0o777 | fs.ModeDir},
	}
	root := &sysfs.AdaptFS{FS: m}

	moduleConfig = gowasm.NewModuleConfig().
		WithFSConfig(gowasm.NewFSConfig().(sysfs.FSConfig).WithSysFSMount(root, "/"))

	// Output:
}

// This example shows how to configure a sysfs.DirFS
func ExampleDirFS() {
	root := sysfs.DirFS(".")

	moduleConfig = gowasm.NewModuleConfig().
		WithFSConfig(gowasm.NewFSConfig().(sysfs.FSConfig).WithSysFSMount(root, "/"))

	// Output:
}

// This example shows how to configure a sysfs.ReadFS
func ExampleReadFS() {
	root := sysfs.DirFS(".")
	readOnly := &sysfs.ReadFS{FS: root}

	moduleConfig = gowasm.NewModuleConfig().
		WithFSConfig(gowasm.NewFSConfig().(sysfs.FSConfig).WithSysFSMount(readOnly, "/"))

	// Output:
}
