package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

func main() {
	if len(os.Args) == 1 {
		glog.Errorf("Please provide input filename as first argument")
		os.Exit(1)
	}
	inputFile := os.Args[1]
	if file, err := os.Open(inputFile); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)

		scanner.Scan()
		boundaryString := scanner.Text()
		boundaryArray := strings.Split(boundaryString, " ")

		boundaryX, err := strconv.Atoi(boundaryArray[0])
		if err != nil {
			glog.Fatalf("Error[%s] while converting boundary X coordinate[%s] to int", err, boundaryArray[0])
			os.Exit(1)
		}

		boundaryY, err := strconv.Atoi(boundaryArray[1])
		if err != nil {
			glog.Fatalf("Error[%s] while converting boundary Y coordinate[%s] to int", err, boundaryArray[1])
			os.Exit(1)
		}

		boundary := &coordinate{x: boundaryX, y: boundaryY}
		readPosition := true
		readCommand := false
		rover := &Rover{}

		for scanner.Scan() {
			if readPosition {
				// first line is rover's initial position
				initArray := strings.Split(scanner.Text(), " ")
				initX, _ := strconv.Atoi(initArray[0])
				initY, _ := strconv.Atoi(initArray[1])
				rover = Init(boundary, &coordinate{x: initX, y: initY}, initArray[2])
				readCommand = true
				readPosition = false
			} else if readCommand {
				// then read the command
				command := scanner.Text()
				if err := rover.Move(command); err != nil {
					glog.Errorf("Error[%s] while moving rover", err)
				} else {
					fmt.Println(rover)
				}
				readCommand = false
				readPosition = true
			}
		}
	} else {
		glog.Fatal(err)
	}

}
