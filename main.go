package main

import (
	"fmt"
	"os"
	"strings"
	"go-reloaded/pkg"
)

func main() {
	filesName := os.Args[1:]
	if len(filesName) != 2 {
		fmt.Println("Error: Argemant number required is not equal to 2")
	} else if filesName[0] != "sample.txt" || filesName[1] != "result.txt" {
		fmt.Println("Error: there is error in the name of files")
	} else {
		readFile, err := os.ReadFile(filesName[0])
		if err != nil {
			fmt.Println(err)
		}
		fileText := string(readFile)
		text := pkg.GoReloaded(fileText)
		if strings.Contains(text, "Error -") {
			err := os.WriteFile(filesName[1], []byte("Error"), 0666)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(text)
		} else {
			err := os.WriteFile(filesName[1], []byte(text), 0666)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
