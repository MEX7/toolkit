package main

import (
	"fmt"

	"github.com/kl7sn/toolkit/stock"
)

func main() {
	var (
		low  float64
		high float64
		now  float64
		coe  float64
		num  float64
	)
	fmt.Printf("当前股价: ")
	fmt.Scanln(&now)

	fmt.Printf("网格最低价: ")
	fmt.Scanln(&low)

	fmt.Printf("网格最高价: ")
	fmt.Scanln(&high)

	fmt.Printf("涨跌系数（例如涨跌3%%交易，则输入 3）: ")
	fmt.Scanln(&coe)

	fmt.Printf("单次交易股数: ")
	fmt.Scanln(&num)

	coe = coe / 100
	ups := stock.ReMult(now, high, coe)
	downs := stock.ReDivide(now, low, coe)

	cost := int(float64(ups) * now * num)
	reserved := int(float64(downs) * now * num)

	fmt.Printf("网格总数：%d, 下方网格：%d, 上方网格：%d\n", ups+downs, downs, ups)
	fmt.Printf("底仓数量：%d, 买入成本：%d, 预留资金：%d \n", ups*int(num), cost, reserved)
	fmt.Printf("网格收益 %.2f%%~%.2f%%", (high-now)*100/(2*now), (high-now)*100/now)

}
