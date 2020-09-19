package main

import (
	"io/ioutil"
	s "strings"
	"strconv"
	"fmt"
	"math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func roundDown(number float64) (int64) {
	if math.Mod(number, 1) < 0.5 {
		return int64(math.Round(number))
	}

	return int64(math.Round(number - 1))
}

func fuelRequired(mass int64) (int64) {
	return (roundDown(float64(mass / 3))) - 2
}

func fuelForFuel(fuel int64) (int64) {
	next := fuelRequired(fuel)
	if next <= 0 {
		return 0;
	}

	return next + fuelForFuel(next)
}

func main() {

	raw, err := ioutil.ReadFile("./modules.txt")
	check(err)
	
	chunks := s.Split(string(raw), "\n")

	fuelForModules := make([]int64, len(chunks))

	for i:=0; i < len(chunks); i++ {
		if len(chunks[i]) < 1 {
			fuelForModules[i] = 0
			continue
		}

		n, err := strconv.ParseInt(chunks[i], 10, 64)
		check(err)

		baseModuleFuel := fuelRequired(n)
		fuelForModules[i] = baseModuleFuel + fuelForFuel(baseModuleFuel)
	}

	var grandTotal int64 = 0

	for _, value := range(fuelForModules) {
		grandTotal += value
	}

	fmt.Printf("grandTotal: %d\n", grandTotal)
}