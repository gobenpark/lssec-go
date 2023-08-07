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

//
//
//Element	한글명	type	Required	Length	Description
//t1302InBlock	t1302InBlock	Object	Y	-
//-shcode	단축코드	String	Y	6
//-gubun	작업구분	String	Y	1	0:30초
//1:1분
//2:3분
//3:5분
//4:10분
//5:30분
//6:60분
//-time	시간	String	Y	6	처음 조회시는 Space
//연속 조회시에 이전 조회한 OutBlock의 cts_time 값으로 설정
//cnt	건수	Number	Y	3	1이상 900 이하
