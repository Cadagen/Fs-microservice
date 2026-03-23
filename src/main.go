package main

import (
	"cadagen/microservices/fs/src/commands"
	"cadagen/microservices/fs/src/utils"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	operation := os.Args[1]

	switch operation {
	case "ls":
		returnOk(commands.Ls(os.Args[2]))
	case "read":
		readCmd := flag.NewFlagSet("read", flag.ExitOnError)

		start := readCmd.Int("start", -1, "Start line")
		end := readCmd.Int("end", -1, "End line")

		readCmd.Parse(os.Args[2:])

		returnOk(commands.Read(readCmd.Arg(0), *start, *end))
	case "glob":
		returnOk(commands.Glob(os.Args[2], os.Args[3], strings.Split(os.Args[4], ",")))
	case "grep":
		returnOk(commands.Grep(os.Args[2], os.Args[4], os.Args[3], strings.Split(os.Args[5], ",")))
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
