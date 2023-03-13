package top100

/*
单调栈,入栈时,保证当前块左侧所有雨水都结算完毕
*/

func trap(height []int) (result int) {
	monoStack := Stack{}
	for i := 0; i < len(height); i++ {
		for len(monoStack) > 0 && height[monoStack.Top()] <= height[i] {
			cur := monoStack.Pop()
			if len(monoStack) > 0 {
				deep := minInt(height[monoStack.Top()], height[i]) - height[cur]
				width := i - monoStack.Top() - 1
				result += deep * width
			}
		}
		monoStack.Push(i)
	}
	return
}
