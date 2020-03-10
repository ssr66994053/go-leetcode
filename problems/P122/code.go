package P121

func maxProfit(prices []int) int {
	min := -1
	max := -1
	p := 0
	for i := 0; i < len(prices); i++ {
		if min == -1 || (prices[i] < min && max == -1) {
			min = prices[i]
			continue
		}

		if prices[i] > max {
			max = prices[i]
		} else {
			p += max - min
			min = prices[i]
			max = -1
		}
	}

	if max != -1 {
		p += max - min
	}

	return p
}
