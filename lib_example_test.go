package envloader_test

import (
	"log"

	envloader "github.com/4strodev/env-loader"
)

func ExampleMarshal() {
	type EnvVariables struct {
		Workers       int    `env:"WORKERS,required"`
		ApiToken      string `env:",required"`
		RemoteLogging bool   // This is optional
	}

	// Expected env variables
	// WORKERS=10
	// ApiToken="some api token"
	var envVariables EnvVariables
	err := envloader.Marshal(&envVariables)
	if err != nil {
		log.Fatal(err)
	}
}
