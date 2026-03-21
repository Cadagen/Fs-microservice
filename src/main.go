package main

import (
	"cadagen/microservices/fs/src/commands"
	"cadagen/microservices/fs/src/utils"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func main() {
	operation := os.Args[1]

	switch operation {
	case "ls":
		returnOk(commands.Ls(os.Args[2]))
	default:
		utils.Throw(errors.New("Invalid command"))
	}
}

func returnOk[T any](data utils.CliResponse[T]) {
	json, err := json.Marshal(data)
	if err != nil {
		utils.Throw(err)
	}

	fmt.Println(string(json))
	os.Exit(0)
}
