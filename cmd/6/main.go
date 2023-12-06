package main

import (
	_ "embed"
	"strconv"
	"strings"
)

type Race struct {
	Time     int
	Distance int
}

func parseRaces(data string) (races []*Race) {
	numbers := strings.Fields(data)
	length := len(numbers)

	for i := 1; i < length/2; i++ {
		time, err := strconv.Atoi(numbers[i])
		if err != nil {
			panic(err)
		}

		distance, err := strconv.Atoi(numbers[i+length/2])
		if err != nil {
			panic(err)
		}

		race := &Race{
			Time:     time,
			Distance: distance,
		}
		races = append(races, race)
	}

	return races
}

func calculateWinningStrategies(r *Race) (strategies int) {
	for i := 1; i <= r.Time; i++ {
		if (r.Time-i)*i > r.Distance {
			strategies++
		}
	}

	return strategies
}

//go:embed data.txt
var data string

func main() {
	races := parseRaces(data)
	productOfStrategies := 1

	for _, race := range races {
		productOfStrategies = productOfStrategies * calculateWinningStrategies(race)
	}

	println(productOfStrategies)
}
