package forth

import (
	"fmt"
	"strconv"
	"strings"
)

func isNumber(c byte) bool {
	return '0' <= c && c <= '9'
}

type forth struct {
	stack   []int
	command map[string][]string
}

var words = []struct {
	name string
	n    int
	f    func(*forth, []int) error
}{
	{"+", 2, func(f *forth, v []int) error { f.push(v[0] + v[1]); return nil }},
	{"-", 2, func(f *forth, v []int) error { f.push(v[0] - v[1]); return nil }},
	{"*", 2, func(f *forth, v []int) error { f.push(v[0] * v[1]); return nil }},
	{"/", 2, func(f *forth, v []int) error {
		if v[1] == 0 {
			return fmt.Errorf("divide by zero")
		}
		f.push(v[0] / v[1])
		return nil
	}},
	{"dup", 1, func(f *forth, v []int) error { f.push(v[0], v[0]); return nil }},
	{"drop", 1, func(f *forth, v []int) error { return nil }},
	{"swap", 2, func(f *forth, v []int) error { f.push(v[1], v[0]); return nil }},
	{"over", 2, func(f *forth, v []int) error { f.push(v[0], v[1], v[0]); return nil }},
}

func (f *forth) push(values ...int) {
	f.stack = append(f.stack, values...)
}

func (f *forth) pop(n int) ([]int, error) {
	if len(f.stack) < n {
		return nil, fmt.Errorf("stack underflow")
	}
	vals := f.stack[len(f.stack)-n:]
	f.stack = f.stack[:len(f.stack)-n]
	return vals, nil
}

func (f *forth) define(input string) error {
	parts := strings.Split(input, " ")

	if isNumber(parts[1][0]) {
		return fmt.Errorf("can not redefine numbers")
	}

	var subcmds []string
	for _, part := range parts[2 : len(parts)-1] {
		if cmds := f.command[part]; cmds != nil {
			subcmds = append(subcmds, cmds...)
		} else {
			subcmds = append(subcmds, part)
		}
	}

	f.command[parts[1]] = subcmds

	return nil
}

func (f *forth) eval(input string) error {
	if isNumber(input[0]) {
		val, _ := strconv.Atoi(input)
		f.push(val)
		return nil
	} else if cmds := f.command[input]; cmds != nil {
		for _, cmd := range cmds {
			if err := f.eval(cmd); err != nil {
				return err
			}
		}
		return nil
	}

	for _, word := range words {
		if input == word.name {
			val, err := f.pop(word.n)
			if err != nil {
				return err
			}
			return word.f(f, val)
		}
	}

	return fmt.Errorf("unknown word %s", input)
}

// Forth runs the input and returns the corresponding stack
func Forth(inputs []string) ([]int, error) {
	f := forth{command: make(map[string][]string)}

	for _, input := range inputs {
		input = strings.ToLower(input)
		if input[0] == ':' {
			if err := f.define(input); err != nil {
				return nil, err
			}
		} else {
			for _, cmd := range strings.Split(input, " ") {
				if err := f.eval(cmd); err != nil {
					return nil, err
				}
			}
		}
	}

	return f.stack, nil
}