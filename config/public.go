package config

import (
	"encoding/json"
	"github.com/JackyBaooo/common-tools/internal"
	"github.com/spf13/viper"
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
	internal.Infof("GlobalViper: %s", string(out))
	internal.Infof("GlobalConfig: %+v", GlobalConfig)
	return
}
