package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func CreateFile(filename string, content string) {
	var _, err = os.Stat(filename)
	// create file if not exists
	if os.IsNotExist(err) {
		// If the file doesn't exist, create it, or append to the file
		f, err := os.OpenFile(filename, os.O_CREATE, 0644)
		// f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.Write([]byte(content))
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
	fmt.Println("File Created Successfully :", filename)
}

func ReadFileByte(filename string) []byte {
	output, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return output

	// if find != "" && replace != "" {
	// 	output = bytes.Replace(output, []byte(find), []byte(replace), -1)
	// }
	// return string(output)
}

func ReadFileString(filename string) string {
	return string(ReadFileByte(filename))
}
