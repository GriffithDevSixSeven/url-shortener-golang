package main

import (
	"fmt"

	"github.com/GriffithDevSixSeven/url-shortener-golang/configs"
)

func main() {

	EnvCfg := configs.EnvConfig{
		EnvFile: ".env",
	}

	fmt.Println(EnvCfg.GetDbUrl())
}
