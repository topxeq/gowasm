package gowasm_test

import (
	"embed"
	"io/fs"
	"log"

	"github.com/topxeq/gowasm"
)

//go:embed testdata/index.html
var testdataIndex embed.FS

var moduleConfig gowasm.ModuleConfig

// This example shows how to configure an embed.FS.
func Example_fsConfig() {
	// Strip the embedded path testdata/
	rooted, err := fs.Sub(testdataIndex, "testdata")
	if err != nil {
		log.Panicln(err)
	}

	moduleConfig = gowasm.NewModuleConfig().
		// Make "index.html" accessible to the guest as "/index.html".
		WithFSConfig(gowasm.NewFSConfig().WithFSMount(rooted, "/"))

	// Output:
}
