package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("please enter of your txt file name")

	// address of file
	var address string
	fmt.Scanln(&address)

	//open file
	file, err := os.Open(address)
	if err != nil {
		log.Fatal("error has occured in file opening")
		return
	}
	defer file.Close()

	//read file
	scanner := bufio.NewScanner(file)

	// delete duplicated line
	duplicate := make(map[string]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if duplicate[line] == false {
			duplicate[line] = true
		}
	}

	//create file
	file, err = os.Create("remove_duplicated_file.txt")
	if err != nil {
		log.Fatal("error has been occured in creating file")
		return
	}
	defer file.Close()

	//entering file in new file
	for key := range duplicate {
		_, err := file.WriteString(key + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
