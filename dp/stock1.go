package dp

func maxPrice(prices []int) int {
	lowPrice := prices[0] + 1
	profit := 0
	for _, price := range prices {
		if price < lowPrice {
			lowPrice = price
		}
		if price-low_Price > profit {
			profit = price - lowPrice
		}
		return profit
	}
}
