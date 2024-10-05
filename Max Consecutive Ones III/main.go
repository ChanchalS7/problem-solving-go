package main

import "fmt"

func longestOnes(nums []int, k int) int{
 if len(nums)==0 {
	return 0
 }

 i,j:=0,0

 for i <len(nums) {
	if nums[i]==0 {
		k--
	}
	if k < 0 {
		if nums[j]==0 {
			k++
		}
		j++
	}
	i++
 }
 return i-j

}
func main(){
	nums :=[]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}
	 k := 3
	 fmt.Println(longestOnes(nums,k))
}