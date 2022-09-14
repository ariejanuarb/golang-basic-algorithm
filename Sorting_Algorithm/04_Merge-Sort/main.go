package main
import "fmt"

// split the array/slice into two roughly equal parts
func mergeSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[0:len(items)/2])
	second := mergeSort(items[len(items)/2])
	return merge(first, second)
}

// merging two sorted list back to a single sorted list
func merge(a []int, b []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
}

