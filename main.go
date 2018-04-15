package main

import (
	"os"
	"path/filepath"
	"strings"
)

const boxPathEnvKey = "__BOX_ACTIVATED_PATH"

type box struct {
	activatedPath string
	pwd           string
	shell         shell
}

func main() {
	box := newBox()
	box.execute()
}

func newBox() box {
	s := shell{}
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		s.echof("error getting current directory: %v", err)
		os.Exit(1)
	}
	return box{
		activatedPath: os.Getenv(boxPathEnvKey),
		pwd:           pwd,
		shell:         s,
	}
}

func (b *box) execute() {
	if b.activatedPath != "" {
		if strings.HasPrefix(b.pwd, b.activatedPath) {
			// Inside activated.
			return
		}
		// Outside activated.
		c := config{b.activatedPath}
		b.deactivate(c.getVirtualEnvironmentMap())
		return
	}
	c := config{b.pwd}
	b.activate(c.getVirtualEnvironmentMap())
}

func (b *box) activate(venv map[string]string) {
	if venv == nil {
		return
	}
	b.shell.export(boxPathEnvKey, b.activatedPath)
	b.shell.echof("activated %v", venv)
}

func (b *box) deactivate(venv map[string]string) {
	b.activatedPath = ""
	b.shell.unset(boxPathEnvKey)
	if venv != nil {
		b.shell.echof("deactivated %v", venv)
	}
}
