package second

func getMax(first, second int) int {
	if first > second {
		return first
	}

	return second
}

func getHalf(maximum int) int {
	if maximum > 0 {
		diff := maximum / 2

		if maximum%2 != 0 {
			diff++
		}

		return diff
	}

	return 0
}
