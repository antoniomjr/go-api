package main

import "github.com/antoniomjr/go/9-apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}