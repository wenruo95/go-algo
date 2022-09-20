package algo

import "testing"

func TestGetSum(t *testing.T) {
	type testData struct {
		a, b int
		sum  int
	}

	datas := []*testData{
		{a: 1, b: 2, sum: 3},
		{a: 2, b: 3, sum: 5},
		{a: 0, b: 0, sum: 0},
		{a: 1, b: -1, sum: 0},
		{a: 2, b: -2, sum: 0},
		{a: -3, b: 5, sum: 2},
		{a: 3, b: -5, sum: -2},
		{a: -3, b: -5, sum: -8},
	}
	for _, data := range datas {
		if sum := GetSumOfTwoIntegers(data.a, data.b); sum != data.sum {
			t.Errorf("get_sum_of_two_integers a:%v b:%v sum:%v expect:%v", data.a, data.b, sum, data.sum)
		}
	}

}
