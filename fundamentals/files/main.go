package main

import (
	"fmt"
	"howardsolutions/go/files/fileutils"
	"os"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"

	content, err := fileutils.ReadTextFile(filePath)

	if err != nil {
		fmt.Printf("ERROR PANIC!! %v", err)
		newContent := fmt.Sprintf("Orignal: %v\n Double Original: %v%v", content, content, content)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Println(content)
	}
}
