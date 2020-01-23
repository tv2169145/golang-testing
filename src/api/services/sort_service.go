package services

import "github.com/tv2169145/golang-testing/src/api/utils/sort"

func Sort(element []int) {
	if len(element) >= 10000 {
		sort.BasicSortAsc(element)
	} else {
		sort.BubbleSortAsc(element)
	}
}
