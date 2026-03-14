package main

import (
	"testing"

	"github.com/topxeq/gowasm/internal/testing/maintester"
	"github.com/topxeq/gowasm/internal/testing/require"
)

// Test_main ensures the following will work:
//
//	go run greet.go gowasm
func Test_main(t *testing.T) {
	stdout, _ := maintester.TestMain(t, main, "greet", "gowasm")
	require.Equal(t, `wasm >> Hello, gowasm!
go >> Hello, gowasm!
`, stdout)
}
