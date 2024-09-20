package main

import "fmt"

func gcdOfString(str1 string, str2 string)string{
	//check if the concatenation of two string is different order is equal or not
	if str1+str2 != str2 +str1{
		return ""
	}
	//calculate the length of str1 and str2 
	gcd:=gcd(len(str1), len(str2))
	//returns the substring of str1 from index 0 to the gcd length, this substring is the larges potential divisor string 
	return str1[0:gcd]
}
//this function takes two integer and returns their gcd
func gcd(a,b int)int{
	//this loop uses the Elucidean algorithm to calculate gcd
	for b!=0{
		a,b=b,a%b 
	}
	return a 
}

func main() {
	gcd := gcdOfString("ABCABC", "ABC")
	fmt.Println(gcd)
}