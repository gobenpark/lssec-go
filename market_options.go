package ebest_go

import (
	"fmt"
	"time"
)

// 현재가 호가 조회
type CurrentAskingPriceOption struct {
	Ticker string
}

func (CurrentAskingPriceOption) String() string {
	return "t1101"
}

func (c CurrentAskingPriceOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s"
  }
}`, c.String(),
		c.Ticker,
	)), nil
}

func (c CurrentAskingPriceOption) Path() string {
	return "/stock/market-data"
}

type CurrentPriceOption struct {
	Ticker string
}

func (CurrentPriceOption) String() string {
	return "t1102"
}

func (c CurrentPriceOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s"
  }
}`, c.String(),
		c.Ticker,
	)), nil
}

func (c CurrentPriceOption) Path() string {
	return "/stock/market-data"
}

type CurrentPriceMemoOption struct {
	Ticker string
}

func (CurrentPriceMemoOption) String() string {
	return "t1104"
}

func (c CurrentPriceMemoOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s"
  }
}`, c.String(),
		c.Ticker,
	)), nil
}

func (CurrentPriceMemoOption) Path() string {
	return "/stock/market-data"
}

// 주식피봇/디마크조회
type SearchPivotDemarkOption struct {
	Ticker string
}

func (SearchPivotDemarkOption) String() string {
	return "t1105"
}

func (c SearchPivotDemarkOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s"
  }
}`, c.String(),
		c.Ticker,
	)), nil
}

func (SearchPivotDemarkOption) Path() string {
	return "/stock/market-data"
}

// 시간외체결량
type OvertimeTransactionCountOption struct {
	Ticker     string
	DanChetime time.Time
	Index      int
}

func (OvertimeTransactionCountOption) String() string {
	return "t1109"
}

func (c OvertimeTransactionCountOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s"
	"dan_chetime": "%s"
	"idx": %d
  }
}`, c.String(),
		c.Ticker,
		c.DanChetime.Format("150405"),
		c.Index,
	)), nil
}

func (OvertimeTransactionCountOption) Path() string {
	return "/stock/market-data"
}

// 주식시간대별체결조회
type TimeOfDayTransactionOption struct {
	Ticker    string
	Volume    int
	StartTime time.Time
	EndTime   time.Time
	CtsTime   time.Time
}

func (TimeOfDayTransactionOption) String() string {
	return "t1301"
}

func (c TimeOfDayTransactionOption) Path() string {
	return "/stock/market-data"
}

func (c TimeOfDayTransactionOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"cvolume": %d,
	"starttime": "%s",
	"endtime": "%s",
	"cts_time": "%s"
  }
}`, c.String(),
		c.Ticker,
		c.Volume,
		c.StartTime.Format("1504"),
		c.EndTime.Format("1504"),
		c.CtsTime.Format("1504"),
	)), nil
}

// 주식분별주가조회
type MinuteOfDayPriceOption struct {
	Ticker string
	// 1:1분
	// 2:3분
	// 3:5분
	// 4:10분
	// 5:30분
	// 6:60분
	Gubun string
	Time  time.Time
	Count int
}

func (MinuteOfDayPriceOption) Path() string {
	return "/stock/market-data"
}

func (MinuteOfDayPriceOption) String() string {
	return "t1302"
}

func (c MinuteOfDayPriceOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"gubun": "%s",
	"time": "%s",
	"cnt": %d
  }
}`, c.String(),
		c.Ticker,
		c.Gubun,
		c.Time.Format("1504"),
		c.Count,
	)), nil
}
