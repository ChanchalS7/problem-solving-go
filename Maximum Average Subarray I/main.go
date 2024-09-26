package main

import "fmt"


func findMaxAverage(nums []int, k int) float64{
/*
1. use sliding window technique to maintain a sum of k elements
2. Initially, calculate the sum of the first k elements
3. Then slide the window across the array by subtracting the element that is left behind and adding
 	the new element that comes into that window.
 4. Keep track of the maximum sum encountered during the sliding process.
 5. Return the maximum average by dividing the maximum sum by k.
*/

//calculate the sum of first k  elements
windowSum:=0
for i:=0; i<k; i++ {
	windowSum +=nums[i]
}
//track the maximum sum(initially the sum of the first window)
maxSum:=windowSum

//slide the window across the array
for i:=k; i<len(nums); i++ {
	//update the window sum by subtracting the element that is left behind and adding the new elements
	windowSum = windowSum - nums[i-k]+nums[i]
	//track the maximum sum

	if windowSum > maxSum {
		maxSum = windowSum
	}
}
// calculate and return the maximum average
return float64(maxSum)/float64(k)

}

/*
func findMaxAverage(nums []int, k int) float64 {
	n := len(nums)
	if n == 0 || k == 0 {
		return 0.0
	}
	
	// Calculate the initial sum of the first window
	windowSum := 0
	for i := 0; i < k; i++ {
		windowSum += nums[i]
	}
	
	// Set initial maxSum to the first window sum
	maxSum := windowSum
	
	// Slide the window through the array
	for i := k; i < n; i++ {
		// Update the window sum by sliding: remove the element left behind, add the new one
		windowSum = windowSum - nums[i-k] + nums[i]
		
		// Track the maximum sum encountered
		maxSum = int(math.Max(float64(maxSum), float64(windowSum)))
	}
	
	// Return the maximum average
	return float64(maxSum) / float64(k)
}
*/


func main(){
nums:=[]int{1,12,-5,-6,50,3}
k:=4
//[5], k=1
fmt.Println(findMaxAverage(nums,k))
}