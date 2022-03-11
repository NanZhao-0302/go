package fifthHomework

func bubbleSort(arr []*Person) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j].fatRate > arr[j+1].fatRate {
				arr[j].fatRate, arr[j+1].fatRate = arr[j+1].fatRate, arr[j].fatRate
			}
		}
	}
}
