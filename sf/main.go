package main

import "fmt"

func main() {
	a := []int{1, 4, 12, 6, 3, -63, 0}
	selectionSort(a)
}

func selectionSort(a []int) {
	l := len(a)
	if a == nil || l < 2 {
		return
	}
	for i := 0; i < l-1; i++ {
		minIndex := i
		for j := i + 1; j < l; j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			a[i] ^= a[minIndex]
			a[minIndex] ^= a[i]
			a[i] ^= a[minIndex]
		}
		fmt.Printf("%d ", a[i])
	}
	fmt.Printf("%d", a[l-1])

}
