package main

/*
 * 象棋棋盘
 * 整个棋盘放入第一象限，棋盘的最左上角是(0,0)位置
 * x 坐标 [0,8]，y 坐标 [0,9]
 * 求“马”从 (0,0) 出发，走 k 步，最后落在 (x,y) 位置的方法数有多少种？
 */

func HourseJumpDp(targetXPos int, targetYPos int, steps int) int {
	if targetXPos < 0 || targetXPos > 9 || targetYPos < 0 || targetYPos > 10 || steps <= 0 {
		return 0
	}

	// leftSteps 范围是 [0, steps]
	dp := make([][][]int, steps+1) // z 轴
	for i := 0; i < steps+1; i++ { // x 轴
		dp[i] = make([][]int, 9)
		for j := 0; j < 9; j++ { // y 轴
			dp[i][j] = make([]int, 10)
		}
	}

	// 初始化，递归结束条件
	dp[0][targetXPos][targetYPos] = 1

	// 从第0层开发计算
	for leftSteps := 1; leftSteps <= steps; leftSteps++ {
		for sourceXpos := 0; sourceXpos < 9; sourceXpos++ {
			for sourceYPos := 0; sourceYPos < 10; sourceYPos++ {
				// 逆时针
				ways := 0
				if sourceXpos+1 < 9 && sourceYPos+2 < 10 { // 右上侧立"日"
					ways += dp[leftSteps-1][sourceXpos+1][sourceYPos+2]
				}
				if sourceXpos+2 < 9 && sourceYPos+1 < 10 { // 右上侧卧"日"
					ways += dp[leftSteps-1][sourceXpos+2][sourceYPos+1]
				}
				if sourceXpos+2 < 9 && sourceYPos-1 >= 0 { // 右下侧卧"日"
					ways += dp[leftSteps-1][sourceXpos+2][sourceYPos-1]
				}
				if sourceXpos+1 < 9 && sourceYPos-2 >= 0 { // 右下侧立"日"
					ways += dp[leftSteps-1][sourceXpos+1][sourceYPos-2]
				}
				if sourceXpos-1 >= 0 && sourceYPos-2 >= 0 { // 左下侧立"日"
					ways += dp[leftSteps-1][sourceXpos-1][sourceYPos-2]
				}
				if sourceXpos-2 >= 0 && sourceYPos-1 >= 0 { // 左下侧卧"日"
					ways += dp[leftSteps-1][sourceXpos-2][sourceYPos-1]
				}
				if sourceXpos-2 >= 0 && sourceYPos+1 < 10 { //左上侧卧"日"
					ways += dp[leftSteps-1][sourceXpos-2][sourceYPos+1]
				}
				if sourceXpos-1 >= 0 && sourceYPos+2 < 10 { // 左上侧立"日"
					ways += dp[leftSteps-1][sourceXpos-1][sourceYPos+2]
				}
				dp[leftSteps][sourceXpos][sourceYPos] = ways
			}
		}
	}

	return dp[steps][0][0]
}
