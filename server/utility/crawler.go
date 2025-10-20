package utility

func BoolTrueInt(flag bool) int {
	if flag {
		return 1
	} else {
		return 2
	}
}

func BoolIntBoll(flag int) bool {
	if flag == 1 {
		return true
	} else {
		return false
	}
}
