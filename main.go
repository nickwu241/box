package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

const boxPathEnvKey = "__BOX_ACTIVATED_PATH"

type state struct {
	activatedPath string
	venv          map[string]string
}

func main() {
	pwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Printf("echo 'error getting current directory: %v';\n", err)
		os.Exit(1)
	}
	box(pwd)
}

func box(pwd string) {
	s, stateExisted := newState(pwd)
	if stateExisted {
		if strings.HasPrefix(pwd, s.activatedPath) {
			// fmt.Println("echo 'activated already';")
			return
		}
		s.deactivate()
		return
	}
	s.activate()
}

func newState(pwd string) (state, bool) {
	ap := os.Getenv(boxPathEnvKey)
	if ap != "" {
		return state{
			activatedPath: ap,
			venv:          getVenv(ap),
		}, true
	}

	return state{
		activatedPath: pwd,
		venv:          getVenv(pwd),
	}, false
}

func getVenv(configPath string) map[string]string {
	if !configFileExists(configPath) {
		// fmt.Printf("echo 'config doesnt exist in %s';\n", configPath)
		return nil
	}
	viper.SetConfigName("box")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("echo 'reading box config: %v';\n", err)
		os.Exit(1)
	}

	venv := map[string]string{}
	goVersion := viper.GetString("go")
	if goVersion != "" {
		venv["go"] = goVersion
	}
	pythonVersion := viper.GetString("python")
	if pythonVersion != "" {
		venv["python"] = pythonVersion
	}
	if len(venv) == 0 {
		return nil
	}

	// fmt.Printf("echo 'getVenv: %v';\n", venv)
	return venv
}

func configFileExists(path string) bool {
	if _, err := os.Stat(filepath.Join(path, "box.yml")); !os.IsNotExist(err) {
		return true
	}
	if _, err := os.Stat(filepath.Join(path, "box.yaml")); !os.IsNotExist(err) {
		return true
	}
	return false
}

func (s *state) activate() {
	if s.venv == nil {
		// fmt.Println("echo 'nothing to activate';")
		return
	}
	fmt.Printf("export %s='%s';\n", boxPathEnvKey, s.activatedPath)
	fmt.Printf("echo activated '%v';\n", s.venv)
}

func (s *state) deactivate() {
	fmt.Printf("unset '%s';\n", boxPathEnvKey)
	fmt.Printf("echo deactivated '%v';\n", s.venv)
	s.activatedPath = ""
	s.venv = nil
}
