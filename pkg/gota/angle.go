package gota

import (
	"github.com/zeromicro/go-zero/core/logx"
)

func calcAng(data []float64) []float64 {
	var delta []float64

	datalen := len(data)
	logx.Infof("~~~~~~~~~~~~~~~~len(data): %v", datalen)

	for i := 0; i < datalen; i++ {
		// filter first data
		if i == 0 {
			continue
		}

		last := data[i-1]
		now := data[i]
		diff := now - last

		delta = append(delta, diff)
	}

	logx.Infof("delta: %v", delta)

	return delta
}
