package binarysearch

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	var mid int

	for low <= high {
		mid = (low + high) / 2
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// [ 1, 2, 3, 4, 5, 6]

// target = 1

// 6
// mid := len(arr)/2
// mid = 3  // idx 3
// val = 4

// divide the arr again
// berapa nilai array nyaa sekarng

// mid := len(arr)/2
// kita cek apakah nilai ini sama dengan target kalo sama
