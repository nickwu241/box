package main

import (
	"fmt"
)

type shell struct{}

func (s *shell) echof(format string, a ...interface{}) {
	fmt.Printf("echo '%s';\n", fmt.Sprintf(format, a...))
}

func (s *shell) export(key, value string) {
	fmt.Printf("export %s='%s';\n", key, value)
}

func (s *shell) unset(key string) {
	fmt.Printf("unset '%s';\n", key)
}
