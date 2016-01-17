package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoverMove(t *testing.T) {
	assert := assert.New(t)
	r := Init(&coordinate{x: 5, y: 5}, &coordinate{x: 1, y: 2}, "N")
	assert.Equal(r.currentPosition, &position{coordinate: &coordinate{x: 1, y: 2}, direction: "N"})
	err := r.Move("LMLMLMLMM")
	assert.Nil(err)
	assert.Equal(r.currentPosition, &position{coordinate: &coordinate{x: 1, y: 3}, direction: "N"})

	r = Init(&coordinate{x: 5, y: 5}, &coordinate{x: 3, y: 3}, "E")
	assert.Equal(r.currentPosition, &position{coordinate: &coordinate{x: 3, y: 3}, direction: "E"})
	err = r.Move("MMRMMRMRRM")
	assert.Nil(err)
	assert.Equal(r.currentPosition, &position{coordinate: &coordinate{x: 5, y: 1}, direction: "E"})
}

func TestXBoundaryMove(t *testing.T) {
	assert := assert.New(t)
	r := Init(&coordinate{x: 5, y: 5}, &coordinate{x: 5, y: 2}, "E")
	err := r.Move("M")
	assert.Nil(err)
	assert.Equal(r.currentPosition, &position{coordinate: &coordinate{x: 5, y: 2}, direction: "E"})

	r = Init(&coordinate{x: 5, y: 5}, &coordinate{x: 0, y: 2}, "W")
	err = r.Move("M")
	assert.Nil(err)
	assert.Equal(r.currentPosition, &position{coordinate: &coordinate{x: 0, y: 2}, direction: "W"})
}

func TestYBoundaryMove(t *testing.T) {
	assert := assert.New(t)
	r := Init(&coordinate{x: 5, y: 5}, &coordinate{x: 2, y: 5}, "N")
	err := r.Move("M")
	assert.Nil(err)
	assert.Equal(r.currentPosition, &position{coordinate: &coordinate{x: 2, y: 5}, direction: "N"})

	r = Init(&coordinate{x: 5, y: 5}, &coordinate{x: 2, y: 0}, "S")
	err = r.Move("M")
	assert.Nil(err)
	assert.Equal(r.currentPosition, &position{coordinate: &coordinate{x: 2, y: 0}, direction: "S"})
}

func TestSingleMoveLeft(t *testing.T) {
	assert := assert.New(t)
	r := Init(&coordinate{x: 5, y: 5}, &coordinate{x: 2, y: 5}, "N")
	err := r.singleMove("L")
	assert.Nil(err)
	assert.Equal("W", r.currentPosition.direction)
	err = r.singleMove("L")
	assert.Nil(err)
	assert.Equal("S", r.currentPosition.direction)
	err = r.singleMove("L")
	assert.Nil(err)
	assert.Equal("E", r.currentPosition.direction)
	err = r.singleMove("L")
	assert.Nil(err)
	assert.Equal("N", r.currentPosition.direction)

	err = r.singleMove("N")
	assert.NotNil(err)
}

func TestSingleMoveRight(t *testing.T) {
	assert := assert.New(t)
	r := Init(&coordinate{x: 5, y: 5}, &coordinate{x: 2, y: 5}, "N")
	err := r.singleMove("R")
	assert.Nil(err)
	assert.Equal("E", r.currentPosition.direction)
	err = r.singleMove("R")
	assert.Nil(err)
	assert.Equal("S", r.currentPosition.direction)
	err = r.singleMove("R")
	assert.Nil(err)
	assert.Equal("W", r.currentPosition.direction)
	err = r.singleMove("R")
	assert.Nil(err)
	assert.Equal("N", r.currentPosition.direction)

	err = r.singleMove("N")
	assert.NotNil(err)
}
