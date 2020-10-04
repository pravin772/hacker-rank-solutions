package arrays

func minimumSwaps(arr []int32) int {
	var c int = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != int32(i+1) {
			for j := i + 1; j < len(arr); j++ {
				if arr[j] == int32(i+1) {
					arr[i], arr[j] = arr[j], arr[i]
					c++
					break
				}
			}
		}
	}
	return c
}
