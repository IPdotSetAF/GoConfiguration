package configuration

import (
	"encoding/json"
	"os"

	"github.com/mitchellh/mapstructure"
)

type IConfigurationManager interface {
	GetSection(key string) (interface{}, error)
}

type ConfigurationManager struct {
	FileName string
	Config   map[string]interface{}
}

func NewConfigurationManager(fileName string) (*ConfigurationManager, error) {
	if fileName == "" {
		fileName = "appSettings.json"
	}

	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var config map[string]interface{}
	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		return nil, err
	}

	return &ConfigurationManager{
		FileName: fileName,
		Config:   config,
	}, nil
}

func (instance *ConfigurationManager) GetSection(key string) interface{} {
	if key == "" {
		return instance.Config
	}
	return instance.Config[key]
}

func (instance *ConfigurationManager) GetOptions(key string, options any) error {
	if key == "" {
		return mapstructure.Decode(instance.Config, &options)
	}
	return mapstructure.Decode(instance.Config[key], &options)
}
