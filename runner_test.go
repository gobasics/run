package run

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

func TestRunnerRun(t *testing.T) {
	for k, v := range []struct {
		args []string
		want error
	}{
		{[]string{}, errors.New(ArgsErr)},
		{[]string{"bar"}, fmt.Errorf(ImplErr, "bar")},
		{[]string{"foo"}, nil},
	} {
		t.Run(strconv.Itoa(k), func(t *testing.T) {
			m := Runner{
				"foo": func(Container) error { return nil },
			}

			got := m.Run(Container{Config: Config{Args: v.args}})
			if (got == nil) != (v.want == nil) {
				t.Errorf("want %v, got %v", v.want, got)
			}
		})
	}
}
