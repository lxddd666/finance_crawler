package stock

import (
	"fmt"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/model/entity"
)

// Reverse 查询股票倒叙
func ReverseKline(slice []*entity.FinanceKline) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func GetCode(code string, exchange string) (completeCode string) {
	return fmt.Sprintf("%s%s", gstr.ToLower(exchange), code)
}
