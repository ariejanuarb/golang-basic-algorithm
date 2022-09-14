/*
selection sort is a sorting algorithm that is used to sort the elements by repeatedly finding the minimum element
and placing it at the first in the unsorted part of the array.

after placing the minimum element in one end of the array,
the array is then divided into two subparts which can be again sorted using the algorithm

Algorithm :
- take an array of integers as the input
- find the index of the minimum element by iterating over the array
- if the number found is minimum, then swap with its previous element
- now return the sorted array
*/

package main
import "fmt"
func SelectionSort(array[] int, size int) []int {
	var min_index int
	var temp int

	for i := 0; i < size - 1; i++ {
		min_index = i
		for j := i + 1; j < size; j++ {
			if array[j] < array[min_index] {
				min_index = j
			}
		}
		temp = array[i]
		array[i] = array[min_index]
		array[min_index] = temp
	}
	return array
}

func main() {
	var num = 7
	array := []int{2,4,3,1,6,8,5}
	fmt.Println(SelectionSort(array, num))
}