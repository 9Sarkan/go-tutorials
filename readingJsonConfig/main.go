package main

import (
	"github.com/spf13/viper"
	"fmt"
)

func main(){
	viper.SetConfigType("json")
	viper.SetConfigFile("./configFile.json")
	viper.ReadInConfig()

	fmt.Printf("Using Config: %s\n", viper.ConfigFileUsed())

	if viper.IsSet("kurdish.key1"){
		fmt.Printf("hello in kurdish is: %s\n", viper.Get("kurdish.key1"))
	}
	if viper.IsSet("spanish.key1"){
		fmt.Printf("hello in spanish is: %s\n", viper.Get("spanish.key1"))
	} else {
		fmt.Printf("Spanish is not configed!\n")
	}
}
