package run

import "os"

var DefaultConfig = Config{
	Args:   os.Args,
	Env:    os.Environ(),
	Stderr: os.Stderr,
	Stdin:  os.Stdin,
	Stdout: os.Stdout,
}
