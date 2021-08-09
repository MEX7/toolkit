package main

import (
	"fmt"
	"strconv"

	"github.com/kl7sn/toolkit/stock"
)

func main() {
	lowest := 1.84 // 输入预计最低价
	highest := 2.6 // 最高价
	// nowPrice := 0.6 // 当前价格

	buyCoe := 0.035 // 建仓阶段系数，每跌 3.5% 买入一次
	buyTimes := 3   // 买入 3 次
	// buyMoney := 4000 // 第一次买入数量

	sellCoe := 0.01 // 卖出系数，需要在达到最高价时进行最后一次卖出
	sellTimes := 3  // 卖出 3 次
	// sellMoney := 4000

	knockCoe := 0.01 // 用 2000 对敲
	knockStart := stock.ReDivide(lowest, buyCoe, buyTimes)
	knockEnd := stock.ReMult(highest, sellCoe, sellTimes)

	dim, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", knockEnd/knockStart), 64)
	knockSum := float64(stock.LogX(1+knockCoe, dim))

	fmt.Printf("建仓开始价格：%.3f, 从 %.3f 开始每下跌2.5%% 执行金字塔买入\n", lowest, knockStart)
	fmt.Printf("对敲开始价格：%.3f\n", knockStart)
	fmt.Printf("网格数量：%d\n", int(knockSum))
	fmt.Printf("对敲结束价格：%.3f, 每上涨 1%% 执行金字塔卖出\n", knockEnd)
	fmt.Printf("清仓结束价格：%.3f\n", highest)

	firstBuyMoney := knockSum * 1000 / 3
	fmt.Printf("底仓资金：%d , 网格交易资金：%.2f 万, 建议第一次建仓: %.2f 万\n", int(firstBuyMoney*7), knockSum*0.2, firstBuyMoney/10000)

	firstBuy := stock.ReDivide(lowest, buyCoe, 2)
	firstBuyEarn := firstBuyMoney * ((knockStart - firstBuy) / firstBuy)

	secondBuy := stock.ReDivide(lowest, buyCoe, 1)
	secondBuyEarn := firstBuyMoney * 2 * ((knockStart - secondBuy) / secondBuy)

	// 建仓一次，买入 4000，经过两次对敲后卖完
	earn1Sum := knockSum / 3
	earn1 := 2000*knockCoe*earn1Sum*(earn1Sum+1)/2 + firstBuyEarn
	earn1SellPrice := stock.Increase(knockStart, knockCoe, int(earn1Sum))
	fmt.Printf("建仓一次成功，清仓价格 %.2f，成本：%.2f 万，获利：%.2f 元, 收益率：%.2f %% \n", earn1SellPrice, firstBuyMoney/10000, earn1, earn1/firstBuyMoney*100)

	// 建仓两次，买入 4000
	earn2 := 2000*knockCoe*knockSum*(knockSum+1)/2 + firstBuyEarn + secondBuyEarn
	earn2SellPrice := stock.Increase(knockStart, knockCoe, int(knockSum))
	fmt.Printf("建仓两次成功，清仓价格 %.2f，成本：%.2f 万，获利：%.2f 元, 收益率：%.2f %% \n", earn2SellPrice, firstBuyMoney*3/10000, earn2, earn2*100/(firstBuyMoney*3))

	// 建仓三次，买入 8000
	// earn3 := 4000*buyCoe + 8000*buyCoe*2 + 160000*buyCoe*3
	// fmt.Printf("建仓三次成功，成本：4000 元，获利：%.2f 元, 收益率：%.2f %% \n", earn3, earn3/280)

}
