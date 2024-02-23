package binarysearch

// https://leetcode.com/problems/binary-search/
// O(log n) Iterative approach
func search(nums []int, target int) int {
    pt, pt1 := 0, len(nums)-1
    

    for pt <= pt1 {
        mid := pt + ((pt1 -pt) / 2)

        if nums[mid] < target {
            mid++
            pt = mid
        } else if nums[mid] > target {
            mid--
            pt1 = mid
        } else {
            return mid
        }
    }

    return -1
}

func recursiveBinSearch(nums []int, target int) int {
  if len(nums) == 0 {
    return -1
  }

  return binSearch(nums, target, 0, len(nums) -1)
}

func binSearch(nums []int, target, left, right int) int {
  mid := left + ((right - left) / 2)
  if left <= right {
    if nums[mid] > target {
      mid--
      return binSearch(nums, target, left, mid)
    } else if nums[mid] < target {
      mid++
      return binSearch(nums, target, mid, right)
    } else {
      return mid
    } 
  }

  return -1
}


