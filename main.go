package main

import (
	"syscall"
)

var (
	mcp2210, _ = syscall.LoadLibrary("mcp2210_dll_um_x64.dll")
)

func main() {
}