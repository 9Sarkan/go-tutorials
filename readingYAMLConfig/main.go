package main

import (
	"fmt"
	"github.com/spf13/viper"
	flag "github.com/spf13/pflag"
	"os"
)

func main(){
	var configFile *string = flag.String("configFile", "myConfig", "set yaml config file!")
	flag.Parse()

	_, err := os.Stat(*configFile)
	if err == nil {
		fmt.Printf("Using User Specified Configuration file!\n")
		viper.SetConfigFile(*configFile)
	} else {
		viper.SetConfigName("myConfig")
		viper.AddConfigPath("/tmp")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
	}

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("can not read config!\n")
		return
	} else {
		fmt.Printf("Using Config file: %s!\n", viper.ConfigFileUsed())
	}

	if viper.IsSet("item1.k1"){
		fmt.Printf("itme 1, k2: %s\n", viper.Get("item1.k2"))
	} else {
		fmt.Printf("item not found!\n")
	}

}
