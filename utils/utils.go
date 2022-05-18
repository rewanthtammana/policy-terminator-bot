// Helper function file
package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

// Checks for error & kills the app
func CheckIfError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	return
}

// Checks for error & logs the error
func CheckIfErrorAndLog(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

// Just another debugging function
func FormatToJSON(i interface{}) string {
	jsonString, err := json.MarshalIndent(i, "", "\t")
	CheckIfError(err)
	return string(jsonString)
}
