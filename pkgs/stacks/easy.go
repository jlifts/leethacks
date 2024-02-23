package stacks

// Valid Parenthesis https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
  if len(s) == 0 || len(s)%2 == 1 {
    return false
  }
  // O(n)
  pairs := map[rune]rune{
    '(': ')',
    '{': '}',
    '[': ']',
  }
  stack := []rune{}

  for _, c := range s {
    if _, ok := pairs[c]; ok {
      stack = append(stack, c)
     } else if len(stack) == 0 || pairs[stack[len(stack)-1]] != c {
       return false
     } else {
      stack = stack[:len(stack)-1]
    }
  }

  return len(stack) == 0
}
