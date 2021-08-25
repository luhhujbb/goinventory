package main

import (
    api "github.com/luhhujbb/goinventory/api"
    viper "github.com/spf13/viper"
)


func main (){
    viper.SetConfigName("inventory")
    api.Server()
}
