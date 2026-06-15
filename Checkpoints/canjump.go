package piscine

func CanJump(steps []uint) bool {
	if len(steps) == 0 {
		return false
	}

	pos := 0
	lastItem := len(steps) - 1
	size := len(steps)

	for pos != lastItem {
		next := pos + int(steps[pos])

		if next >= size || next == pos {
			return false
		}

		pos = next
	}
	return true
}
