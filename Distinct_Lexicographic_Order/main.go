package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	text := scanFileByUrl("https://gist.githubusercontent.com/Jekiwijaya/0b85de3b9ff551a879896dd78256e9b8/raw/e9d58da5d4df913ad62e6e8dd83c936090ee6ef4/gistfile1.txt")

	fmt.Println(getFirstOccurance(text))
	fmt.Print(getFirstLexicographicalOrder(text))
}

func scanFileByUrl(url string) string {
	response, _ := http.Get(url)
	data, _ := ioutil.ReadAll(response.Body)
	return string(data)
}

func getFirstOccurance(text string) string {
	var firstOccurance string
	var seen [255]bool
	for _, character := range text {
		if !seen[character] {
			firstOccurance += string(character)
			seen[character] = true
		}
	}

	return firstOccurance
}

func getFirstLexicographicalOrder(text string) string {
	var firstLexi string
	var count [255]int
	var onList [255]bool

	for _, character := range text {
		count[character]++
	}

	for _, character := range text {
		count[character]--
		// Not optimize, Optimize with map dictionary
		// if strings.Contains(firstLexi, string(character)) {
		// 	continue
		// }
		if onList[character] {
			continue
		}

		for len(firstLexi) > 0 && count[firstLexi[len(firstLexi)-1]] > 0 && byte(character) < firstLexi[len(firstLexi)-1] {
			onList[firstLexi[len(firstLexi)-1]] = false
			firstLexi = firstLexi[:len(firstLexi)-1]
		}
		firstLexi += string(character)
		onList[character] = true
	}

	return firstLexi
}
