package main

import (
	engine "ecomm/pkg/server"
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	serv := engine.New()
	fmt.Println("Server is running on port 5000")
	serv.Engine.Run(":5000")
}