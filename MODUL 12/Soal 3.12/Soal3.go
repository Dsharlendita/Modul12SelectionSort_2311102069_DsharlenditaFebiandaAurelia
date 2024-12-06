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
	data := []int{}

	for {
		var x int
		fmt.Scan(&x)

		if x == -5313 {
			break
		} else if x == 0 {
			selectionSort(data)
			n := len(data)
			if n == 0 {
				continue
			}
			var median int
			if n%2 == 1 {
				median = data[n/2]
			} else {
				median = (data[n/2-1] + data[n/2]) / 2
			}
			fmt.Println(median)
		} else {
			data = append(data, x)
		}
	}
}
