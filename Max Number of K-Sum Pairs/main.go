package main

import (
	"fmt"
	"sort"
)

func maxOperations(nums []int, k int) int{
sort.Ints(nums)
//two pointer approach to find the solutions
left,right:=0, len(nums)-1
//variable which stores the count the number of operation and return the value
ans:=0
//runt till the value of left point is lesser than the right pointer
for left < right {
	if nums[left] + nums[right] == k {
		ans++; right--; left++
	}else if nums[left] + nums[right] > k {
		right--;
	}else{
		left++;
	}
}
//return the number of operations
return ans
}
/*
ans, obj:=0, make(map[int] int)
for _, num:= range nums{
 diff:=k-num
 if obj[diff] > 0 {
  ans++
  obj[diff]--
 }else{
 obj[num]++
 }
}
 return ans;
*/

func main(){
	nums:=[]int{1,2,3,4}
	k:=5
 fmt.Println(maxOperations(nums,k))
}