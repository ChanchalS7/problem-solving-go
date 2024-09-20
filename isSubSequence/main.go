package main 
import "fmt"
func isSubsequence(s string, t string) bool {
	//take two pointer i and j
		i,j:=0,0
		// this loop run till the point value does not become equal to len of s and t
		for i!=len(s) && j!=len(t) {
			// if we find the s's element in t's element then increase our pointer by 1 , else just increase pointer respective to t
				if s[i] == t[j] {
					i,j = i+1, j+1
				}else{
					j++
				}
		}
		// return the value of i pointer equal to len of s
		return i ==len(s)
}

func main(){
	ans:=isSubsequence("abc","ahbgdc") 
	fmt.Println(ans)
}