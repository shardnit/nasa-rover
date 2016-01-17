package main

import (
	"errors"
	"fmt"
)

type Rover struct {
	initPosition    *position
	currentPosition *position
	boundary        *coordinate
}

type position struct {
	*coordinate
	direction string
}

type coordinate struct {
	x int
	y int
}

func Init(boundary *coordinate, init *coordinate, d string) *Rover {
	r := &Rover{
		initPosition: &position{
			coordinate: init,
			direction:  d,
		},
		boundary: boundary,
	}
	r.currentPosition = r.initPosition
	return r
}

// this method accepts the whole command string, processes it, and set the appropriate position of rover
func (r *Rover) Move(command string) error {
	for _, s := range command {
		if err := r.singleMove(string(s)); err != nil {
			return err
		}
	}
	return nil
}

// this method is used to act on a single command like "L" or "M" or "R"
func (r *Rover) singleMove(command string) error {
	switch command {
	case "L":
		r.turnLeft()
	case "R":
		r.turnRight()
	case "M":
		r.moveForward()
	default:
		return errors.New(fmt.Sprintf("Unknown singular command: %s", command))
	}
	return nil
}

func (r *Rover) turnLeft() {
	switch r.currentPosition.direction {
	case "N":
		r.currentPosition.direction = "W"
	case "S":
		r.currentPosition.direction = "E"
	case "E":
		r.currentPosition.direction = "N"
	case "W":
		r.currentPosition.direction = "S"
	}
}

func (r *Rover) turnRight() {
	switch r.currentPosition.direction {
	case "N":
		r.currentPosition.direction = "E"
	case "S":
		r.currentPosition.direction = "W"
	case "E":
		r.currentPosition.direction = "S"
	case "W":
		r.currentPosition.direction = "N"
	}
}

// we assume that the rover cannot ever cross the plateau boundary
func (r *Rover) moveForward() {
	switch r.currentPosition.direction {
	case "N":
		if r.currentPosition.y < r.boundary.y {
			r.currentPosition.y = r.currentPosition.y + 1
		}
	case "S":
		if r.currentPosition.y > 0 {
			r.currentPosition.y = r.currentPosition.y - 1
		}
	case "E":
		if r.currentPosition.x < r.boundary.x {
			r.currentPosition.x = r.currentPosition.x + 1
		}
	case "W":
		if r.currentPosition.x > 0 {
			r.currentPosition.x = r.currentPosition.x - 1
		}
	}
}

func (r *Rover) String() string {
	return fmt.Sprintf("%v %v %s", r.currentPosition.x, r.currentPosition.y, r.currentPosition.direction)
}
