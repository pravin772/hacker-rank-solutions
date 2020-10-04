package arrays

func hourglassSum(arr [][]int32) int32 {
	var sum int32 = -100
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			var top int32 = arr[i][j] + arr[i][j+1] + arr[i][j+2]
			var middle int32 = arr[i+1][j+1]
			var bottom int32 = arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]

			if top+middle+bottom > sum {
				sum = top + middle + bottom
			}
		}
	}
	return sum
}
