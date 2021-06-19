package config

import (
	"log"
	_"os"
     _"path/filepath"
	"github.com/spf13/viper"
)

var Config *Configuration

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type DatabaseConfiguration struct {
	Driver       string
	Dbname       string
	Username     string
	Password     string
	Host         string
	Port         string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
}

type ServerConfiguration struct {
	Port   string
	Secret string
	Mode   string
}

// SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Configuration

	viper.SetConfigName("config.yml")
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	/*
    path ,err1 := os.Getwd()
		if err1 != nil{
	}
	log.Println("path",path)
		err2 := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				log.Println(err)	
			}
			log.Println(path, info.Size(), info.IsDir(),info.Name())
			return nil
		})
		if err2 != nil{
			log.Println(err2)
		}

		*/
	if err := viper.ReadInConfig(); err != nil {
		
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	Config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
