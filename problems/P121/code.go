package P121

func maxProfit(prices []int) int {
	min := -1
	max := 0
	for i := 0; i < len(prices); i++ {
		if min == -1 {
			min = prices[i]
			continue
		}

		if prices[i] < min {
			min = prices[i]
			continue
		}

		if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max
}
