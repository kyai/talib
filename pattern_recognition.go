package talib

import (
	"math"
)

const (
	KRISE = 100
	KFALL = -100
	KFLAT = 0
)

// Two Crows
func CDL2CROWS(klines []Kline) int {
	if klines[0].isYang() && klines[1].isYin() && klines[2].isYin() {
		if klines[1].Close > klines[0].Close {
			if klines[2].Open > klines[1].Open && klines[2].Close < klines[1].Close {
				return KFALL
			}
		}
	}
	return 0
}

// Three Black Crows
func CDL3BLACKCROWS(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYin() && klines[2].isYin() {
		if klines[0].Open > klines[1].Open && klines[1].Open > klines[2].Open {
			if klines[0].Close > klines[1].Close && klines[1].Close > klines[2].Close {
				return KFALL
			}
		}
	}
	return 0
}

// Three Inside Up/Down
func CDL3INSIDE(klines []Kline) int {
	if klines[2].Close > klines[0].Open {
		if contain(klines[0], klines[1]) {
			return KRISE
		}
	}
	return 0
}

// Three-Line Strike
func CDL3LINESTRIKE(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYin() && klines[2].isYin() && klines[3].isYang() {
		if klines[0].Open > klines[1].Open && klines[1].Open > klines[2].Open && klines[3].Open > klines[0].Open {
			if klines[0].Close > klines[1].Close && klines[1].Close > klines[2].Close && klines[3].Close < klines[2].Close {
				return KRISE
			}
		}
	}

	if klines[0].isYang() && klines[1].isYang() && klines[2].isYang() && klines[3].isYin() {
		if klines[0].Open < klines[1].Open && klines[1].Open < klines[2].Open && klines[2].Open < klines[3].Open {
			if klines[0].Close < klines[1].Close && klines[1].Close < klines[2].Close && klines[3].Close < klines[0].Close {
				return KFALL
			}
		}
	}
	return 0
}

// Three Outside Up/Down
func CDL3OUTSIDE(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYang() && klines[2].isYang() {
		if klines[0].Open < klines[1].Open && klines[1].Open < klines[2].Open {
			if klines[0].Close > klines[1].Close && klines[1].Close < klines[2].Close {
				return KRISE
			}
		}
	}

	if klines[0].isYang() && klines[1].isYin() && klines[2].isYin() {
		if klines[0].Open < klines[1].Open && klines[1].Open > klines[2].Open {
			if klines[0].Close > klines[1].Close && klines[1].Close > klines[2].Close {
				return KFALL
			}
		}
	}
	return 0
}

// Three Stars In The South
func CDL3STARSINSOUTH(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYin() && klines[2].isYin() {
		if klines[0].Open > klines[1].Open && klines[1].Open > klines[2].Open {
			if klines[0].Close > klines[1].Close && klines[1].Close < klines[2].Close {
				return KRISE
			}
		}
	}
	return 0
}

// Three Advancing White Soldiers
func CDL3WHITESOLDIERS(klines []Kline) int {
	if klines[0].isYang() && klines[1].isYang() && klines[2].isYang() {
		if klines[0].Open < klines[1].Open && klines[1].Open < klines[2].Open {
			if klines[0].Close < klines[1].Close && klines[1].Close < klines[2].Close {
				return KRISE
			}
		}
	}
	return 0
}

// Abandoned Baby
func CDLABANDONEDBABY(klines []Kline) int {
	return 0
}

// Advance Block
func CDLADVANCEBLOCK(klines []Kline) int {
	if klines[0].isYang() && klines[1].isYang() && klines[2].isYang() {
		if klines[0].Open < klines[1].Open && klines[1].Open < klines[2].Open {
			if klines[0].Close < klines[1].Close && klines[1].Close < klines[2].Close {
				return KRISE
			}
		}
	}
	return 0
}

// Belt-hold
func CDLBELTHOLD(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYang() {
		if klines[1].Open == klines[1].Low && isNear(klines[1].Close, klines[1].High) != 0 {
			return KRISE
		}
	}
	return 0
}

// Breakaway
func CDLBREAKAWAY(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYin() && klines[4].isYang() {
		if klines[4].Close < klines[0].Close && klines[4].Close > klines[1].Open {
			return KRISE
		}
	}
	return 0
}

// Closing Marubozu
func CDLCLOSINGMARUBOZU(klines []Kline) int {
	if klines[0].isYang() {
		if klines[0].Low < klines[0].Open && klines[0].Close == klines[0].High {
			return KFLAT
		}
	}
	return 0
}

// Concealing Baby Swallow
func CDLCONCEALBABYSWALL(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYin() && klines[2].isYin() && klines[3].isYin() {
		if klines[0].High == klines[0].Open && klines[0].Close == klines[0].Low && klines[1].High == klines[1].Open && klines[1].Close == klines[1].Low {
			return KRISE
		}
	}
	return 0
}

// Counterattack
func CDLCOUNTERATTACK(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYang() {
		if klines[0].Open > klines[1].Open && klines[0].Close > klines[1].Close && klines[0].Close == klines[1].Open {
			return KRISE
		}
	}
	return 0
}

// Dark Cloud Cover
func CDLDARKCLOUDCOVER(klines []Kline) int {
	if klines[0].isYang() && klines[1].isYin() {
		if klines[0].Open < klines[1].Open && klines[0].Close < klines[1].Close {
			return KFALL
		}
	}
	return 0
}

// Doji
func CDLDOJI(klines []Kline) int {
	if klines[0].Open == klines[1].Close {
		return KFLAT ////
	}
	return 0
}

// Doji Star
func CDLDOJISTAR(klines []Kline) int {
	if klines[0].isStar() { ////
		return KRISE
	}
	return 0
}

// Dragonfly Doji
func CDLDRAGONFLYDOJI(klines []Kline) int {
	if klines[0].High == klines[0].Open && klines[0].High == klines[0].Close {
		return KFLAT
	}
	return 0
}

// Engulfing Pattern
//两日K线模式，以多头吞噬为例，第一日为阴线， 第二日阳线，第一日的开盘价和收盘价在第二日开盘价收盘价之内，但不能完全相同。
func CDLENGULFING(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYang() {
		if klines[0].Open < klines[1].Close && klines[0].Close > klines[1].Open {
			return KRISE
		}
	}
	return 0
}

// Evening Doji Star
// 三日K线模式，基本模式为暮星，第二日收盘价和开盘价相同，预示顶部反转。
func CDLEVENINGDOJISTAR(klines []Kline) int {
	if klines[0].isYang() && klines[2].isYin() {
		if klines[1].Open == klines[1].Close {
			return KFALL
		}
	}
	return 0
}

// Evening Star
// 三日K线模式，与晨星相反，上升趋势中, 第一日阳线，第二日价格振幅较小，第三日阴线，预示顶部反转。
func CDLEVENINGSTAR(klines []Kline) int {
	if klines[0].isYang() && klines[2].isYin() {
		if klines[1].isSmall() {
			return KFALL
		}
	}
	return 0
}

// Up/Down-gap side-by-side white lines
func CDLGAPSIDESIDEWHITE(klines []Kline) int {
	return 0
}

// Gravestone Doji
// 一日K线模式，开盘价与收盘价相同，上影线长，无下影线，预示底部反转。
func CDLGRAVESTONEDOJI(klines []Kline) int {
	if klines[0].Open == klines[0].Close {
		//什么叫长?
		if klines[0].Low == klines[0].Open && klines[0].High > 5 { //TODO

		}
	}
	return 0
}

// Hammer
func CDLHAMMER(klines []Kline) int {
	if klines[0].isHammerLine() {
		return KFLAT
	}
	return 0
}

// Hanging Man
// 一日K线模式，形状与锤子类似，处于上升趋势的顶部，预示着趋势反转。
func CDLHANGINGMAN(klines []Kline) int {
	if klines[0].isHammerLine() {
		return KFALL
	}
	return 0
}

// Harami Pattern
// 二日K线模式，多头母子为例，在下跌趋势中，第一日K线长阴， 第二日开盘价收盘价在第一日价格振幅之内，为阳线，预示趋势反转，股价上升。
func CDLHARAMI(klines []Kline) int {
	if klines[0].isYin() && klines[0].isBig() && klines[1].isYang() {
		if klines[1].Open > klines[0].Low && klines[1].Close < klines[0].High {
			return KRISE
		}
	}
	return 0
}

// Harami Cross Pattern
//二日K线模式，多头为例,与母子线类似，若第二日K线是十字线， 便称为十字孕线，预示着趋势反转。
func CDLHARAMICROSS(klines []Kline) int {
	if klines[0].isYin() && klines[0].isBig() {
		if klines[1].isStar() {
			return KRISE
		}
	}
	return 0
}

// High-Wave Candle
//三日K线模式，具有极长的上/下影线与短的实体，预示着趋势反转。
func CDLHIGHWAVE(klines []Kline) int {
	return 0
}

// Hikkake Pattern
//三日K线模式，与母子类似，第二日价格在前一日实体范围内, 第三日收盘价高于前两日，反转失败，趋势继续。
func CDLHIKKAKE(klines []Kline) int {
	return 0
}

// Modified Hikkake Pattern
//三日K线模式，与陷阱类似，上升趋势中，第三日跳空高开； 下跌趋势中，第三日跳空低开，反转失败，趋势继续。
func CDLHIKKAKEMOD(klines []Kline) int {
	return 0
}

// Homing Pigeon
// 二日K线模式，与母子线类似，不同的的是二日K线颜色相同， 第二日最高价、最低价都在第一日实体之内，预示着趋势反转。
func CDLHOMINGPIGEON(klines []Kline) int {

	return 0
}

// Identical Three Crows
// 三日K线模式，上涨趋势中，三日都为阴线，长度大致相等， 每日开盘价等于前一日收盘价，收盘价接近当日最低价，预示价格下跌。
func CDLIDENTICAL3CROWS(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYin() && klines[2].isYin() {
		if (klines[0].Open-klines[0].Close) == (klines[1].Open-klines[1].Close) && (klines[1].Open-klines[1].Close) == (klines[2].Open-klines[2].Close) {
			if klines[1].Open == klines[0].Close && klines[2].Open == klines[1].Close {
				if isNear(klines[0].Close, klines[0].Low) != 0 && isNear(klines[1].Close, klines[1].Low) != 0 && isNear(klines[2].Close, klines[2].Low) != 0 {
					return KFALL
				}
			}

		}
	}

	return 0
}

// In-Neck Pattern
// 二日K线模式，下跌趋势中，第一日长阴线， 第二日开盘价较低，收盘价略高于第一日收盘价，阳线，实体较短，预示着下跌继续。
func CDLINNECK(klines []Kline) int {
	if klines[0].isYin() && klines[0].isBig() && klines[1].isYang() && klines[1].isSmall() {
		if klines[1].Open < klines[0].Low && isNear(klines[1].Close, klines[0].Close) == 1 {
			return KFALL
		}
	}
	return 0
}

// Inverted Hammer
// 一日K线模式，上影线较长，长度为实体2倍以上， 无下影线，在下跌趋势底部，预示着趋势反转。
func CDLINVERTEDHAMMER(klines []Kline) int {
	if klines[0].isinvertedHammer() {
		return KFLAT
	}
	return 0
}

// Kicking
// 二日K线模式，与分离线类似，两日K线为秃线，颜色相反，存在跳空缺口。
func CDLKICKING(klines []Kline) int {
	return 0
}

// Kicking - bull/bear determined by the longer marubozu
// 二日K线模式，与反冲形态类似，较长缺影线决定价格的涨跌。
func CDLKICKINGBYLENGTH(klines []Kline) int {
	return 0
}

// Ladder Bottom
//五日K线模式，下跌趋势中，前三日阴线， 开盘价与收盘价皆低于前一日开盘、收盘价，第四日倒锤头，第五日开盘价高于前一日开盘价， 阳线，收盘价高于前几日价格振幅，预示着底部反转。
func CDLLADDERBOTTOM(klines []Kline) int {
	return 0
}

// Long Legged Doji
//一日K线模式，开盘价与收盘价相同居当日价格中部，上下影线长， 表达市场不确定性。
func CDLLONGLEGGEDDOJI(klines []Kline) int {
	Middle := (klines[0].High - klines[0].Low) / 2
	if klines[0].Close == klines[0].Open && klines[0].Close == Middle {
		// 如果影线很长
		return KFLAT
	}
	return 0
}

// Long Line Candle
//一日K线模式，K线实体长，无上下影线。
func CDLLONGLINE(klines []Kline) int {
	if klines[0].isBig() {
		if klines[0].isYin() {
			if klines[0].High == klines[0].Close && klines[0].Low == klines[0].Open {
				return KFLAT
			}
		}

		if klines[0].isYang() {
			if klines[0].High == klines[0].Open && klines[0].Low == klines[0].Close {
				return KFLAT
			}
		}
	}
	return 0
}

// Marubozu
//一日K线模式，上下两头都没有影线的实体， 阴线预示着熊市持续或者牛市反转，阳线相反。
func CDLMARUBOZU(klines []Kline) int {
	if klines[0].isYin() {
		if klines[0].High == klines[0].Close && klines[0].Low == klines[0].Open {
			return KFALL
		}
	}

	if klines[0].isYang() {
		if klines[0].High == klines[0].Open && klines[0].Low == klines[0].Close {
			return KFALL
		}
	}

	return 0
}

// Matching Low
// 二日K线模式，下跌趋势中，第一日长阴线， 第二日阴线，收盘价与前一日相同，预示底部确认，该价格为支撑位。
func CDLMATCHINGLOW(klines []Kline) int {
	if klines[0].isYin() && klines[0].isBig() && klines[1].isYin() && klines[1].Close == klines[0].Close {
		return KRISE
	}
	return 0
}

// Mat Hold
//五日K线模式，上涨趋势中，第一日阳线，第二日跳空高开影线， 第三、四日短实体影线，第五日阳线，收盘价高于前四日，预示趋势持续。
func CDLMATHOLD(klines []Kline) int {
	return 0
}

// Morning Doji Star
// 三日K线模式， 基本模式为晨星，第二日K线为十字星，预示底部反转。
func CDLMORNINGDOJISTAR(klines []Kline) int {
	if klines[0].isYin() && klines[2].isYang() && klines[1].isStar() && klines[2].isBig() {
		return KRISE
	}
	return 0
}

// Morning Star
// 三日K线模式，下跌趋势，第一日阴线， 第二日价格振幅较小，第三天阳线，预示底部反转。
func CDLMORNINGSTAR(klines []Kline) int {
	if klines[0].isYin() && klines[2].isYang() && klines[1].isSmall() && klines[2].isBig() {
		return KRISE
	}
	return 0
}

// On-Neck Pattern
//二日K线模式，下跌趋势中，第一日长阴线，第二日开盘价较低， 收盘价与前一日最低价相同，阳线，实体较短，预示着延续下跌趋势。
func CDLONNECK(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYang() {
		if klines[0].isBig() && klines[1].isSmall() && klines[1].Open < klines[0].Low && klines[1].Close == klines[0].Low {
			return KFALL
		}
	}
	return 0
}

// Piercing Pattern
//两日K线模式，下跌趋势中，第一日阴线，第二日开盘价低于前一日最低价， 收盘价处在第一日实体上部，预示着底部反转。
func CDLPIERCING(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYang() {
		if klines[1].Open < klines[0].Low && klines[1].Close > klines[0].Close && klines[1].Close < klines[0].Open {
			return KRISE
		}
	}
	return 0
}

// Rickshaw Man
// 一日K线模式，与长腿十字线类似， 若实体正好处于价格振幅中点，称为黄包车夫。
// 什么叫长?
func CDLRICKSHAWMAN(klines []Kline) int {
	return 0
}

// Rising/Falling Three Methods
func CDLRISEFALL3METHODS(klines []Kline) int {
	return 0
}

// Separating Lines
//二日K线模式，上涨趋势中，第一日阴线，第二日阳线， 第二日开盘价与第一日相同且为最低价，预示着趋势继续。
func CDLSEPARATINGLINES(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYang() {
		if klines[1].Open == klines[0].Open && klines[1].Open == klines[1].Low {
			return KRISE
		}
	}
	return 0
}

// Shooting Star
//一日K线模式，上影线至少为实体长度两倍， 没有下影线，预示着股价下跌
func CDLSHOOTINGSTAR(klines []Kline) int {
	Entity := math.Abs(klines[0].Close - klines[0].Open)

	if klines[0].isYang() {
		if klines[0].High-klines[0].Close >= 2*Entity && klines[0].Low == klines[0].Open {
			return KFALL
		}
	}

	if klines[0].isYin() {
		if klines[0].High-klines[0].Open >= 2*Entity && klines[0].Low == klines[0].Close {
			return KFALL
		}
	}

	return 0
}

// Short Line Candle
//一日K线模式，实体短，无上下影线
func CDLSHORTLINE(klines []Kline) int {
	if isNear(klines[0].Close, klines[0].Open) != 0 {
		if klines[0].isYang() {
			if klines[0].Low == klines[0].Open && klines[0].High == klines[0].Close {
				return KFLAT
			}
		}

		if klines[0].isYin() {
			if klines[0].High == klines[0].Open && klines[0].Low == klines[0].Close {
				return KFLAT
			}
		}
	}
	return 0
}

// Spinning Top
//一日K线，实体小。
func CDLSPINNINGTOP(klines []Kline) int {
	if isNear(klines[0].Close, klines[0].Open) != 0 {
		return KFLAT
	}
	return 0
}

// Stalled Pattern
//三日K线模式，上涨趋势中，第二日长阳线， 第三日开盘于前一日收盘价附近，短阳线，预示着上涨结束
func CDLSTALLEDPATTERN(klines []Kline) int {
	if klines[0].isYang() && klines[1].isYang() && klines[2].isYang() {
		if klines[1].isBig() && isNear(klines[2].Open, klines[1].Close) != 0 && klines[2].isSmall() {
			return KFLAT
		}
	}
	return 0
}

// Stick Sandwich
//三日K线模式，第一日长阴线，第二日阳线，开盘价高于前一日收盘价，第三日开盘价高于前两日最高价，收盘价于第一日收盘价相同。
func CDLSTICKSANDWICH(klines []Kline) int {
	if klines[0].isYin() && klines[0].isBig() && klines[1].isYang() && klines[2].isYin() {
		if klines[1].Open > klines[0].Close && klines[2].Open > klines[0].High && klines[2].Open > klines[1].High && klines[2].Close == klines[0].Close {
			return KFALL
		}
	}
	return 0
}

//一日k线 大致与蜻蜓十字相同，下影线长度长
// Takuri (Dragonfly Doji with very long lower shadow)
func CDLTAKURI(klines []Kline) int {
	return 0
}

// Tasuki Gap
func CDLTASUKIGAP(klines []Kline) int {
	if klines[0].isYang() && klines[1].isYang() && klines[2].isYin() {
		if klines[1].Open > klines[0].Close && klines[2].Close < klines[1].Open && klines[2].Close > klines[0].Close {
			return KRISE
		}
	}
	return 0
}

// Thrusting Pattern
// 二日K线模式，与颈上线类似，下跌趋势中，第一日长阴线，第二日开盘价跳空， 收盘价略低于前一日实体中部，与颈上线相比实体较长，预示着趋势持续。
func CDLTHRUSTING(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYang() && klines[0].isBig() {
		if klines[1].Open < klines[0].Low && isNear(klines[1].Close, 0.5*math.Abs(klines[0].Close-klines[0].Open)) == -1 {
			return KFALL
		}
	}
	return 0
}

// Tristar Pattern
func CDLTRISTAR(klines []Kline) int {
	if klines[0].isStar() && klines[1].isStar() && klines[2].isStar() {
		if klines[1].Close < klines[0].Close && klines[1].Close < klines[2].Close {
			return KRISE
		}
	}
	return 0
}

// Unique 3 River
func CDLUNIQUE3RIVER(klines []Kline) int {
	if klines[0].isYin() && klines[1].isYin() && klines[1].isHammerLine() && klines[2].isYang() {
		if contain(klines[1], klines[2]) && klines[2].Close > klines[3].Close {
			return KRISE
		}
	}
	return 0
}

// Upside Gap Two Crows
func CDLUPSIDEGAP2CROWS(klines []Kline) int {
	if klines[0].isYang() && klines[1].isYin() && klines[2].isYin() {
		if klines[0].Close < klines[1].Close && contain(klines[2], klines[1]) {
			return KRISE
		}
	}
	return 0
}

// Upside/Downside Gap Three Methods
func CDLXSIDEGAP3METHODS(klines []Kline) int {
	return 0
}
