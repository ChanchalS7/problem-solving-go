package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var b []bool
	max := max(candies)
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			b = append(b, true)
		} else {
			b = append(b, false)
		}
	}
	return b
}
func max(n []int) int {
	max := 0
	for i := 0; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}
	return max
}
