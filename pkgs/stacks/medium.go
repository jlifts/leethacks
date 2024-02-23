package stacks

import (
"strconv"
)

// https://leetcode.com/problems/min-stack/
type MinStack struct {
  nums []int
  minVal []int
}


func Constructor() MinStack {
  return  MinStack {
    nums: []int{},
    minVal: []int{},
  }
}


func (this *MinStack) Push(val int)  {
  this.nums = append(this.nums, val)

  if len(this.nums) == 1 {
    this.minVal = []int{val}

    return
  }

  if val <= this.minVal[len(this.minVal)-1] {
    this.minVal = append(this.minVal, val)
  }
}


func (this *MinStack) Pop()  {
  num := this.nums[len(this.nums)-1]

  if num == this.minVal[len(this.minVal)-1] {
    this.minVal = this.minVal[:len(this.minVal)-1]
  }

  this.nums = this.nums[:len(this.nums)-1]
}


func (this *MinStack) Top() int {
  return this.nums[len(this.nums)-1]
}


func (this *MinStack) GetMin() int {
  return this.minVal[len(this.minVal)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// https://leetcode.com/problems/evaluate-reverse-polish-notation/
// O(n)

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	expressions := map[string]bool{"+": true, "-": true, "*": true, "/": true}

	for _, element := range tokens {
		if !expressions[element] {
			num, _ := strconv.Atoi(element)
			stack = append(stack, num)
		} else {
			num2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			switch element {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, int(num1/num2))
			}
		}
	}

	return stack[0]
}
