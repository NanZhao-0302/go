package fifthHomework

func quickSorter(arr []*Person, start, end int) {
	middle := (start + end) / 2
	midFatRate := arr[middle].fatRate
	s, e := partation(arr, start, end, midFatRate)
	if s == e {
		s++
		e--
	}
	if e > start {
		quickSorter(arr, start, e)
	}
	if s < end {
		quickSorter(arr, s, end)
	}
}

func partation(arr []*Person, left int, right int, midFatRate float64) (int, int) {
	for left <= right {
		for arr[left].fatRate < midFatRate {
			left++
		}
		for arr[right].fatRate > midFatRate {
			right--
		}
		if left >= right {
			break
		}

		arr[left].fatRate, arr[right].fatRate = arr[right].fatRate, arr[left].fatRate
		left++
		right--
	}
	return left, right
}
