package config

import "github.com/spf13/viper"

// InitConfig 初始化配置文件
func InitConfig(fileName, fileType, filePath string) error {
	// config file
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	// config file path
	viper.AddConfigPath(filePath)
	// try to load config file
	err := viper.ReadInConfig()
	return err
}
