func gcdOfStrings(str1 string, str2 string) string {
	a := minDivider(str1)
	b := minDivider(str2)
	if a != b {
		return ""
	}
	c := len(str1) / len(a)
	d := len(str2) / len(a)
	g := gcd(c, d)
	res := ""
	for i := 0; i < g; i++ {
		res += a
	}
	return res
}

func minDivider(s string) string {
	for i := 1; i <= len(s); i++ {
		if isDivider(s[:i], s) {
			return s[:i]
		}
	}
	return ""
}

func isDivider(a string, s string) bool {
	if len(s)%len(a) != 0 {
		return false
	}
	for i := len(a); i < len(s); i++ {
		if s[i] != a[i%len(a)] {
			return false
		}
	}
	return true
}

func gcd(a int, b int) int {
	if a == b {
		return a
	} else if a < b {
		return gcd(b, a)
	} else {
		return gcd(a-b, b)
	}
}

/*
func gcdOfStrings(str1 string, str2 string) string {
    gcdValue := gcd(len(str1), len(str2))

    gcdStr := str1[:gcdValue]

    var newStr1, newStr2 string

    for i := 0; i < len(str1)/gcdValue; i++ {
        newStr1 += gcdStr
    }

    if newStr1 != str1 {
        return ""
    }

    for i := 0; i < len(str2)/gcdValue; i++ {
        newStr2 += gcdStr
    }

    if newStr2 != str2{
        return ""
    }

    return gcdStr
}

func gcd(a, b int) int {
      for b != 0 {
              t := b
              b = a % b
              a = t
      }

      return a
}




func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	gcd := gcd(len(str1), len(str2))
	return str1[0:gcd]

}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}











*/