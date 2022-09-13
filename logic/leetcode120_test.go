package logic

import "testing"

func TestMaxProfit(t *testing.T) {

	type testData struct {
		prices []int
		result int
	}
	datas := []*testData{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			result: 5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			result: 0,
		},
		{
			prices: []int{1, 2},
			result: 1,
		},
	}

	for _, data := range datas {
		if result := MaxProfit(data.prices); result != data.result {
			t.Errorf("max_profix prices:%+v result:%v expect:%v", data.prices, result, data.result)
		}
		if result := MaxProfit2(data.prices); result != data.result {
			t.Errorf("max_profix2 prices:%+v result:%v expect:%v", data.prices, result, data.result)
		}
		if result := MaxProfit3(data.prices); result != data.result {
			t.Errorf("max_profix3 prices:%+v result:%v expect:%v", data.prices, result, data.result)
		}

	}

}

func BenchmarkMaxProfit(b *testing.B) {
	benchmarkMaxProfit(b, MaxProfit)
}

func BenchmarkMaxProfit2(b *testing.B) {
	benchmarkMaxProfit(b, MaxProfit2)
}

func BenchmarkMaxProfit3(b *testing.B) {
	benchmarkMaxProfit(b, MaxProfit3)
}

func benchmarkMaxProfit(b *testing.B, fn func([]int) int) {
	type testData struct {
		prices []int
		result int
	}
	datas := []*testData{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			result: 5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			result: 0,
		},
		{
			prices: []int{1, 2},
			result: 1,
		},
	}

	for i := 0; i < b.N; i++ {
		for j := 0; j < len(datas); j++ {
			fn(datas[j].prices)
		}
	}
}
