package main

import (
	//"bufio"
	//"os"
	"fmt"
	"io/ioutil"
	//"log"
	//"strings"
	//"strconv"
)

//read in file (mass of fuel)
func input() {

	data, err := ioutil.ReadFile("input.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:\n", string(data))
}

func calcFuel() {

	
}



func main() {

	input() //read in file (mass of fuel)

}