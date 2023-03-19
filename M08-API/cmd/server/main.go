package main

import "github.com/amauryeuzebio/goexpert/M08-API/configs"

func main() {
	config, _ := configs.LoadConfig(".")

	println(config.DBDriver)
}
