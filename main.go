package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	defaults = map[string]interface{}{
		"username": "admin",
		"password": "password1",
		"host":     "localhost",
		"port":     3306,
		"database": "test",
	}
	configPaths = []string{
		".",
	}
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

func main() {
	ENV := os.Getenv("ENV")
	if ENV == "" {
		ENV = "dev"
	}
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	viper.SetConfigName(ENV)
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	err := viper.ReadInConfig()
	if err != nil {
		log.Println(err)
	}
	viper.AutomaticEnv()
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("could not decode config into struct: %v", err)
	}
	fmt.Printf("Config struct: %v\n", config)
	// changed := false
	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	err = viper.Unmarshal(&config)
	// 	if err != nil {
	// 		log.Printf("could not decode config after changed: %v", err)
	// 	}
	// 	changed = true
	// })
	// for !changed {
	// 	time.Sleep(time.Second)
	// 	fmt.Printf("Config struct: %v\n", config)
	// }
}
