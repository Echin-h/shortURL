package configs

import (
	"ShortUrl/internal/global/log"
	"ShortUrl/tools"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
)

var Path = "configs/config.yml"
var c = Config{}

func Init() {
	viper.SetConfigFile(Path)
	viper.AddConfigPath(".")

	if tools.FileExist(Path) {
		log.SugarLogger.Info("Config Info is ok -- in log ")
		tools.PanicIfErr(
			viper.ReadInConfig(),
			viper.Unmarshal(&c),
		)
	} else {
		fmt.Println("Config file not exist in ", Path, ". Using environment variables.")
		if err := envconfig.Process("", &c); err != nil {
			panic(err)
		}
	}
	fmt.Println("Config read successfully")
}

func Get() Config {
	return c
}
