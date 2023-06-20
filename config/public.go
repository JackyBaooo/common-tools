package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
)

var (
	GlobalViper  *viper.Viper
	GlobalConfig CommonConfig
)

func Init(file string) (err error) {
	v := viper.New()
	v.AddConfigPath("./")
	v.AddConfigPath("./config")
	if file != "" {
		v.SetConfigFile(file)
	} else {
		v.SetConfigName("config")
	}
	err = v.ReadInConfig()
	if err != nil {
		return
	}
	GlobalViper = v
	err = v.Unmarshal(&GlobalConfig)
	if err != nil {
		return
	}
	out, _ := json.Marshal(GlobalViper.AllSettings())
	log.Printf("GlobalViper: %s", string(out))
	log.Printf("GlobalConfig: %+v", GlobalConfig)
	return
}
