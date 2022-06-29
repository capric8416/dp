package main

func HourseJumpRecurse(targetXPos int, targetYPos int, steps int) int {
	if targetXPos < 0 || targetXPos > 9 || targetYPos < 0 || targetYPos > 10 || steps <= 0 {
		return 0
	}

	return process(steps, 0, 0, targetXPos, targetYPos, "")
}

func process(leftSteps int, sourceXpos int, sourceYPos int, targetXPos int, targetYPos int, indent string) int {
	if leftSteps == 0 {
		if sourceXpos == targetXPos && sourceYPos == targetYPos {
			// fmt.Printf("%v (%v,%v) 到达\n", indent, sourceXpos, sourceYPos)
			return 1
		}
		// fmt.Printf("%v (%v,%v) 不可达\n", indent, sourceXpos, sourceYPos)
		return 0
	}
	// fmt.Printf("%v (%v,%v) 尝试\n", indent, sourceXpos, sourceYPos)

	// 逆时针
	ways := 0
	if sourceXpos+1 < 9 && sourceYPos+2 < 10 { // 右上侧立"日"
		ways += process(leftSteps-1, sourceXpos+1, sourceYPos+2, targetXPos, targetYPos, indent+"\t")
	}
	if sourceXpos+2 < 9 && sourceYPos+1 < 10 { // 右上侧卧"日"
		ways += process(leftSteps-1, sourceXpos+2, sourceYPos+1, targetXPos, targetYPos, indent+"\t")
	}
	if sourceXpos+2 < 9 && sourceYPos-1 >= 0 { // 右下侧卧"日"
		ways += process(leftSteps-1, sourceXpos+2, sourceYPos-1, targetXPos, targetYPos, indent+"\t")
	}
	if sourceXpos+1 < 9 && sourceYPos-2 >= 0 { // 右下侧立"日"
		ways += process(leftSteps-1, sourceXpos+1, sourceYPos-2, targetXPos, targetYPos, indent+"\t")
	}
	if sourceXpos-1 >= 0 && sourceYPos-2 >= 0 { // 左下侧立"日"
		ways += process(leftSteps-1, sourceXpos-1, sourceYPos-2, targetXPos, targetYPos, indent+"\t")
	}
	if sourceXpos-2 >= 0 && sourceYPos-1 >= 0 { // 左下侧卧"日"
		ways += process(leftSteps-1, sourceXpos-2, sourceYPos-1, targetXPos, targetYPos, indent+"\t")
	}
	if sourceXpos-2 >= 0 && sourceYPos+1 < 10 { //左上侧卧"日"
		ways += process(leftSteps-1, sourceXpos-2, sourceYPos+1, targetXPos, targetYPos, indent+"\t")
	}
	if sourceXpos-1 >= 0 && sourceYPos+2 < 10 { // 左上侧立"日"
		ways += process(leftSteps-1, sourceXpos-1, sourceYPos+2, targetXPos, targetYPos, indent+"\t")
	}

	return ways
}
