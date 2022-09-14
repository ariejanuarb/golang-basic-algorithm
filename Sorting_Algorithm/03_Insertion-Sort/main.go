/*
As you can see, the insertionSort() function starts by iterating over the entire slice in a nested for loop.
The job of the inner for loop is to consume one value for each repetition, and grow the sorted output list, which are all the elements before index i

At each repetition, insertion sort removes one element from the input data,
finds the location it belongs within the sorted first section, and inserts it there. 
It repeats until no input elements remain.
*/


package main
import "fmt"

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func main() {
	fmt.Println(insertionSort([]int{5,3,2,1,0,4}))

}