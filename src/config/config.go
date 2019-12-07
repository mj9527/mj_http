package config

import (
	"fmt"
	toml "github.com/BurntSushi/toml"
	"path/filepath"
)

type CustomConfig struct {
	ServerInfo struct {
		Port    int
		Timeout int
	}
}

var ServerConfig CustomConfig // global config

func init() {
	filePath, err := filepath.Abs("../config/config_test.toml")
	if err != nil {
		fmt.Println("can't find file ", err)
		return
	}
	if _, err := toml.DecodeFile(filePath, &ServerConfig); err != nil {
		panic(err)
	}
	fmt.Println("reuslt ", ServerConfig)
}
