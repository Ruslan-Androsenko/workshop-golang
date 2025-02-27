package second

func findBestSeatDist(row []int) int {
	var (
		counter, maxBegin, maxMiddle int
		hasBegin, hasMiddle, hasEnd  bool
		lastIndex                    = len(row) - 1
	)

	for index, val := range row {
		if val == 0 {
			if index == 0 {
				hasBegin = true
			} else if !hasBegin && !hasMiddle && index > 0 && index < lastIndex {
				hasMiddle = true
			}
		}

		if index == lastIndex {
			hasEnd = true
		}

		if val == 0 {
			counter++
		} else if counter > 0 {
			if hasBegin {
				maxBegin = counter
				hasBegin = false
			} else if hasMiddle && counter > maxMiddle {
				maxMiddle = counter
			}

			if index == lastIndex {
				hasEnd = false
			}

			counter = 0
		}
	}

	if counter > maxMiddle {
		maxMiddle = counter
	}

	halfMiddle := getHalf(maxMiddle)

	if maxBegin > 0 && maxBegin >= halfMiddle {
		return maxBegin
	}

	if hasEnd && counter >= halfMiddle {
		return counter
	}

	return halfMiddle
}
