package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Mapper struct {
	Name      string
	SrcRanges [][]int
	DstRanges [][]int
}

func (m *Mapper) MapNumber(number int) int {
	for i := 0; i < len(m.SrcRanges); i++ {
		if number >= m.SrcRanges[i][0] && number <= m.SrcRanges[i][1] {
			return m.DstRanges[i][0] + (number - m.SrcRanges[i][0])
		}
	}

	return number
}

// MapRange maps a range of numbers to a new range of numbers.
// ===========================================================
//
// Seed: (1,3)
//        src  =>  dst
// Map1: (1,4) => (2,5)
//
// Seed (1,3) is in range of Map1 => mapped to list of ranges [(2,4)]
// -----------------------------------------------------------
// Seed: (3,5)
//        src  =>  dst
// Map1: (1,4) => (2,5)
// Map2: (5,8) => (0,3)
//
// Seed (3,5) is in range of Map1 and Map2 => mapped to list of ranges [(4,5) (0)]

func (m *Mapper) MapRange(start, end int) (mapped [][]int) {
	for i := 0; i < len(m.SrcRanges); i++ {
		if start >= m.SrcRanges[i][0] && end <= m.SrcRanges[i][1] {
			mapped = append(mapped, []int{m.DstRanges[i][0] + (start - m.SrcRanges[i][0]), m.DstRanges[i][0] + (end - m.SrcRanges[i][0])})
		} else {
			mapped = append(mapped, []int{start, end})
		}
	}

	return mapped
}

func createMapper(name string, mappingRules [][]int) Mapper {
	m := Mapper{
		Name:      name,
		SrcRanges: [][]int{},
		DstRanges: [][]int{},
	}

	for _, mr := range mappingRules {
		m.SrcRanges = append(m.SrcRanges, []int{mr[1], mr[1] + mr[2] - 1})
		m.DstRanges = append(m.DstRanges, []int{mr[0], mr[0] + mr[2] - 1})
	}

	return m
}

func parseMappingRules(input string) Mapper {
	lines := strings.Split(input, "\n")

	name := strings.Trim(lines[0], ":")
	mappingRules := [][]int{}

	for _, line := range lines[1:] {
		mappingRule := []int{}
		for _, field := range strings.Fields(line) {
			number, err := strconv.Atoi(field)
			if err != nil {
				panic(err)
			}
			mappingRule = append(mappingRule, number)
		}
		mappingRules = append(mappingRules, mappingRule)
	}

	return createMapper(name, mappingRules)
}

func getSeeds(input string) (seeds [][]int) {
	fields := strings.Fields(strings.Split(input, ":")[1])

	var seedRanges []int
	for _, field := range fields {
		seed, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		seedRanges = append(seedRanges, seed)
	}

	for i := 0; i < len(seedRanges); i = i + 2 {
		start, end := seedRanges[i], seedRanges[i]+seedRanges[i+1]-1
		seeds = append(seeds, []int{start, end})
	}

	return seeds
}

//go:embed data.txt
var data string

func main() {
	mappingRules := strings.Split(data, "\n\n")
	seedRanges := getSeeds(mappingRules[0])

	var listOfMappers []Mapper
	for i := 1; i < len(mappingRules); i++ {
		mapper := parseMappingRules(mappingRules[i])
		listOfMappers = append(listOfMappers, mapper)
	}

	var locationNumber int
	for i := 0; i < len(seedRanges); i++ {
		for seed := seedRanges[i][0]; seed <= seedRanges[i][1]; seed++ {
			mappedSeed := seed
			for _, mapper := range listOfMappers {
				mappedSeed = mapper.MapNumber(mappedSeed)
			}
			if mappedSeed < locationNumber || locationNumber == 0 {
				locationNumber = mappedSeed
			}
		}
	}

	fmt.Println("Lowest location number:", locationNumber)
}
