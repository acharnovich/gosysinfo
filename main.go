package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	meminfo, err := os.Open("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	defer meminfo.Close()

	var total, available uint64
	for {
		var key, value string
		_, err := fmt.Fscanf(meminfo, "%s %s kB\n", &key, &value)
		if err != nil {
			break
		}
		val, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}
		switch key {

		case "MemTotal:":
			available = uint64(val) * 1024

			total = uint64(val) * 1024
		case "MemAvailable:":
			available = uint64(val) * 1024
		}
	}

	fmt.Printf("Total RAM: %d MB\n", total/1024/1024)
	fmt.Printf("Available RAM: %d MB\n", available/1024/1024)

}
