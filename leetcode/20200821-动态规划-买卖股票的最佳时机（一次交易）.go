package leetcode

func maxProfit(prices []int) int {
	var min, max, profit int
	for i, p := range prices {
		if i==0 {
			min = p
			continue
		}

		if min > p {
			max = p
			min = p
		}
		if max < p {
			max = p
		}

		if max - min > profit {
			profit = max -min
		}

	}
	return profit
}
