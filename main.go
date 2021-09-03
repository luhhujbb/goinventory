package main

import (
    api "github.com/luhhujbb/goinventory/api"
    viper "github.com/spf13/viper"
    inventory "github.com/luhhujbb/goinventory/inventory"
    aws "github.com/luhhujbb/goinventory/provider/aws"
    "log"
    "os"
)

func setDefaultConfig(){
    viper.SetDefault("store.file.bucket","/etc/inventory")
    viper.SetDefault("store.s3.region","eu-west-1")
}

func main (){
    viper.SetConfigName("config")
    viper.AddConfigPath("/etc/inventory/")   // path to look for the config file in
    viper.AddConfigPath("$HOME/.inventory")  // call multiple times to add many search paths
    viper.AddConfigPath(".")
    setDefaultConfig()
    if err := viper.ReadInConfig(); err != nil {
    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
        log.Fatal("configNotFound")
    } else {
        // Config file was found but another error was produced
    }
    }
    if os.Getenv("AWS_REGION") == ""{
        os.Setenv("AWS_REGION", viper.Get("store.s3.region").(string))
    }
    if viper.Get("inventory") != nil {
        log.Print("init inventory")
        inventory.ConfigureInventory(viper.Get("inventory"))
    }
    if viper.Get("aws") != nil {
        log.Print("init aws")
        aws.ConfigureAWS(viper.Get("inventory"))
    }
    api.Server()
}
