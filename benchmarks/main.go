package main

import (
	"ethbench/cmd"
	"fmt"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("error loading .env file")
	}
	fmt.Println("EthBench Started .....")
	cmd.Execute()
}
