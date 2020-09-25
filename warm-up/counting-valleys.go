package warmup

func countingValleys(steps int32, path string) int32 {
	var levels int32 = 0
	var valleys int32 = 0
	for i := 0; i < int(steps); i++ {
		if path[i] == 'U' {
			levels = levels + 1
			if levels == 0 {
				valleys = valleys + 1
			}
		} else {
			levels = levels - 1
		}
	}
	return valleys
}
