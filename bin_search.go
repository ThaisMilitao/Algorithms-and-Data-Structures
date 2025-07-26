package main

//binary search on a sorted ascending array
func bin_search(val int, list []int, start int, end int) int {

	if start > end {
		return -1
	}
	mid := (end + start) / 2

	if list[mid] == val {
		return mid

	} else if list[mid] > val {
		end = (mid - 1)

	} else {
		start = (mid + 1)
	}
	return bin_search(val, list, start, end)
}

// binary search on a sorted descending array
func rev_bin_search(val int, list []int, start int, end int) int {
	if start > end {
		return -1
	}
	mid := (start + end) / 2

	if list[mid] == val {
		return mid
	} else if list[mid] > val {
		start = mid + 1
	} else {
		end = mid - 1
	}

	return rev_bin_search(val, list, start, end)
}

func main() {

}
