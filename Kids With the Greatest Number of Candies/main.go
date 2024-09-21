package main

import "fmt"


func kidsWithCandies(candies []int, extraCandies int) []bool {
	var ans []bool
	max:= max(candies)
	for i:=0; i<len(candies); i++ {
		if candies[i] + extraCandies >=max {
			ans=append(ans,true)
		}else{
			ans=append(ans,false)
		}
	}
	return ans
}
//function to find out the maximum number
func max(n []int) int{
	max:=0
	for i:=0; i<len(n); i++ {
		if n[i] > max{
			max=n[i]
		}
	}
	return max
}
//main function body
func main(){
candies:=[]int{2,3,5,1,3}
extraCandies:=3
fmt.Println(kidsWithCandies(candies, extraCandies))

}