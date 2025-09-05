package consts

const (
	KlineTypeMinutes   = 1 // 1分钟
	KlineType5Minutes  = 2 // 5分钟
	KlineType15Minutes = 3 // 15分钟
	KlineType30Minutes = 4 // 30分钟
	KlineTypeHours     = 5 // 1小时
	//Kline2Hours        = 6 // 2小时（不支持）
	//Kline4Hours        = 7 // 4小时（不支持）
	KlineTypeDay   = 8  // 日K
	KlineTypeWeek  = 9  // 周K
	KlineTypeMonth = 10 // 月K
)

const (
	KlineTimestampEndNow = 0 // 传0表示从当前最新的交易日往前查k线
)

const (
	AdjustTypeExRights = 0 // 除权
	//AdjustTypeCurrentRight  = 1 // 当前权 目前不支持
)
