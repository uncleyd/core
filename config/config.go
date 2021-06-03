package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	cfg = &AdminConfig{}
)

const (
	configPath = "./config/env/%s/admin.json"
)

func Get() *AdminConfig {
	return cfg
}

func Parse(env string) {
	if len(env) < 1 {
		env = "local" // default set to 'local'
		fmt.Errorf("env is null.default set to 'local'\n")
	}

	var path = fmt.Sprintf(configPath, env)

	defer func() {
		fmt.Printf("env config path = %s\n", path)
	}()

	LoadConfig(path)
	fmt.Printf("env=`%s` admin.json load complete.\n", env)
}

func LoadConfig(path string) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("load admin.json file error: %s\n", err)
		panic(err)
		return
	}

	err = json.Unmarshal(bs, cfg)
	if err != nil {
		fmt.Printf("parse admin.json file error: %s\n", err)
		panic(err)
	}
}
