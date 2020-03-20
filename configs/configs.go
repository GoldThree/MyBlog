package configs

import (
	"bytes"
	"errors"
	"io/ioutil"

	"github.com/spf13/viper"
)

// InitConfigs 初始化配置
func InitConfigs() error {

	viper.SetConfigType("yaml")

	appEnv := "./configs/development.yaml"

	conf, err := ioutil.ReadFile(appEnv)
	if err != nil {
		return errors.New("tonic_error.configs.missing_configs_file")
	}

	err = viper.ReadConfig(bytes.NewBuffer(conf))
	if err != nil {
		return errors.New("configs_error.configs.invalid_format")
	}
	return nil
}

// Get 从配置文件中获取interface类型的value
func Get(key string) interface{} {
	return viper.Get(key)
}

// GetString 从配置文件中获取string类型的value
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt 从配置文件中获取int类型的value
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetBool 从配置文件中获取bool类型的value
func GetBool(key string) bool {
	return viper.GetBool(key)
}
