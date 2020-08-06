package main

func findSecretWord(wordlist []string, master *Master) {
	l := len(wordlist)
	a := make([][]int, l)
	for i := 0; i < l; i++ {
		a[i] = make([]int, l)
	}
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			if i == j {
				a[i][j] = 6
			} else {
				r := 0
				for k := 0; k < 6; k++ {
					if wordlist[i][k] == wordlist[j][k] {
						r++
					}
				}
				a[i][j] = r
				a[j][i] = r
			}
		}
	}

	result := make([]int, len(wordlist))
	for i := 0; i < len(wordlist); i++ {
		result[i] = i
	}
	for len(result) > 1 {
		n := compute(result, a)
		val := master.Guess(wordlist[n])
		if val == 6 {
			return
		}
		newResult := []int{}
		for i := 0; i < len(result); i++ {
			if a[n][result[i]] == val {
				newResult = append(newResult, result[i])
			}
		}
		result = newResult
	}
	master.Guess(wordlist[result[0]])
}

func compute(l []int, a [][]int) int {
	minv := len(a)
	minvIdx := -1
	for n, _ := range a {
		maxl := maxL(l, n, a)
		if maxl < minv {
			minvIdx = n
			minv = maxl
		}
	}
	return minvIdx
}

func maxL(l []int, n int, a [][]int) int {
	m := map[int]int{}
	ret := 1
	for _, v := range l {
		if val, ok := m[a[n][v]]; ok {
			m[a[n][v]] = val + 1
			ret = max(ret, m[a[n][v]])
		} else {
			m[a[n][v]] = 1
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
