package forth

import (
	"errors"
	"strconv"
	"strings"
)

type word struct {
	numArgs int
	isValid func([]int) bool
	run     func([]int, []int) []int
}

var div = func(in []int) bool { return in[1] != 0 }
var ops = map[string]word{
	"+":    {2, nil, func(s, a []int) []int { return append(s, a[0]+a[1]) }},
	"-":    {2, nil, func(s, a []int) []int { return append(s, a[0]-a[1]) }},
	"*":    {2, nil, func(s, a []int) []int { return append(s, a[0]*a[1]) }},
	"/":    {2, div, func(s, a []int) []int { return append(s, a[0]/a[1]) }},
	"swap": {2, nil, func(s, a []int) []int { return append(s, a[1], a[0]) }},
	"over": {2, nil, func(s, a []int) []int { return append(s, a[0], a[1], a[0]) }},
	"dup":  {1, nil, func(s, a []int) []int { return append(s, a[0], a[0]) }},
	"drop": {1, nil, func(s, a []int) []int { return s }},
}

func Forth(lines []string) ([]int, error) {
	macros := make(map[string]string)
	for i := range lines {
		line := strings.ToLower(lines[i])

		if line[0] == ':' {
			line = line[2 : len(line)-2]
			endMacroName := strings.IndexByte(line, ' ')
			macros[line[:endMacroName]] = expandMacros(line[endMacroName:], macros)
		} else {
			lines[i] = expandMacros(line, macros)
		}
	}

	stack := make([]int, 0, 8)
	for _, word := range strings.Fields(lines[len(lines)-1]) {
		if w, ok := ops[word]; ok {
			if len(stack) < w.numArgs {
				return nil, errors.New("stack underflow")
			}
			argsIdx := len(stack) - w.numArgs
			if w.isValid != nil && !w.isValid(stack[argsIdx:]) {
				return nil, errors.New("invalid arguments")
			}
			stack = w.run(stack[:argsIdx], stack[argsIdx:])
		} else if arg, err := strconv.Atoi(word); err == nil {
			stack = append(stack, arg)
		} else {
			return nil, errors.New("invalid word")
		}
	}
	return stack, nil
}

func expandMacros(line string, macros map[string]string) string {
	for k, v := range macros {
		line = strings.ReplaceAll(line, k, v)
	}
	return line
}
