package leetcode

/*
单调栈的性质：
	1. 弹出操作完毕后，栈顶一定是左边第一个小于等于当前元素的
	2. 栈顶被弹出时，当前元素一定是右边第一个小于栈顶元素的
*/

/*
*
对于每一列，以它为高，左右尽量扩展宽度求出面积，需要找到左右第一个比它小的列。
*/
func largestRectangleArea(heights []int) int {
	heights = append(heights, 0) // 保证最终heights全部弹出，因为弹出时才做结算
	monoStack := Stack{-1}       // 单调栈用于确认左边界。最左-1永远不会弹出
	result := 0
	left := make([]int, len(heights)) // 记录左侧第一个比它小的下标
	for i := 0; i < len(heights); i++ {
		for len(monoStack) > 1 && heights[monoStack.Top()] >= heights[i] {
			toBalance := monoStack.Pop()
			area := (i - left[toBalance] - 1) * heights[toBalance]
			if result < area {
				result = area
			}
		}
		left[i] = monoStack.Top()
		monoStack.Push(i)
	}
	return result
}
