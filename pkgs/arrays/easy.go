package arrays

import "strconv"

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)

	for index, num := range nums {
		if idx, ok := indexMap[target-num]; ok {
			return []int{idx, index}
		}
		indexMap[num] = index
	}

	return []int{}
}

func containsDuplicate(nums []int) bool {
	hm := make(map[int]bool)

	for _, v := range nums {
		if hm[v] {
			return true
		}

		hm[v] = true
	}

	return false
}

func isAnagram(s string, t string) bool {
	chars := make(map[rune]int)

	for _, v := range s {
		chars[v]++
	}
	for _, v := range t {
		chars[v]--
	}
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}

func majorityElement(nums []int) int {
	hm := make(map[int]int)

	for _, v := range nums {
		hm[v] += 1
	}

	num := 0
	res := 0

	for i, val := range hm {
		if val > num {
			num = val
			res = i
		}
	}

	return res
}

func fizzBuzz(n int) []string {
	var answer = make([]string, n, n)

	for i := 1; i <= n; i++ {
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		currentStr := ""

		if divisibleBy3 {
			currentStr += "Fizz"
		}

		if divisibleBy5 {
			currentStr += "Buzz"
		}

		if currentStr == "" {
			currentStr += strconv.Itoa(i)
		}

		answer[i-1] = currentStr
	}

	return answer
}
