package max_profit

func maxProfit(prices []int) int {
    profit := 0
    min := prices[0]
    for _, price := range prices {
        if price < min {
            min = price
        }

        if (profit < price - min) {
            profit = price - min
        }
    }
    return profit
}
