package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"webservice/config"
)

// start args
var (
	// envFilepath .
	envFilepath = flag.String("env_filepath", "", "env filepath")
)

// Run preparing for the start of the service
func Run(ctx context.Context) {
	// parse the args
	if !flag.Parsed() {
		flag.Parse()
	}

	fmt.Println()
	fmt.Printf("start args: \n - %+v\n\n", os.Args)
	fmt.Println()
	fmt.Println("service will load default env file first, then read os env variables")
	fmt.Println()

	err := config.Parse(*envFilepath)
	if err != nil {
		panic("read env file failed, err: " + err.Error())
	}
}
