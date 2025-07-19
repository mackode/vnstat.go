package main

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"math"
	"os/exec"
	"regexp"
)

func vnstat() (float64, float64, error) {
	rx := float64(0)
	tx := float64(0)

	cmd := exec.Command("ssh", "-p", "8022", "admin@192.168.0.1", "vnstat", "-i", "eth0", "-tr", "2")
	output, err := cmd.Output()
	if err != nil {
		return rx, tx, err
	}
	rateRex := regexp.MustCompile(`(?m)^\s+([rt]x)\s+([\d.]+)\s+(\S+)`)
	matches := rateRex.FindAllStringSubmatch(string(output), 2)

	for _, match := range matches {
		if match[1] == "rx" {
			rx, err = toBits(match[2], match[3])
		} else if match[1] == "tx" {
			tx, err = toBits(match[2], match[3])
		} else {
			return rx, tx, fmt.Errorf("Unknown entry: %s", match[1])
		}
		if err != nil {
			return rx, tx, err
		}
	}

	return rx, tx, nil
}

func toBits(str string, unit string) (float64, error) {
	s := str + string(unit[0])
	i, err := humanize.ParseBytes(s)
	return float64(i), err
}

func toBitRate(bps float64) string {
	return humanize.Bytes(uint64(bps)) + "it/sec"
}

func speedFromRate(x float64) float64 {
	return math.Sqrt(x / 1000.0)
}