package main

/*
 * @description
 * How many ways are there from place A to place B after some steps
 * 问：机器人有多少种走法，
 * 可以花 steps 步从 currentPos 走到 targetPos？
 *
 * currentPos 范围是是[1, N]
 * 机器人每次只能走一步
 *
 * 如果 currentPos == 1，则下一步只能来到2
 * 如果 currentPos == N，则下一步只能来到N-1
 * 如果 1 < currentPos < N，则下一步既可以来到 currentPos-1，也可以来到 currentPos+1
 *
 * @maxPos: 右边界（可达），左边界为1（可达）
 * @currentPos: 当前机器人位置
 * @moveSteps: 机器人可以走多步
 * @targetPos: 机器人要前往的位置
 * @return: 总共有多少种走法
 */
func NumberOfWaysDp(maxPos, currentPos, moveSteps, targetPos int) int {
	// 如果剩余步数小于1，那么就不能动了，所以有0种走法
	// 如果右边界小于2，那么只能停留在1位置（左边界和右边界都是1），所以有0种走法
	// 如果当前位置小于左边界, 或者大于右边界，那么机器人就在有效区间外了，所以有0种走法
	// 如果目标位置小于左边界, 或者大于右边界，那么机器人就在有效区间外了，所以有0种走法
	if moveSteps < 1 || maxPos < 2 || currentPos < 1 || currentPos > maxPos || targetPos < 1 || targetPos > maxPos {
		return 0
	}

	// 递归函数有两个可变参数，所以是二维dp，然后再看这两个参数的变化范围
	// moveSteps 行，范围是 [0, moveSteps]
	// maxPos+1 列，范围是 [1, maxPos]，不使用第0列
	dp := make([][]int, moveSteps+1)
	for i := 0; i <= moveSteps; i++ {
		dp[i] = make([]int, maxPos+1)
	}

	// leftSteps == 0
	// stayPos == targetPos （对应递归终止条件，剩余步数为0，刚好停在目标位置）
	dp[0][targetPos] = 1

	// 上面已初始化第0行，所以应该从上到下，从左到右计算
	// 1 <= leftSteps <= steps
	// 1 <= stayPos <= maxPos
	for leftSteps := 1; leftSteps <= moveSteps; leftSteps++ {
		// stayPos == 1 （对应递归在第1个位置，只能向右移动）
		dp[leftSteps][1] = dp[leftSteps-1][2]
		// stayPos == maxPos （对应递归在最后一个位置，只能向左移动）
		dp[leftSteps][maxPos] = dp[leftSteps-1][maxPos-1]
		// 1 < stayPos < maxPos （对应递归在中间位置，可能向右，或者向左移动）
		for stayPos := 2; stayPos < maxPos; stayPos++ {
			dp[leftSteps][stayPos] = dp[leftSteps-1][stayPos+1] + dp[leftSteps-1][stayPos-1]
		}
	}

	// 对应递归函数首次被调用时的情形
	// return process(moveSteps, currentPos, targetPos, maxPos)
	return dp[moveSteps][currentPos]
}
