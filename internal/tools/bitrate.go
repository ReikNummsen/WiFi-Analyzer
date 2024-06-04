package tools

import (
	"github.com/showwin/speedtest-go/speedtest"
)

func GetBitrate() (float64, float64) {
	var speedtestClient = speedtest.New()
	var downloadSpeed float64
	var uploadSpeed float64

	serverList, _ := speedtestClient.FetchServers()
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		s.DownloadTest()
		s.UploadTest()
		downloadSpeed = s.DLSpeed
		uploadSpeed = s.ULSpeed
		s.Context.Reset()
	}
	return downloadSpeed, uploadSpeed
}
