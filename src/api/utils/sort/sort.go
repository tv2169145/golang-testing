package sort

import "sort"

func BubbleSortAsc(element []int) {
	running := true
	for running {
		running = false
		for i := 0; i < len(element)-1; i++ {
			if element[i] > element[i+1] {
				element[i], element[i+1] = element[i+1], element[i]
				running = true
			}
		}
	}
}

func BasicSortAsc(element []int) {
	sort.Ints(element)
}

//--------------

func BubbleSortDesc(element []int) {
	running := true
	for running {
		running = false
		for i := 0; i < len(element)-1; i++ {
			if element[i] < element[i+1] {
				element[i], element[i+1] = element[i+1], element[i]
				running = true
			}
		}
	}
}


func BasicSortDesc(element []int) {
	sort.Ints(element)
	sort.Sort(sort.Reverse(sort.IntSlice(element)))
}
