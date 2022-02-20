package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var (
	lock           = &sync.Mutex{}
	configInstance *Config
)

const ConfigFile = ".env"

type Config struct {
	App
	Database
	UpTimeRobot
}

type App struct {
	Env      string
	Port     int
	LogLevel string
	LogPath  string
}

type Database struct {
	Dialect        string
	Host           string
	Port           string
	Name           string
	User           string
	Password       string
	MaxConnections int
}

type UpTimeRobot struct {
	Url      string
	Interval int
}

// GetConfig Function to get initialized config struct.
// Will initialize it if configInstance is equal to nil.
func GetConfig() *Config {
	if configInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if configInstance == nil {
			c := &Config{}
			err := c.Configure(readConfigs)
			if err != nil {
				panic(err)
			}

			configInstance = c
		}
	}

	return configInstance
}

// Initialize Function to configure program
// configs at the start of the program.
// Have to be used at the main package only.
func Initialize() {
	GetConfig()
}

// Configure This method receives preconfigured
// viper instance from callback argument and parses settings to Config struct
func (c *Config) Configure(viper func() *viper.Viper) error {
	v := viper()

	c.App = App{
		Env:      v.GetString("APP_ENV"),
		LogLevel: v.GetString("LOG_LEVEL"),
		LogPath:  v.GetString("LOG_PATH"),
	}
	c.Database = Database{
		Dialect:        v.GetString("DB_DIALECT"),
		Host:           v.GetString("DB_HOST"),
		User:           v.GetString("DB_USER"),
		Password:       v.GetString("DB_PASSWORD"),
		Name:           v.GetString("DB_NAME"),
		Port:           v.GetString("DB_PORT"),
		MaxConnections: v.GetInt("DB_MAX_CONN"),
	}
	c.UpTimeRobot = UpTimeRobot{
		Url:      v.GetString("UPTIME_URL"),
		Interval: v.GetInt("UPTIME_INTERVAL"),
	}

	return nil
}

func (d *Database) GetDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", d.User, d.Password, d.Host+":"+d.Port, d.Name)
}

// readConfigs is a callback function to read configs from configs.yaml from project root dir
func readConfigs() *viper.Viper {
	viper.SetConfigFile(getConfigFilePath())
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	return viper.GetViper()
}

func getConfigFilePath() string {
	rootFolder, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return rootFolder + "/" + ConfigFile
}
