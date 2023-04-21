package args

import (
	"fmt"
)

type Arguments struct {
	Args []string
}

func New(args []string) *Arguments {
	if len(args) == 0 {
		return nil
	}

	return &Arguments{Args: args}
}

func (a *Arguments) String() string {
	return fmt.Sprintf("%q", a.Args)
}

func (a *Arguments) IndexOf(l ...string) int {
	for i, v := range a.Args {
		for _, s := range l {
			if s == v {
				return i
			}
		}
	}

	return -1
}

func (a *Arguments) AllIndexOf(l ...string) []int {
	in := []int{}
	for i, v := range a.Args {
		for _, s := range l {
			if s == v {
				in = append(in, i)
			}
		}
	}

	return in
}
