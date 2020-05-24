package main

import (
	"fmt"
	"sort"
)

type counterMap map[string]uint

type stringCounter struct {
	name  string
	count uint
}

func (cm counterMap) Inc(key string) {
	cm[key]++
}

func (cm counterMap) PrintSorted() {
	for _, entry := range cm.Sorted() {
		fmt.Println(entry.count, entry.name)
	}
}

func (cm counterMap) Sorted() []stringCounter {
	result := make([]stringCounter, 0, len(cm))
	for name, count := range cm {
		result = append(result, stringCounter{name: name, count: count})
	}

	sort.Slice(result, func(i, j int) bool { return result[i].count > result[j].count })
	return result
}
