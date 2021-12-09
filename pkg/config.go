package pkg

import "github.com/spf13/viper"

type Config struct {
    AppName string
    Url string
    ProjectName string
    Environment string
    Authorization string
    UserId string
    FlagName string
    Properties map[string]string
    InstanceId string
}

var AppConfig Config

func InitConfig(commandName string) {
    if err := viper.UnmarshalKey(commandName, &AppConfig); err != nil {
        panic(err)
    }
}
