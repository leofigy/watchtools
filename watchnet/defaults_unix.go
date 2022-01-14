//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || linux || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd js,wasm linux netbsd openbsd solaris

package watchnet

const (
	NEWLINE     = "\n"
	lsofCommand = "lsof"
)

var (
	defaultArgs = []string{"-i", "-P", "-n"}
)
