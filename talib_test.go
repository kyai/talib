package talib

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestTalib(t *testing.T) {
	klines := randKlines(3)
	fmt.Println("klines: ", klines)
	// klines = split(klines, 1)
	// reverse(klines)
	// fmt.Println("klines: ", klines)
}

func randKlines(num int) []Kline {
	klines := make([]Kline, num)
	for k, v := range klines {
		v.Open = randFloat(100, 200)
		v.High = randFloat(v.Open, 200)
		v.Low = randFloat(100, v.High)
		v.Close = randFloat(v.Low, v.Open)
		klines[k] = v
	}
	return klines
}

func randRange(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(max-min) + min
	return randNum
}

func randFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	num := rand.Float64()*(max-min+1) + min
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}
