package analyzer

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func acceptExtension(file string) bool {
	var config map[string][]interface{}
	content, err := os.ReadFile("config/authorized.json")
	separetedList := strings.Split(file, ".")
	extension := separetedList[len(separetedList) - 1]

	if err != nil {
		log.Fatalln(err)
	}
	
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatalln(err)
	}

	for _, ext := range config["authorized"] {
		if ext == extension {
			return true
		}
	}
	return false
}

func inspectCode(file string) {
	if !acceptExtension(file) {
		return
	}
}

func CodeReport(directory string, recursive bool) {
	files, err := listFiles(directory, recursive)
	
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		inspectCode(file)
	}
}