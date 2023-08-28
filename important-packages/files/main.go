package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	createFile()
	readFile()
	readFileStream()
	removeFile()
}

func createFile() {
	file, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	length, err := file.WriteString("Writing some text\n")
	_, err2 := file.Write([]byte("Writing some text"))

	if err != nil || err2 != nil {
		panic(err)
	}

	fmt.Println("createFile - File created successfully! Length: ", length)

	file.Close()
}

func readFile() {
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("readFile: \n", string(file))
}

func readFileStream() {
	file, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}

	stream := bufio.NewReader(file)
	buffer := make([]byte, 18)

	for {
		line, err := stream.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println("readFileStream:", string(buffer[:line]))
	}
}

func removeFile() {
	err := os.Remove("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("File removed!")
}