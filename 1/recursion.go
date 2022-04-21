package main

func NumberOfWaysRecurse(maxPos, currentPos, moveSteps, targetPos int) int {
	if moveSteps < 1 || maxPos < 2 || currentPos < 1 || currentPos > maxPos || targetPos < 1 || targetPos > maxPos {
		return 0
	}

	return process(moveSteps, currentPos, targetPos, maxPos)
}

func process(leftSteps, currentPos, targetPos, maxPos int) int {
	// 如果剩余步数为0，刚好停在目标位置，则此种走法可行
	if leftSteps == 0 {
		if currentPos == targetPos {
			return 1
		}
		return 0
	}

	// 只能向右移动
	if currentPos == 1 {
		return process(leftSteps-1, currentPos+1, targetPos, maxPos)
	}
	// 只能向左移动
	if currentPos == maxPos {
		return process(leftSteps-1, currentPos-1, targetPos, maxPos)
	}
	// 可以向右移动，或者向左移动
	return process(leftSteps-1, currentPos+1, targetPos, maxPos) + process(leftSteps-1, currentPos-1, targetPos, maxPos)
}
