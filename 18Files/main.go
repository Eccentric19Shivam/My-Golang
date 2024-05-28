package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to file in golang")
	content := "This needs to go un a file - LeanCodeOnline.in"

	file, err := os.Create("./mylcofile.txt")

	checkNilError(err) // shut down the execution of program and show error

	length, err := io.WriteString(file, content)

	checkNilError(err)
	fmt.Println("Length is: ", length)
	defer file.Close() // as we might want to write further code down, so to close at the end we write defer so we write it once ot is remenbered by defer statement
	readFile("./mylcofile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)

	checkNilError(err)

	fmt.Println("text data inside the file is \n", string(databyte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
