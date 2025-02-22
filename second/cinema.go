package second

func findBestSeatDist(row []int) int {
	var counter, maximum int

	for _, val := range row {
		if val == 0 {
			counter++
		} else if counter > 0 {
			if counter > maximum {
				maximum = counter
			}

			counter = 0
		}
	}

	if counter > maximum {
		maximum = counter
	}

	rowBegin := row[:maximum]
	rowEnd := row[len(row)-maximum:]

	if hasExistEmptyRange(rowBegin, maximum) {
		return maximum
	}

	if hasExistEmptyRange(rowEnd, maximum) {
		return maximum
	}

	diff := maximum / 2

	if maximum%2 != 0 {
		diff++
	}

	return diff
}

func hasExistEmptyRange(row []int, maximum int) bool {
	var counter int

	for _, val := range row {
		if val == 0 {
			counter++
		}
	}

	return counter == maximum
}
