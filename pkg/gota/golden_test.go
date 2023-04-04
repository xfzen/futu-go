package gota

import (
	"fmt"
	"testing"

	"github.com/zeromicro/go-zero/core/logx"
)

func TestCalcGoldenRatio(t *testing.T) {
	// CalcGoldenRatio(250, 15)

	// 1000 750 500 250 100 75 50 25 10

	ratios, average := CalcGoldenRatio(1000, 15)
	fmt.Println(ratios)  // 输出每次计算的黄金分割之后的价格
	fmt.Println(average) // 输出黄金分割比例之和的平均值

	logx.Infof("\n")
}
