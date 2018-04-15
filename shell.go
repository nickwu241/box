package main

import (
	"fmt"
)

type shellCLI struct{}

func (s *shellCLI) echof(format string, a ...interface{}) {
	fmt.Printf("echo '%s';\n", fmt.Sprintf(format, a...))
}

func (s *shellCLI) export(key, value string) {
	fmt.Printf("export %s='%s';\n", key, value)
}

func (s *shellCLI) unset(key string) {
	fmt.Printf("unset '%s';\n", key)
}
