package tools

import (
	"fmt"

	"github.com/schollz/wifiscan"
)

func GetStrength() int {
	wifis, err := wifiscan.Scan()
	if err != nil {
		fmt.Println(err)
	}
	return wifis[0].RSSI
}
