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

func parseRacesP1(data string) (races []*Race) {
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

func parseRacesP2(data string) (races *Race) {
	fields := strings.Fields(data)
	length := len(fields)

	time := ""
	distance := ""

	for i := 1; i < length/2; i++ {
		time = time + fields[i]
		distance = distance + fields[i+length/2]
	}

	race := &Race{}
	var err error

	race.Time, err = strconv.Atoi(time)
	if err != nil {
		panic(err)
	}

	race.Distance, err = strconv.Atoi(distance)
	if err != nil {
		panic(err)
	}

	return race
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
	races := parseRacesP1(data)
	productOfStrategies := 1

	for _, race := range races {
		productOfStrategies = productOfStrategies * calculateWinningStrategies(race)
	}

	println("Part 1:", productOfStrategies)

	race := parseRacesP2(data)
	println("Part 2:", calculateWinningStrategies(race))
}
