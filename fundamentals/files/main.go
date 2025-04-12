package main

import (
	"fmt"
	"howardsolutions/go/files/fileutils"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()

	content, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")

	if err != nil {
		fmt.Printf("ERROR PANIC!! %v", err)
	} else {
		fmt.Println(content)
	}
}
