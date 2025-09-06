package main 

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("/temp/defer.txt")
	defer closeFile(file)
	writeFile(file)
}


func createFile(filePath string) *os.File {
	fmt.Println("Creating")
	file, err := os.Create(filePath)
	if err!=nil {
		panic(err)
	}
	return file
}


func writeFile(file *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(file,"done")
}

func closeFile(file *os.File) {
	fmt.Println("closing")
	err := file.Close()

	if err!=nil {
		panic(err)
	}
}
