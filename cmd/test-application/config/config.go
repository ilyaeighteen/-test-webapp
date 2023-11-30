package config

import (
	"fmt"
	logs "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//go:generate mockgen -source=main.go -destination=mocks/mock.go

var Config appConfig

// определяем новый тип для хранения конфигурации, при добавлении новых параметров необходим переопределять!!!
type appConfig struct {
	DB_HOST       string `mapstructure:"db_host"`
	DB_USER       string `mapstructure:"db_user"`
	DB_PASSWD     string `mapstructure:"db_passwd"`
	DB            string `mapstructure:"db"`
	DB_PORT       string `mapstructure:"db_port"`
}

// перечитываем файл конфигурации
func ReloadConfig() error {
    //absolute_path := "C:/Users/Ilya Gulyaev/go/src/test-webapp/cfg"
	absolute_path := "/var/opt/config"
	if config_err := LoadConfig(absolute_path); config_err != nil {
		logs.WithFields(logs.Fields{
			"ReloadConfig": "Open file",
			"PATH":         absolute_path,
		}).Fatal(config_err.Error())
		return config_err
	}
	logs.WithFields(logs.Fields{
		"ReloadConfig": "Open file",
		"PATH":         absolute_path,
	}).Info("Config file Reload!")
	return nil
}

// читаем конфигурацию из файла
func LoadConfig(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("conf") // <- имя конфигурационного файла
	v.SetConfigType("json")
	v.SetEnvPrefix("blueprint")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path) // <- // путь для поиска конфигурационного файла в
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	//парсим json в переменную
	return v.Unmarshal(&Config)
}