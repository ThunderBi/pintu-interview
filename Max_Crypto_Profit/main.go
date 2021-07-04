package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	prices := scanFileByUrl("https://gist.githubusercontent.com/Jekiwijaya/c72c2de532203965bf818e5a4e5e43e3/raw/2631344d08b044a4b833caeab8a42486b87cc19a/gistfile1.txt")

	result := findMaximumProfit(prices[:])

	fmt.Print(result)
}

func scanFileByUrl(url string) []int {
	response, _ := http.Get(url)

	data, _ := ioutil.ReadAll(response.Body)

	var result []int
	for _, value := range strings.Split(string(data), " ") {
		valueInt, _ := strconv.Atoi(value)
		result = append(result, valueInt)
	}
	return result
}

func findMaximumProfit(prices []int) int {
	minimumPrice := prices[0]
	maximumProfit := 0

	for _, price := range prices[1:] {
		minimumPrice = min(minimumPrice, price)
		maximumProfit = max(maximumProfit, price-minimumPrice)
	}

	return maximumProfit
}

func min(value1, value2 int) int {
	if value1 <= value2 {
		return value1
	}
	return value2
}

func max(value1, value2 int) int {
	if value1 <= value2 {
		return value2
	}
	return value1
}
