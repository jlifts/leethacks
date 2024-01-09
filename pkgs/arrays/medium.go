package arrays

// Utilize bucket sort for O(n) time complexity, size of the array is equal to the size of hashmap
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	bucket := make([][]int, len(nums)+1)

	for i, v := range m {
		bucket[v] = append(bucket[v], i)
	}

	answer := make([]int, 0, k)

	for i := len(bucket) - 1; i >= 0; i-- {
		for _, v := range bucket[i] {
			if k > 0 {
				answer = append(answer, v)
				k--
			}
		}
	}

	return answer
}

func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	prefix := 1

	for i := 0; i < len(nums); i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	for i := len(nums) - 1; i >= 0; i-- {
		answer[i] *= postfix
		postfix *= nums[i]
	}

	return answer
}

func isValidSudoku(board [][]byte) bool {

	return false
}
