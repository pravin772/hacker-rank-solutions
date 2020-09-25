package warmup

func sockMerchant(n int32, ar []int32) int32 {
	var pair int32
	for i := 0; i < int(n); i++ {
		c := 1
		for j := i + 1; j < int(n); j++ {
			if ar[i] == ar[j] {
				c++
			}
		}
		if c%2 == 0 {
			pair++
		}
	}
	return pair
}
