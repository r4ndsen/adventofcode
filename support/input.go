package support

import (
	"fmt"
	"strings"
)

type InputType string

type Input interface {
	Lines() []string
	Bytes() [][]byte
	Runes() [][]rune
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

func (i InputType) Bytes() [][]byte {
	res := make([][]byte, 0)
	for _, l := range i.Lines() {
		res = append(res, []byte(l))
	}

	return res
}

func (i InputType) Runes() [][]rune {
	res := make([][]rune, 0)
	for _, l := range i.Lines() {
		res = append(res, []rune(l))
	}

	return res
}

func (i InputType) String() string {
	return string(i)
}
