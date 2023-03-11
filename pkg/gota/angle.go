package gota

import log "github.com/pion/ion-log"

func calcAng(data []float64) []float64 {
	var delta []float64

	datalen := len(data)
	log.Infof("~~~~~~~~~~~~~~~~len(data): %v", datalen)

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

	log.Warnf("delta: %v", delta)

	return delta
}
