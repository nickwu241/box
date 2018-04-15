package main

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type config struct {
	dir string
}

func (c *config) getVirtualEnvironmentMap() map[string]string {
	if !c.configFileExists() {
		return nil
	}
	viper.SetConfigName("box")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(c.dir)
	if err := viper.ReadInConfig(); err != nil {
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

	return venv
}

func (c *config) configFileExists() bool {
	if _, err := os.Stat(filepath.Join(c.dir, "box.yml")); !os.IsNotExist(err) {
		return true
	}
	if _, err := os.Stat(filepath.Join(c.dir, "box.yaml")); !os.IsNotExist(err) {
		return true
	}
	return false
}
