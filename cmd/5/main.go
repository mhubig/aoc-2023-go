package main

import (
	_ "embed"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"slices"
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

func getSeeds(input string) (seeds []int) {
	fields := strings.Fields(strings.Split(input, ":")[1])
	for _, field := range fields {
		seed, err := strconv.Atoi(field)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, seed)
	}
	return seeds
}

//go:embed data.txt
var data string

func main() {

	go func() {
		http.ListenAndServe(":8080", nil)
	}()

	mappingRules := strings.Split(data, "\n\n")
	seeds := getSeeds(mappingRules[0])

	var listOfMappers []Mapper
	for i := 1; i < len(mappingRules); i++ {
		mapper := parseMappingRules(mappingRules[i])
		listOfMappers = append(listOfMappers, mapper)
	}

	locationNumbers := []int{}
	for _, seed := range seeds {
		for _, mapper := range listOfMappers {
			seed = mapper.MapNumber(seed)
		}
		locationNumbers = append(locationNumbers, seed)
	}

	slices.Sort(locationNumbers)
	fmt.Println("lowest location number:", locationNumbers[0])
}
