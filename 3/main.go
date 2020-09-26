package main


import (
	"fmt"
	"strconv"
	"strings"
	"math"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Point struct {
	x int
	y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(a Point, b Point) int {
	return Abs(a.x - b.x) + Abs(a.y-b.y)
}
func Task3() string {
	originX := 0
	originY := 0
	
	minDistance := math.MaxInt64
	stepsDone := math.MaxInt64
	
	mapdata := make(map[Point] int)

	raw, err := ioutil.ReadFile("./input.txt")
	check(err)

	lines := strings.Split(string(raw), "\n")

	for line, text := range(lines) {
		fmt.Println("Line: ", line);
		if(len(text) != 0) {
			steps := 0
			curX := originX
			curY := originY
			for _, data := range strings.Split(text, ",") {
				if data != "" {
					dirX := 0
					dirY := 0
					val := toInt(data[1:])
					
					if(string(data[0]) == "L") {
						dirX = -1;
						dirY = 0;
					} else if(string(data[0]) == "R") {
						dirX = 1;
						dirY = 0;
					} else if(string(data[0]) == "U") {
						dirX = 0;
						dirY = +1;
					} else if(string(data[0]) == "D") {
						dirX = 0;
						dirY = -1;
					}
					for ; val > 0; val-- {
						curX += dirX
						curY += dirY
						steps++

						if(line == 1 && mapdata[Point{curX, curY}] != 0) {
							// Intersection
							dist := manhattanDistance(Point{originX,originY}, Point{curX, curY})

							if dist != 0 && (steps+mapdata[Point{curX, curY}]) < stepsDone {
								stepsDone = (steps+mapdata[Point{curX, curY}])
							}
							if dist != 0 && dist < minDistance {
								minDistance = dist
							}
						}
						
						if (line == 0) {
							mapdata[Point{curX, curY}] = steps
						}
					}
					
				}
			}
		} else {
			break
		}
		
	}
	
	return strconv.Itoa(stepsDone) + " | Distance : " + strconv.Itoa(minDistance)
}

func toInt(v string) int {
	i, err := strconv.Atoi(v)
	if err != nil { panic(err) }
	return i
}

func main() {
	fmt.Println(Task3())
}