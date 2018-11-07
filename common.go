package talib

import (
	"math"
)

func (kline Kline) isYin() bool {
	return kline.Close < kline.Open
}

func (kline Kline) isYang() bool {
	return kline.Close > kline.Open
}

// k1 contain of k2
func contain(k1, k2 Kline) bool {
	if k1.isYang() && k2.isYang() {
		if k1.Open < k2.Open && k1.Close > k2.Close {
			return true
		}
	}
	if k1.isYin() && k2.isYin() {
		if k1.Open > k2.Open && k1.Close < k2.Close {
			return true
		}
	}
	return false
}

// split klines by days
func split(klines []Kline, num int) []Kline {
	if len(klines) <= num {
		return klines
	}
	return klines[:num]
}

// reverse the klines
func reverse(klines []Kline) []Kline {
	for from, to := 0, len(klines)-1; from < to; from, to = from+1, to-1 {
		klines[from], klines[to] = klines[to], klines[from]
	}
	return klines
}

// 实体比较小，一般无上影线(或者特别短)，但是下影线特别长，一般是实体的两倍
func (kline Kline) isHammerLine() bool {
	Entity := math.Abs(kline.Close - kline.Open)
	var UpperShadow float64
	var LowerShadow float64
	if kline.isYang() {
		UpperShadow = kline.High - kline.Close
		LowerShadow = kline.Open - kline.Low
		if UpperShadow < 0.3*Entity && LowerShadow > 2*Entity {
			return true
		}
	} else if kline.isYin() {
		UpperShadow = kline.High - kline.Open
		LowerShadow = kline.Close - kline.Low
		if UpperShadow < 0.3*Entity && LowerShadow > 2*Entity {
			return true
		}
	}
	return false
}

//十字星
func (kline Kline) isStar() bool {
	return kline.isMinimum()
}

//长？ 短？ 略低于？ 如何定义？  暂时按照股票的标准
//极小
func (kline Kline) isMinimum() bool {
	Entity := math.Abs(kline.Close - kline.Open)
	if Entity/kline.Open < 0.006 {
		return true
	}
	return false
}

//小
func (kline Kline) isSmall() bool {
	Entity := math.Abs(kline.Close - kline.Open)
	Amplitude := Entity / kline.Open
	if Amplitude <= 0.015 && Amplitude >= 0.006 {
		return true
	}
	return false
}

//中
func (kline Kline) isMiddle() bool {
	Entity := math.Abs(kline.Close - kline.Open)
	Amplitude := Entity / kline.Open
	if Amplitude <= 0.035 && Amplitude > 0.015 {
		return true
	}
	return false
}

//大
func (kline Kline) isBig() bool {
	Entity := math.Abs(kline.Close - kline.Open)
	Amplitude := Entity / kline.Open
	if Amplitude > 0.035 {
		return true
	}
	return false
}

//附近？ 略低于？ 略高于？
// 0 ： 不在附近  -1 ：略低于 1 ： 略高于
// -1, 1 都代表在附近
func isNear(k1, k2 float64) int {
	Entity := math.Abs(k1 - k2)
	Amplitude := Entity / k1

	if k1 >= k2 && Amplitude <= 0.006 {
		return 1
	}

	if k1 < k2 && Amplitude <= 0.006 {
		return -1
	}

	return 0
}
