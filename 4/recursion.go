package main

func ConvertToLetterStringRecurse(digits string) int {
	if len(digits) == 0 {
		return 0
	}

	return process(digits, 0)
}

func process(digits string, index int) int {
	// 如果来到了最后，说明之前做的决定是有效的，所以返回1
	if index == len(digits) {
		return 1
	}

	// 将当前位置转换为字母
	// 0 不可以转成任何字母
	if digits[index] == '0' {
		return 0
	}
	ways := process(digits, index+1)

	// 看看下一个位置, 可不可以和当前位置的数字一起转换为字母
	if index+1 < len(digits) && (digits[index]-'0')*10+digits[index+1]-'0' < 27 {
		ways += process(digits, index+2)
	}

	return ways
}
