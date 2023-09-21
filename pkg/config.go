package pkg

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	//Ak/sk of cloud providers,Obtain the mirror list on the source side
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	//Cloud Provider Region
	RegionAli string `yaml:"region_ali"`
	RegionHw  string `yaml:"region_hw"`
	//Aliyun ACR Account
	UserAli   string `yaml:"user_ali"`
	PasswdAli string `yaml:"passwd_ali"`
	//Huawei SWR Account
	UserHw   string `yaml:"user_hw"`
	PasswdHw string `yaml:"passwd_hw"`
}

var config *Config

func ReadConfigFromFile(path string) (*Config, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	yaml.Unmarshal(file, &config)
	return config, nil
}
