package utils

import (
	//"encoding/json"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func Throw(thrownError error) {
	json, err := json.Marshal(NewCliError(thrownError.Error()))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
	os.Exit(0)
}
