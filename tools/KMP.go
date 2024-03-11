package tools

func GenKmp(s string, k string) int {
	n := len(k)
	if n == 0 {
		return 0
	}
	next := make([]int, n)
	GetNext(next, k)
	j := 0
	for i := 0; i < len(s); i++ {
		for s[i] != k[j] && j > 0 {
			j = next[j-1]
		}
		if s[i] == k[j] {
			j++
		}
		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func GetNext(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}
