package binarysearch

func SearchInts(list []int, key int) int {
	min, max := 0, len(list)-1

	for min <= max {
		mid := (max + min) / 2
		switch {
		case list[mid] == key:
			return mid
		case list[mid] > key:
			max = mid - 1
		case list[mid] < key:
			min = mid + 1
		}
	}

	return -1
}
