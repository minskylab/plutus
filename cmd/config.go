package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type backendConfig map[string]string

type configBridge struct {
	Backend    string                   `yaml:"backend"` // DEPRECATE IN THE FUTURE
	Backends   map[string]backendConfig `yaml:"backends"`
	PublicKey  string                   `yaml:"public_key"`  // DEPRECATE IN THE FUTURE
	PrivateKey string                   `yaml:"private_key"` // DEPRECATE IN THE FUTURE
}

type deliveryChannelConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
	Port     int64  `yaml:"port"`
}

type companyConfig struct {
	Name         string                 `yaml:"name"`
	OfficialWeb  string                 `yaml:"official_web"`
	SupportEmail string                 `yaml:"support_email"`
	SupportPhone string                 `yaml:"support_phone"`
	Custom       map[string]interface{} `yaml:"custom"`
}

type configFile struct {
	Port     int64                            `yaml:"port"`
	Bridge   configBridge                     `yaml:"bridge"`
	Delivery map[string]deliveryChannelConfig `yaml:"delivery"`
	Company  companyConfig                    `yaml:"company"`
}

func readConfigFile(filename string) (*configFile, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := new(configFile)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
