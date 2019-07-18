package main

import "fmt"

func main() {
	arr := []string{"b", "aaa", "kc", "kk", "abcaaaaa"}
	fmt.Println(MergeSort(arr))
}

func MergeSort(arr []string) []string {
	m := len(arr) / 2
	if len(arr) < 2 {
		return arr
	}

	return Merge(MergeSort(arr[:m]), MergeSort(arr[m:]))
}

func Merge(arr1 []string, arr2 []string) []string {
	size := len(arr1) + len(arr2)
	i, j := 0, 0
	slice := make([]string, size, size)
	for k := 0; k < size; k++ {
		if i > len(arr1)-1 && j <= len(arr2)-1 {
			slice[k] = arr2[j]
			j++
		} else if j > len(arr2)-1 && i <= len(arr1)-1 {
			slice[k] = arr1[i]
		}else if arr1[i] < arr2[j] {
			slice[k] = arr1[i]
			i++
		} else {
			slice[k] = arr2[j]
			j++
		}
	}
	return slice
}
