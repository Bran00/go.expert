package main

import "github.com/Bran00/go.expert/39/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}