package file

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func ReadFile(dataObj any, path string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		log.Println("error can't get the current directory")
		return err
	}

	jsonData, err := os.Open(fmt.Sprintf("%s%s", currentDir, path))
	if err != nil {
		log.Println(err)
		log.Println(fmt.Sprintf("error can't find %s%s file, please create it on config folder", currentDir, path))
		return err
	}

	dataByte, err := io.ReadAll(jsonData)
	if err != nil {
		log.Println(err)
		log.Println("error on the data io")
		return err
	}

	err = json.Unmarshal(dataByte, &dataObj)
	if err != nil {
		log.Println(err)
		log.Println(fmt.Sprintf("error reading %s%s file, please use a correct json form", currentDir, path))
		return err
	}

	return nil
}
