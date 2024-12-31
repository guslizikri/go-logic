package main

import (
	"sort"
	"strconv"
)

// Ez

func romanToInt(s string) int {
	// Peta nilai simbol Romawi
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	total := 0
	length := len(s)

	for i := 0; i < length; i++ {
		currentVal := romanMap[s[i]]
		// Periksa apakah karakter berikutnya memiliki nilai yang lebih besar
		if i+1 < length && currentVal < romanMap[s[i+1]] {
			total -= currentVal
		} else {
			total += currentVal
		}
	}

	return total
}

func isPalindrome(x int) bool {
	var s string
	var sPalindrom string
	s = strconv.Itoa(x)
	left, right := 0, len(s)-1

	for left <= right {
		sPalindrom += string(s[right])
		right--
	}

	return s == sPalindrom
}

func kthDistinct(arr []string, k int) string {
	// Tahap 1: Hitung frekuensi setiap string
	freqMap := make(map[string]int)
	for _, str := range arr {
		freqMap[str]++
	}

	// Tahap 2: Cari string yang unik dan lacak urutan kemunculan
	distinctStrings := []string{}
	for _, str := range arr {
		if freqMap[str] == 1 {
			distinctStrings = append(distinctStrings, str)
		}
	}

	if len(distinctStrings) < k {
		return ""
	}
	return distinctStrings[k-1]
}

func reverseString(s string) bool {
	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	// result:=1
	// for i := num; i >= 1; i--{
	//   result *= i
	// }
	// return result

	return num * factorial(num-1)

}

func duplicateString(s string) string {
	check := make(map[byte]bool)
	length := len(s) - 1
	var result string
	for i := 0; i <= length; i++ {
		if !check[s[i]] {
			check[s[i]] = true
			result += string(s[i])
		}

	}
	return result
}

func twoSum(nums []int, target int) []int {
	// Buat hash map untuk menyimpan nilai dan indeksnya
	numMap := make(map[int]int)

	// Iterasi melalui array nums
	for i, num := range nums {
		// Hitung nilai yang kita cari
		complement := target - num

		// Cek apakah nilai tersebut sudah ada dalam hash map
		if idx, found := numMap[complement]; found {
			// Jika ada, kembalikan indeks dari kedua angka tersebut
			return []int{idx, i}
		}

		// Tambahkan angka dan indeksnya ke dalam hash map
		numMap[num] = i
	}

	// Jika tidak ditemukan, kembalikan array kosong
	return []int{}
}

// Medium
func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(current string, open int, close int)

	backtrack = func(current string, open int, close int) {
		if len(current) == 2*n {
			result = append(result, current)
			return
		}

		if open < n {
			backtrack(current+"(", open+1, close)
		}
		if close < open {
			backtrack(current+")", open, close+1)
		}
	}

	backtrack("", 0, 0)
	return result
}

func appendCharacters(s string, t string) int {
	// Pointer untuk string t
	tPointer := 0

	// Iterasi melalui string s
	for i := 0; i < len(s); i++ {
		if tPointer < len(t) && s[i] == t[tPointer] {
			tPointer++
		}
	}

	// Karakter yang perlu ditambahkan adalah panjang t yang tersisa setelah pointer
	return len(t) - tPointer
}

// Hard
func maxSatisfaction(satisfaction []int) int {
	sort.Ints(satisfaction)
	maxSatisfaction := 0
	currentTotal := 0

	for i := len(satisfaction) - 1; i >= 0; i-- {
		if currentTotal+satisfaction[i] > 0 {
			currentTotal += satisfaction[i]
			maxSatisfaction += currentTotal
		} else {
			break
		}
	}

	return maxSatisfaction
}

func minInsertions(s string) int {
	n := len(s)
	// dp[i][j] akan menyimpan panjang LPS di s[i:j+1]
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	lps := dp[0][n-1]
	return n - lps
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
