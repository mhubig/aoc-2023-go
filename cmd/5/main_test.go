package main

import (
	"testing"
)

func TestMapper_GivenName_ReturnsMapperWithGivenName(t *testing.T) {
	mappingRules := [][]int{}
	name := "foo-to-bar map"
	mapper := createMapper(name, mappingRules)

	if mapper.Name != name {
		t.Errorf("Expected %s, got %s", name, mapper.Name)
	}
}

func TestMapper_GivenOneMappingRule_ReturnsMapperWithOneSrcAndOneDestRange(t *testing.T) {
	name := "foo-to-bar map"
	mappingRules := [][]int{{1, 2, 3}}
	expected := Mapper{
		Name:      name,
		SrcRanges: [][]int{{2, 4}},
		DstRanges: [][]int{{1, 3}},
	}
	mapper := createMapper(name, mappingRules)

	if len(mapper.SrcRanges) != len(expected.SrcRanges) {
		t.Errorf("Expected 1, got %d", len(mapper.SrcRanges))
	}

	if mapper.SrcRanges[0][0] != expected.SrcRanges[0][0] {
		t.Errorf("Expected %d, got %d", expected.SrcRanges[0][0], mapper.SrcRanges[0][0])
	}

	if mapper.SrcRanges[0][1] != expected.SrcRanges[0][1] {
		t.Errorf("Expected %d, got %d", expected.SrcRanges[0][1], mapper.SrcRanges[0][1])
	}

	if len(mapper.DstRanges) != len(expected.DstRanges) {
		t.Errorf("Expected 1, got %d", len(mapper.DstRanges))
	}

	if mapper.DstRanges[0][0] != expected.DstRanges[0][0] {
		t.Errorf("Expected %d, got %d", expected.DstRanges[0][0], mapper.DstRanges[0][0])
	}

	if mapper.DstRanges[0][1] != expected.DstRanges[0][1] {
		t.Errorf("Expected %d, got %d", expected.DstRanges[0][1], mapper.DstRanges[0][1])
	}
}

func TestMapNumber_GivenNumberFromMappingRange_ReturnsMapedNumber(t *testing.T) {
	mappingRules := [][]int{{1, 2, 3}}
	mapper := createMapper("foo-to-bar map", mappingRules)

	if mapper.MapNumber(2) != 1 {
		t.Errorf("Expected 1, got %d", mapper.MapNumber(2))
	}

	if mapper.MapNumber(3) != 2 {
		t.Errorf("Expected 2, got %d", mapper.MapNumber(3))
	}

	if mapper.MapNumber(4) != 3 {
		t.Errorf("Expected 3, got %d", mapper.MapNumber(4))
	}
}

func TestMapNumber_GivenNumberFromOutsideOfMappingRange_ReturnsSameNumber(t *testing.T) {
	mappingRules := [][]int{{1, 2, 3}}
	mapper := createMapper("foo-to-bar map", mappingRules)

	if mapper.MapNumber(1) != 1 {
		t.Errorf("Expected 1, got %d", mapper.MapNumber(2))
	}

	if mapper.MapNumber(5) != 5 {
		t.Errorf("Expected 5, got %d", mapper.MapNumber(3))
	}
}

func TestParseMappingRules_GivenMappingName_ReturnsMapperWithName(t *testing.T) {
	input := "foo-to-bar map:"
	name := "foo-to-bar map"
	mapper := parseMappingRules(input)

	if mapper.Name != name {
		t.Errorf("Expected `%s`, got %s", name, mapper.Name)
	}
}

func TestParseMappingRules_GivenOneMappingRule_ReturnsMapperWithOneSrcAndOneDestRange(t *testing.T) {
	input := "foo-to-bar map:\n1 2 3"
	mapper := parseMappingRules(input)
	expected := Mapper{
		Name:      "foo-to-bar map",
		SrcRanges: [][]int{{2, 4}},
		DstRanges: [][]int{{1, 3}},
	}

	if len(mapper.SrcRanges) != len(expected.SrcRanges) {
		t.Errorf("Expected %d, got %d", len(expected.SrcRanges), len(mapper.SrcRanges))
	}

	if mapper.SrcRanges[0][0] != expected.SrcRanges[0][0] {
		t.Errorf("Expected %d, got %d", expected.SrcRanges[0][0], mapper.SrcRanges[0][0])
	}

	if mapper.SrcRanges[0][1] != expected.SrcRanges[0][1] {
		t.Errorf("Expected %d, got %d", expected.SrcRanges[0][1], mapper.SrcRanges[0][1])
	}

	if len(mapper.DstRanges) != len(expected.DstRanges) {
		t.Errorf("Expected %d, got %d", len(expected.DstRanges), len(mapper.DstRanges))
	}

	if mapper.DstRanges[0][0] != expected.DstRanges[0][0] {
		t.Errorf("Expected %d, got %d", expected.DstRanges[0][0], mapper.DstRanges[0][0])
	}

	if mapper.DstRanges[0][1] != expected.DstRanges[0][1] {
		t.Errorf("Expected %d, got %d", expected.DstRanges[0][1], mapper.DstRanges[0][1])
	}
}

func TestParseMappingRules_GivenTwoMappingRules_ReturnsMapperWithMappingsFromBothRules(t *testing.T) {
	input := "seed-to-soil map:\n50 98 2\n52 50 48"
	mapper := parseMappingRules(input)
	expected := Mapper{
		Name:      "seed-to-soil map",
		SrcRanges: [][]int{{98, 99}, {50, 97}},
		DstRanges: [][]int{{50, 51}, {52, 99}},
	}

	if len(mapper.SrcRanges) != len(expected.SrcRanges) {
		t.Errorf("Expected %d, got %d", len(expected.SrcRanges), len(mapper.SrcRanges))
	}

	if mapper.SrcRanges[0][0] != expected.SrcRanges[0][0] {
		t.Errorf("Expected %d, got %d", expected.SrcRanges[0][0], mapper.SrcRanges[0][0])
	}

	if mapper.SrcRanges[0][1] != expected.SrcRanges[0][1] {
		t.Errorf("Expected %d, got %d", expected.SrcRanges[0][1], mapper.SrcRanges[0][1])
	}

	if mapper.SrcRanges[1][0] != expected.SrcRanges[1][0] {
		t.Errorf("Expected %d, got %d", expected.SrcRanges[1][0], mapper.SrcRanges[1][0])
	}

	if mapper.SrcRanges[1][1] != expected.SrcRanges[1][1] {
		t.Errorf("Expected %d, got %d", expected.SrcRanges[1][1], mapper.SrcRanges[1][1])
	}

	if len(mapper.DstRanges) != len(expected.DstRanges) {
		t.Errorf("Expected %d, got %d", len(expected.DstRanges), len(mapper.DstRanges))
	}

	if mapper.DstRanges[0][0] != expected.DstRanges[0][0] {
		t.Errorf("Expected %d, got %d", expected.DstRanges[0][0], mapper.DstRanges[0][0])
	}

	if mapper.DstRanges[0][1] != expected.DstRanges[0][1] {
		t.Errorf("Expected %d, got %d", expected.DstRanges[0][1], mapper.DstRanges[0][1])
	}

	if mapper.DstRanges[1][0] != expected.DstRanges[1][0] {
		t.Errorf("Expected %d, got %d", expected.DstRanges[1][0], mapper.DstRanges[1][0])
	}

	if mapper.DstRanges[1][1] != expected.DstRanges[1][1] {
		t.Errorf("Expected %d, got %d", expected.DstRanges[1][1], mapper.DstRanges[1][1])
	}
}

func TestGetSeeds_GivenSeeds_ReturnsSeeds(t *testing.T) {
	input := "seeds: 1 2 3 4"
	seeds := getSeeds(input)

	if len(seeds) != 2 {
		t.Errorf("Expected 2, got %d", len(seeds))
	}

	if seeds[0][0] != 1 {
		t.Errorf("Expected 1, got %d", seeds[0][0])
	}

	if seeds[0][1] != 2 {
		t.Errorf("Expected 2, got %d", seeds[0][1])
	}

	if seeds[1][0] != 3 {
		t.Errorf("Expected 3, got %d", seeds[1][0])
	}

	if seeds[1][1] != 6 {
		t.Errorf("Expected 6, got %d", seeds[1][1])
	}
}

func TestMapRange_GivenRangeIsSubsetOfMap_ReturnsMappedRange(t *testing.T) {
	mappingRules := [][]int{{2, 1, 4}}
	mapper := createMapper("foo-to-bar map", mappingRules)
	mapped := mapper.MapRange(1, 3)

	if len(mapped) != 1 {
		t.Errorf("Expected 1, got %d", len(mapped))
	}

	if mapped[0][0] != 2 {
		t.Errorf("Expected 2, got %d", mapped[0][0])
	}

	if mapped[0][1] != 4 {
		t.Errorf("Expected 4, got %d", mapped[0][1])
	}
}

func TestMapRange_GivenRangeIsDisjunctToMap_ReturnsGivenRange(t *testing.T) {
	mappingRules := [][]int{{2, 1, 4}}
	mapper := createMapper("foo-to-bar map", mappingRules)
	mapped := mapper.MapRange(5, 8)

	if len(mapped) != 1 {
		t.Errorf("Expected 1, got %d", len(mapped))
	}

	if mapped[0][0] != 5 {
		t.Errorf("Expected 5, got %d", mapped[0][0])
	}

	if mapped[0][1] != 8 {
		t.Errorf("Expected 8, got %d", mapped[0][1])
	}
}
