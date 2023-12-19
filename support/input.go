package support

import (
	"fmt"
	"strings"
)

type InputType string

type Input interface {
	Lines() []string
	fmt.Stringer
}

var _ Input = new(InputType)

func (i InputType) Lines() []string {

	res := strings.Split(string(i), "\n")
	if res[len(res)-1] == "" {
		return res[:len(res)-1]
	}

	return res
}

func (i InputType) String() string {
	return string(i)
}
