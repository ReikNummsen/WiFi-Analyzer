package app

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"wifi/internal/tools"
)

func WiFiScan(file string, value string) {
	fmt.Print("\nPinging ... ")
	ping := fmt.Sprintf("%f", tools.Ping("192.168.188.99")) // change to actual access point
	fmt.Print("Done!\n")
	fmt.Print("Getting signal strength ... ")
	strength := strconv.Itoa(tools.GetStrength())
	fmt.Print("Done!\n")
	// fmt.Print("Speedtesting ... ")
	// dls, uls := tools.GetBitrate()
	// download := fmt.Sprintf("%f", dls)
	// upload := fmt.Sprintf("%f", uls)
	// fmt.Print("Done!\n\n")

	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	line := strings.Join([]string{value, ping, strength}, ",")
	bytes, err := fmt.Fprintln(f, line)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(bytes, "bytes written successfully:")
	fmt.Println(strings.Join([]string{"Value: ", value, "\n", "Ping: ", ping, "ms\n", "Signal strength: ", strength, "dBm\n"}, ""))
}

func Interaction(file string) {
	var value string
	fmt.Println("Value:")
	fmt.Scan(&value)
	for value != "x" {
		WiFiScan(file, value)
		fmt.Println("Value:")
		fmt.Scan(&value)
	}
}
