package main

import (
	"ApiGolang/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	println(config.DBDrivee)
}
