package main

import "fmt"

func bubble() {
	for i := 0; i < len(sArr); i++ {
		// bubble last num. as loop - i -1
		for j := 0; j < len(sArr)-1-i; j++ {
			if sArr[j] > sArr[j+1] {
				sArr[j], sArr[j+1] = sArr[j+1], sArr[j]
			}
		}
	}
}

func selection() {
	for i := 0; i < len(sArr); i++ {
		min := i
		for j := i + 1; j < len(sArr); j++ {
			if sArr[min] > sArr[j] {
				sArr[min], sArr[j] = sArr[j], sArr[min]
			}
		}
	}
}

func insertion() {
	for i := 1; i < len(sArr); i++ {
		for j := i - 1; j >= 0; j-- {
			if sArr[j] > sArr[j+1] {
				sArr[j+1], sArr[j] = sArr[j], sArr[j+1]
			} else {
				break
			}
		}
	}
}

func quick(left, right int) {
	if left >= right {
		return
	}

	val := sArr[left]
	idx := left

	for i := left + 1; i <= right; i++ {
		if sArr[i] < val {
			sArr[idx] = sArr[i]
			sArr[i] = sArr[idx+1]
			idx++
		}
	}
	sArr[idx] = val

	quick(left, idx-1)
	quick(idx+1, right)
}

func quick2(a []int) []int {
	if len(a) < 2 {
		return a
	}
	basic := a[0]
	result, left, right := make([]int, 0), make([]int, 0), make([]int, 0)

	for i := 1; i < len(a); i++ {
		if a[i] < basic {
			left = append(left, a[i])
		} else {
			right = append(right, a[i])
		}
	}
	result = append(result, quick2(left)...)
	result = append(result, basic)
	result = append(result, quick2(right)...)

	return result
}

var sArr = []int{1, 99, 4, 5, 6, 3, 20}

func main() {

	//bubble(sArr)
	//selection()
	//insertion()
	//quick(0, len(sArr)-1)
	sortedArr := quick2(sArr)
	fmt.Println(sArr)
	fmt.Println(sortedArr)
}
