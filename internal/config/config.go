package config

import (
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type appConfig struct {
	App struct {
		Name    string
		Version string
	}
	Server struct {
		Port        int16
		Timeout     int
		Environment string
	}
	Database struct {
		Mariadb struct {
			Host            string
			Port            int
			User            string
			Password        string
			Dbname          string
			MaxIdleConns    int
			MaxOpenConns    int
			ConnMaxIdleTime int
			ConnMaxLifetime int
		}
	}
}

var Viper = new(appConfig)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(filename)
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(filepath.Join(basepath, "../../"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Viper)
	if err != nil {
		panic(err)
	}
}
