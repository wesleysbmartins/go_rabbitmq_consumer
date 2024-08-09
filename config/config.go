package configs

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct{}

type IConfig interface {
	Load(entity *interface{})
}

func (c *Config) Load(fileWordKey string, entity interface{}) {
	dir, _ := os.Getwd()

	configPath := fmt.Sprintf("%s/config/files/", strings.Replace(dir, "cmd", "", 1))

	dirResponse, err := os.ReadDir(configPath)

	var files []byte

	for _, filebyte := range dirResponse {
		filename := filebyte.Name()

		if strings.Contains(filename, ".yml") && strings.Contains(filename, fileWordKey) {
			file, err := os.ReadFile(configPath + filename)

			if err != nil {
				panic("Error: Failed to read .yml configuration files")
			} else {
				files = append(files, file...)
			}
		}
	}

	if err != nil {
		fmt.Println("Error: Failed to read config file: " + err.Error())
	}

	yaml.Unmarshal(files, entity)
}
