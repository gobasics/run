package run

import (
	"errors"
	"fmt"
)

type Runner map[string]Runnable

func (r Runner) Run(c Container) error {
	if len(c.Args) == 0 {
		return errors.New(ArgsErr)
	}

	fn, ok := r[c.Args[0]]
	if !ok {
		return fmt.Errorf(ImplErr, c.Args[0])
	}

	return fn(c)
}
