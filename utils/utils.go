package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func CheckIfError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	return
}

func CheckIfErrorAndLog(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

func FormatToJSON(i interface{}) string {
	jsonString, err := json.MarshalIndent(i, "", "\t")
	CheckIfError(err)
	return string(jsonString)
}
