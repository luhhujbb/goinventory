package main

import (
    api "github.com/luhhujbb/goinventory/api"
    viper "github.com/spf13/viper"
)

func setDefaultConfig(){
    viper.SetDefault("store.file.bucket","/etc/inventory")
}

func main (){
    viper.SetConfigName("inventory")
    api.Server()
}
