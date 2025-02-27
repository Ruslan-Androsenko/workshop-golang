package second

func maxDistToClosets(seats []int) int {
	var (
		maximum, counter, i int
		lastIndex           = len(seats) - 1
	)

	if seats[i] == 0 {
		for seats[i] == 0 {
			i++
			counter++
		}

		maximum = counter
		counter = 0
	}

	for ; i < len(seats); i++ {
		current := seats[i]

		if i == lastIndex && current == 0 {
			counter++
			maximum = getMax(maximum, counter)
			break
		}

		if current == 1 {
			counter = 0
		} else {
			counter++
			maximum = getMax(maximum, getHalf(counter))
		}
	}

	return maximum
}
