package main

func reverseVowels(s string) string {
	r := []byte(s)
	var index []int

	for i, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' ||
			v == 'A' || v == 'E' || v == 'I' || v == 'O' || v == 'U' {
			index = append(index, i)
		}
	}

	if len(index) >= 2 {
		i := 0
		j := len(index) - 1
		for i < j {
			r[index[i]], r[index[j]] = r[index[j]], r[index[i]]
			i++
			j--
		}
	}

	return string(r)
}
