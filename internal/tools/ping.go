package tools

import (
	"github.com/go-ping/ping"
)

func Ping(host string) float32 {
	pinger, _ := ping.NewPinger(host)
	var avg float32

	pinger.SetPrivileged(true)
	pinger.Count = 10
	pinger.Run()
	pings := pinger.Statistics().Rtts
	mu := pinger.Statistics().AvgRtt
	std := pinger.Statistics().StdDevRtt
	clean_pings := []float32{}

	// Remove outliers:
	for ping := range pings {
		if pings[ping] < mu+2*std && pings[ping] > mu-2*std {
			clean_pings = append(clean_pings, float32(pings[ping].Microseconds())*0.001)
		}
	}

	findArraySum := func(arr []float32) float32 {
		var res float32
		for i := 0; i < len(arr); i++ {
			res += arr[i]
		}
		return res
	}
	avg = findArraySum(clean_pings) / float32(len(clean_pings))
	return avg
}
