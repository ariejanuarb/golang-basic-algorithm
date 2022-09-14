/*
Bubble Sort is a sorting algorithm that works by swapping the elements that are in the wrong order. 
In multiple passes, it checks if the adjacent elements are in the right order (increasing) or not

The time complexity of the bubble sort is O(n^2) since it takes two nested loops to check the adjacent element

Algorithm :
- In two nested loops, compare each element with its adjacent element.
- Swap the element if it is less than the previous element.
- Print the array.
*/

package main
import "fmt"

func BubbleSort(array[] int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if (array[j] > array [j+1]) {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	array := []int{11, 14, 3, 8, 18, 17, 43};
	fmt.Println(BubbleSort(array))
}