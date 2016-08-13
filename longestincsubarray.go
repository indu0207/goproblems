package main

import "fmt"

func main() {

	input := [10]int{3, 6, 2, 5, 7, 8, 9, 1, 14, 17}
	maxlength := 1
	i := 0

	var j int

	for (i + maxlength) < len(input) {

		skipped := false
		for j = i + maxlength; j >= i+1; j-- {

			if j > 0 {
				fmt.Printf("%v\n", input)
				fmt.Printf("i = %d , j= %d , maxlength=%d,a[j]=%d,a[j-1]=%d\n", i, j, maxlength, input[j], input[j-1])
			}
			if (j > 0) && (input[j] < input[j-1]) {
				skipped = true
				break
			}

		}

		if skipped {
			fmt.Printf("Skipped\n\n")
			i = j
		} else {
			fmt.Printf("Next iteration\n\n")
			maxlength = maxlength + 1
		}

	}
	fmt.Printf("The max length is %d", maxlength)
}
