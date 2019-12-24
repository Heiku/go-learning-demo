package main

import "fmt"

type Filter interface {
	Detail() string
	Process([]int) []int
}

type UniqueFilter struct {
}

func (UniqueFilter) Detail() string {
	return "delete repeated number"
}

func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pushed := make(map[int]bool)

	for _, n := range inputs {
		if !pushed[n] {
			pushed[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

type MultipleFilter int

func (mf MultipleFilter) Detail() string {
	return fmt.Sprintf("save %v times", mf)
}

func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n%int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

func filterAndPrint(fltr Filter, unFiltered []int) []int {
	filtered := fltr.Process(unFiltered)
	fmt.Println(fltr.Detail()+":\n\t", filtered)
	return filtered
}

func main() {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("before filter: \n\t", numbers)

	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2),
		MultipleFilter(3),
	}

	for _, fltr := range filters {
		numbers = filterAndPrint(fltr, numbers)
	}
}
