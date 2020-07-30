package run

import "io"

type Config struct {
	Args   []string
	Env    string
	Stderr io.Writer
	Stdin  io.Reader
	Stdout io.Writer
}
