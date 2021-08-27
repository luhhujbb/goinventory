package main

import (
    api "github.com/luhhujbb/goinventory/api"
    viper "github.com/spf13/viper"
    inventory "github.com/luhhujbb/goinventory/inventory"
)

func setDefaultConfig(){
    viper.SetDefault("store.file.bucket","/etc/inventory")
}

func main (){
    viper.SetConfigName("config")
    viper.AddConfigPath("/etc/inventory/")   // path to look for the config file in
    viper.AddConfigPath("$HOME/.inventory")  // call multiple times to add many search paths
    viper.AddConfigPath(".")
    inventory.ConfigureInventory(viper.Get("inventory"))
    api.Server()
}
