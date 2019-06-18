package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type configBridge struct {
	Backend    string `yaml:"backend"`
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
}

type deliveryChannelConfig struct {
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

type configFile struct {
	Bridge   configBridge                     `yaml:"bridge"`
	Delivery map[string]deliveryChannelConfig `yaml:"delivery"`
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
