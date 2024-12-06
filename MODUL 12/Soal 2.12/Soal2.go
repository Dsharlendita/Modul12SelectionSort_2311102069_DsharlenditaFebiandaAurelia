package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		rumah := make([]int, m)
		ganjil := []int{}
		genap := []int{}

		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
			if rumah[j]%2 == 0 {
				genap = append(genap, rumah[j])
			} else {
				ganjil = append(ganjil, rumah[j])
			}
		}

		selectionSort(ganjil)
		selectionSort(genap)

		for _, v := range ganjil {
			fmt.Print(v, " ")
		}
		for j := len(genap) - 1; j >= 0; j-- {
			fmt.Print(genap[j], " ")
		}
		fmt.Println()
	}
}
