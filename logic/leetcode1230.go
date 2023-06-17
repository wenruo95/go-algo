package logic

// leetcode 1232: https://leetcode.com/problems/check-if-it-is-a-straight-line/
func CheckStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}

	var k float64

	special := (coordinates[0][0] - coordinates[1][0]) == 0
	if !special {
		k = float64(coordinates[0][1]-coordinates[1][1]) / float64(coordinates[0][0]-coordinates[1][0])
	}

	for i := 2; i < len(coordinates); i++ {
		if special {
			if (coordinates[0][0] - coordinates[i][0]) != 0 {
				return false
			}
			continue
		}
		k2 := float64(coordinates[0][1]-coordinates[i][1]) / float64(coordinates[0][0]-coordinates[i][0])
		if k != k2 {
			return false
		}
	}
	return true
}
