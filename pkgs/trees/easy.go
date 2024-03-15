package trees

import "math"

// Definition for a binary tree node.
 type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
 
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func invertTree(root *TreeNode) *TreeNode {
    if root != nil {
        root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
    }
    return root
}

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
 //DFS
 //Recursive O(n)
func DFSRecMaxDepth(root *TreeNode) int {
    if root != nil {
        return 1 + int(math.Max(float64(DFSRecMaxDepth(root.Left)), float64(DFSRecMaxDepth(root.Right))))
    }
    return 0
}

//Iterative
type StackItem struct {
	TreeNode *TreeNode
	Depth    int
}

func IterMaxDepth(root *TreeNode) int {
    if root != nil {
        var stack []*StackItem
	    stack = append(stack, &StackItem{root, 1})
        res := 0

        for len(stack) > 0 {
            l := len(stack)
            item := stack[l-1]
            stack = stack[:l-1]
            
            node, depth := item.TreeNode, item.Depth
            res = int(math.Max(float64(res), float64(depth)))

            if node.Left != nil {
                stack = append(stack, &StackItem{node.Left, depth + 1})
                
            }
            if node.Right != nil {
                stack = append(stack, &StackItem{node.Right, depth + 1})
                
            }
        }
        return res
    }
    return 0
}

//BFS - level order traversal - using queues
// Iterative O(n)
func BFSMaxDepth(root *TreeNode) int {
    if root != nil {
        level := 0
        q := []*TreeNode{root}

        for len(q) > 0 {
            for _, node := range q {
                q = q[1:]
                if node.Left != nil {
                    q = append(q, node.Left)
                }
                if node.Right != nil {
                    q = append(q, node.Right)
                }
            }
            level++
        }

        return level
    }
    return 0
}


// https://leetcode.com/problems/diameter-of-binary-tree/
func diameterOfBinaryTree(root *TreeNode) int {
    maxD := 0
    var dfs func(node *TreeNode) int

    dfs = func(node *TreeNode) int {
      if root != nil {
        l := dfs(root.Left)
        r := dfs(root.Right)
        localMax := l + r
        maxD = max(maxD, localMax)

        return 1 + max(r, l)
      }

      return 0
    }
    dfs(root)
    return maxD
}