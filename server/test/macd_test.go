package test

import (
	"fmt"
	"testing"
)

// 计算EMA
func calculateEMA(data []float64, period int) []float64 {
	if len(data) < period {
		return make([]float64, len(data))
	}

	alpha := 2.0 / float64(period+1)
	ema := make([]float64, len(data))

	// 第一个EMA是前period个值的SMA
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += data[i]
	}
	ema[period-1] = sum / float64(period)

	// 计算后续EMA
	for i := period; i < len(data); i++ {
		ema[i] = data[i]*alpha + ema[i-1]*(1-alpha)
	}

	return ema
}

// 计算MACD
func CalculateMACD(closePrices []float64, fastPeriod, slowPeriod, signalPeriod int) ([]float64, []float64, []float64) {
	if len(closePrices) < slowPeriod+signalPeriod {
		return nil, nil, nil
	}

	// 计算快慢EMA
	ema12 := calculateEMA(closePrices, fastPeriod)
	ema26 := calculateEMA(closePrices, slowPeriod)

	// 计算DIF线
	dif := make([]float64, len(closePrices))
	for i := 0; i < len(closePrices); i++ {
		dif[i] = ema12[i] - ema26[i]
	}

	// 计算信号线(DEA) - DIF的EMA
	dea := calculateEMA(dif, signalPeriod)

	// 计算MACD柱状图
	macdHist := make([]float64, len(closePrices))
	for i := 0; i < len(closePrices); i++ {
		macdHist[i] = dif[i] - dea[i]
	}

	return dif, dea, macdHist
}

func TestMacd(t *testing.T) {
	// 示例数据
	closePrices := []float64{
		100, 101, 102, 103, 104, 105, 106, 107, 108, 109,
		110, 111, 112, 113, 114, 115, 116, 117, 118, 119,
		120, 121, 122, 123, 124, 125, 126, 127, 128, 129,
		130, 131, 132, 133, 134, 135, 136, 137, 138, 139,
	}

	// 计算MACD (12, 26, 9)
	dif, dea, macd := CalculateMACD(closePrices, 12, 26, 9)

	fmt.Println("MACD计算结果 (12, 26, 9):")
	fmt.Println("天数\t收盘价\tDIF(12-26)\tDEA(DIF-9)\tMACD柱")
	for i := 25; i < len(closePrices); i++ {
		fmt.Printf("%d\t%.2f\t%.4f\t%.4f\t%.4f\n",
			i+1, closePrices[i], dif[i], dea[i], macd[i])
	}

	// 演示不同signalPeriod的影响
	fmt.Println("\n不同signalPeriod的比较:")
	periods := []int{5, 9, 12, 14}
	for _, period := range periods {
		_, dea, _ := CalculateMACD(closePrices, 12, 26, period)
		fmt.Printf("SignalPeriod=%d: 最新DEA=%.4f\n", period, dea[len(dea)-1])
	}
}
