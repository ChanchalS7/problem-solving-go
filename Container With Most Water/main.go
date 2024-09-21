package main
import "fmt"
func maxArea(height []int)int {
 left, right := 0, len(height)-1
 maxArea :=0

 for left < right {
	//calculate the width between the two line
	width := right - left

	//find the minimum height of the two lines
	h:=min(height[left],height[right])
	//calculate the area formed by these two lines
	area := h*width 

	// update the maxArea if the current area is larger
	if area > maxArea {
		maxArea = area 
	}
	// move the pointer for the shorter line
	if height[left] < height[right]{
		left ++
	}else {
		right --
	}
 }
 return maxArea;
}

func main(){
height :=[]int{1,8,6,2,5,4,8,3,7}
fmt.Println(maxArea(height)) //outupt : 49
}