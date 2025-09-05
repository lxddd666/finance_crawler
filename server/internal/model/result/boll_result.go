package result

// BollResult 布林带计算结果
type BollResult struct {
	Timestamp         int64   `json:"timestamp"`
	MiddleBand        float64 // 中轨（移动平均线）
	UpperBand         float64 // 上轨（移动平均线 + 两倍标准差）
	LowerBand         float64 // 下轨（移动平均线 - 两倍标准差）
	StandardDeviation float64 // 标准差
	ClosePrice        float64 // 当前时间
}
