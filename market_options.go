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

// 기간별주가
type PeriodPriceOption struct {
	Ticker string
	//1:일, 2:주, 3:월
	DateType int
	Date     string
	Count    int
}

func (PeriodPriceOption) Path() string {
	return "/stock/market-data"
}

func (PeriodPriceOption) String() string {
	return "t1305"
}

func (c PeriodPriceOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"dwmcode": %d,
	"date": "%s",
	"cnt": %d
  }
}`, c.String(),
		c.Ticker,
		c.DateType,
		c.Date,
		c.Count,
	)), nil
}

// 주식시간대별체결조회챠트
type TimeOfDayTransactionChartOption struct {
	Ticker string
	// hhmm
	StartTime string
	// hhmm
	EndTime string
	// 분간격
	MinInterval string
}

func (TimeOfDayTransactionChartOption) Path() string {
	return "/stock/market-data"
}

func (TimeOfDayTransactionChartOption) String() string {
	return "t1308"
}

func (c TimeOfDayTransactionChartOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"starttime": "%s",
	"endtime": "%s",
	"bun_time": "%s"
  }
}`, c.String(),
		c.Ticker,
		c.StartTime,
		c.EndTime,
		c.MinInterval,
	)), nil
}

// 주식당일전일분틱조회
type DayMinuteTickOption struct {
	// 0:당일 1:전일
	DayType string
	// 0:분 1:틱
	TimeType string
	Ticker   string
	// 처음 조회시 시간 입력값. t1310OutBlock1.chetime <= endtime 인 데이터 조회됨.
	EndTime string
	//처음 조회시 Space 다음 조회시 t1310OutBlock.cts_time 값 입력
	ContinuesTime string
}

func (DayMinuteTickOption) Path() string {
	return "/stock/market-data"
}

func (DayMinuteTickOption) String() string {
	return "t1310"
}

func (c DayMinuteTickOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"daygb": "%s",
	"timegb": "%s",
	"shcode": "%s",
	"endtime": "%s",
	"cts_time": "%s"
  }
}`, c.String(),
		c.DayType,
		c.TimeType,
		c.Ticker,
		c.EndTime,
		c.ContinuesTime,
	)), nil
}

// 관리/불성실/투자유의조회
type ManagementUnfaithfulInvestmentCautionOption struct {
	// 0:전체 1:코스피 2:코스닥
	Market string
	// 1:관리 2:불성실공시 3:투자유의 4.투자환기
	Check string
	// 	처음 조회시는 Space 연속 조회시에 이전 조회한 OutBlock의 cts_shcode 값으로 설정
	ContinuesTicker string
}

func (ManagementUnfaithfulInvestmentCautionOption) Path() string {
	return "/stock/market-data"
}

func (ManagementUnfaithfulInvestmentCautionOption) String() string {
	return "t1404"
}

func (c ManagementUnfaithfulInvestmentCautionOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"jongchk": "%s",
	"cts_shcode": "%s"
  }
}`, c.String(),
		c.Market,
		c.Check,
		c.ContinuesTicker,
	)), nil
}

// 투자경고/매매정지/정리매매조회
type InvestmentWarningStopClearanceOption struct {
	// 0:전체 1:코스피 2:코스닥
	Market string
	// 1:투자경고 2:매매정지 3:정리매매 4 : 투자주의 5 : 투자위험 6 : 위험예고 7 : 단기과열지정 8 : 이상급등종목 9 : 상장주식수 부족
	Check           string
	ContinuesTicker string
}

func (InvestmentWarningStopClearanceOption) Path() string {
	return "/stock/market-data"
}

func (InvestmentWarningStopClearanceOption) String() string {
	return "t1405"
}

func (c InvestmentWarningStopClearanceOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"jongchk": "%s",
	"cts_shcode": "%s"
  }
}`, c.String(),
		c.Market,
		c.Check,
		c.ContinuesTicker,
	)), nil
}

// 초저유동성조회
type LowLiquidityOption struct {
	// 0:전체 1:코스피 2:코스닥
	Market          string
	ContinuesTicker string
}

func (LowLiquidityOption) Path() string {
	return "/stock/market-data"
}

func (LowLiquidityOption) String() string {
	return "t1410"
}

func (c LowLiquidityOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"cts_shcode": "%s"
  }
}`, c.String(),
		c.Market,
		c.ContinuesTicker,
	)), nil
}

// 상/하한
type UpperLowerLimitOption struct {
	//1:20종목씩 조회 2:전체조회
	QueryType string
	// 0:전체 1:코스피 2:코스닥
	Market string
	// 	0:당일 1:전일
	PredayType string
	// 	1:상한 4:하한
	UpDown string
	//대상제외값(설정시 저장됨) 증거금50 : 0x00400000 증거금100 : 0x00800000 증거금50/100 : 0x00200000 관리종목 : 0x00000080 시장경보 : 0x00000100 거래정지 : 0x00000200 우선주 : 0x00004000 투자유의 : 0x04000000 정리매매 : 0x01000000 불성실공시 : 0x80000000
	Exclude int
	// 현재가 >= sprice
	StartPrice int
	// 현재가 <= eprice
	EndPrice int
	// 거래량 >= volume
	Volume int
	Index  int
}

func (UpperLowerLimitOption) Path() string {
	return "/stock/market-data"
}

func (UpperLowerLimitOption) String() string {
	return "t1422"
}

func (c UpperLowerLimitOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"qrygb": "%s",
	"gubun": "%s",
	"jnilgubun": "%s",
	"sign": "%s",
	"jc_num": %d,
	"sprice": %d,
	"eprice": %d,
	"volume": %d,
	"idx": %d
  }
}`, c.String(),
		c.QueryType,
		c.Market,
		c.PredayType,
		c.UpDown,
		c.Exclude,
		c.StartPrice,
		c.EndPrice,
		c.Volume,
		c.Index,
	)), nil
}

// 상/하한가직전
type UpperLowerLimitBeforeOption struct {
	// 1:20종목씩 조회 그외:전체조회
	QueryType string
	// 0:전체 1:코스피 2:코스닥
	Market string
	// 1:상한직전 2:하한직전
	UpDown string
	// 등락율 signgubun 이 '1'(상한직전)인 경우 diff 이상 signgubun 이 '1'(상한직전)인 경우 -diff 이하
	ChangeRate int
	// 	대상제외값(설정시 저장됨) Default:000000000000 000000000128(0x00000080):관리종목 000000000256(0x00000100):시장경보 000000000512(0x00000200):거래정지 000000016384(0x00004000):우선주 000002097152(0x00200000):증거금50/100 000004194304(0x00400000):증거금50 000008388608(0x00800000):증거금100 000016777216(0x01000000):정리매매 000067108864(0x04000000):투자유의 002147483648(0x80000000):불성실공시 ex) 관리종목, 시장경보 종목 제외시 : 000000000384( 128 + 256 )
	Exclude    int
	StartPrice int
	EndPrice   int
	Volume     int
	Index      int
	// 	c' or 'C' 입력시 전일 상/하한가 제외
	PredayExclude string
}

func (UpperLowerLimitBeforeOption) Path() string {
	return "/stock/market-data"
}

func (UpperLowerLimitBeforeOption) String() string {
	return "t1427"
}

func (c UpperLowerLimitBeforeOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"qrygb": "%s",
	"gubun": "%s",
	"signgubun": "%s",
	"diff": %d,
	"jc_num": %d,
	"sprice": %d,
	"eprice": %d,
	"volume": %d,
	"idx": %d,
	"jshex": "%s"
  }
}`, c.String(),
		c.QueryType,
		c.Market,
		c.UpDown,
		c.ChangeRate,
		c.Exclude,
		c.StartPrice,
		c.EndPrice,
		c.Volume,
		c.Index,
		c.PredayExclude,
	)), nil
}

// 신고/신저가
type NewHighLowOption struct {
	// 0:전체 1:코스피 2:코스닥
	Market string
	// 신고신저 	0:신고 1:신저
	UpDown string
	// 기간 0:전일 1:5일 2:10일 3:20일 4:60일 5:90일 6:52주 7:년중
	Period string
	// 	유지여부 	0:일시돌파 1:돌파유지
	Keep string
	// 	대상제외 대상제외값(설정시 저장됨) 증거금50 : 0x00400000 증거금100 : 0x00800000 증거금50/100 : 0x00200000 관리종목 : 0x00000080 시장경보 : 0x00000100 거래정지 : 0x00000200 우선주 : 0x00004000 투자유의 : 0x04000000 정리매매 : 0x01000000 불성실공시 : 0x80000000
	Exclude    int
	StartPrice int
	EndPrice   int
	Volume     int
	Index      int
	// 	기본 => 000000000000 상장지수펀드 => 000000000001 선박투자회사 => 000000000002 스펙 => 000000000004 ETN => 000000000008(0x00000008) 투자주의 => 000000000016(0x00000010) 투자위험 => 000000000032(0x00000020) 위험예고 => 000000000064(0x00000040) 담보불가 => 000000000128(0x00000080) 두개 이상 제외시 해당 값을 합산한다.
	SecondExclude int
}

func (NewHighLowOption) Path() string {
	return "/stock/market-data"
}

func (NewHighLowOption) String() string {
	return "t1442"
}

func (c NewHighLowOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"type1": "%s",
	"type2": "%s",
	"type3": "%s",
	"jc_num": %d,
	"sprice": %d,
	"eprice": %d,
	"volume": %d,
	"idx": %d,
	"jc_num2": %d
  }
}`, c.String(),
		c.Market,
		c.UpDown,
		c.Period,
		c.Keep,
		c.Exclude,
		c.StartPrice,
		c.EndPrice,
		c.Volume,
		c.Index,
		c.SecondExclude,
	)), nil
}

// 가격대별매매비중조회
type PriceRangeTransactionRatioOption struct {
	Ticker string
	// 1:당일 2:전일 3:당일+전일
	DateType string
}

func (PriceRangeTransactionRatioOption) Path() string {
	return "/stock/market-data"
}

func (PriceRangeTransactionRatioOption) String() string {
	return "t1449"
}

func (c PriceRangeTransactionRatioOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"dategb": "%s"
  }
}`, c.String(),
		c.Ticker,
		c.DateType,
	)), nil
}

// 시간대별호가잔량추이
type TimeOfDayAskingPriceOption struct {
	Ticker string
	// 분구분
	// 	00:30초 01:1분 02:2분 03:3분 .....
	Minute string
	// 기본값 : Space, 현재시간을 기준으로 함 연속조회시에 직전 조회결과인 OutBlock의 time 값으로 설정
	Time string
	// 요청건수( 1 이상 500 이하값만 유효) ex) 10건 요청시 "010"
	Count string
}

func (TimeOfDayAskingPriceOption) Path() string {
	return "/stock/market-data"
}

func (TimeOfDayAskingPriceOption) String() string {
	return "t1471"
}

func (c TimeOfDayAskingPriceOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"gubun": "%s",
	"time": "%s",
	"cnt": "%s"
  }
}`, c.String(),
		c.Ticker,
		c.Minute,
		c.Time,
		c.Count,
	)), nil
}

// 체결강도추이
type TransactionIntensityOption struct {
	Ticker string
	// 상승하락 	시간별/일별 구분 0 : 시간별 1 : 일별
	UpDown string
	// 스페이스 입력시 최대 20개 데이터 조회됨.
	DataCount int
	//다음 조회시 입력. 이전 조회시 OutBlock.date값 입력
	Date int
	// 다음 조회시 입력. 이전 조회시 OutBlock.time값 입력
	Time int
	//일반 조회 : 0 차트 조회 : 1 OutBlock1의 volume 필드 값 구분함. 일반 : 누적거래량, 차트 : 체결량
	QueryType string
}

func (TransactionIntensityOption) Path() string {
	return "/stock/market-data"
}

func (TransactionIntensityOption) String() string {
	return "t1475"
}

func (c TransactionIntensityOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"vptype": "%s",
	"datacnt": %d,
	"date": %d,
	"time": %d,
	"gubun": "%s"
  }
}`, c.String(),
		c.Ticker,
		c.UpDown,
		c.DataCount,
		c.Date,
		c.Time,
		c.QueryType,
	)), nil
}

// 시간별예상체결가
type TimeOfDayExpectedTransactionPriceOption struct {
	Ticker        string
	ContinuesTime string
	Count         int
}

func (TimeOfDayExpectedTransactionPriceOption) Path() string {
	return "/stock/market-data"
}

func (TimeOfDayExpectedTransactionPriceOption) String() string {
	return "t1486"
}

func (c TimeOfDayExpectedTransactionPriceOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"shcode": "%s",
	"cts_time": "%s",
	"cnt": %d
  }
}`, c.String(),
		c.Ticker,
		c.ContinuesTime,
		c.Count,
	)), nil
}

// 예상체결가등락율상위조회
type ExpectedTransactionPriceChangeRateTopOption struct {
	// 0:전체 1:코스피 2:코스닥
	Market string
	// 	1:상승 2:하락
	UpDown string
	// 장구분
	// 	1:장전 2:장후 3:직전대비
	MarketType string
	// 	종목체크 	대상제외값 0x00000080:관리종목 0x00000100:시장경보 0x00000200:거래정지 0x00004000:우선주 0x00200000:증거금50/100 0x00400000:증거금50 0x00800000:증거금100 0x01000000:정리매매 0x04000000:투자유의 0x80000000:불성실공시
	TickerCheck string
	//처음 조회시는 Space 연속 조회시에 이전 조회한 OutBlock의 idx 값으로 설정
	Index int
	// 	전체:0 1만주이상:1 5만주이상:2 10만주이상:3 50만주이상:4 백만주이상:5
	Volume string
	// 예상체결시작가격 yesprice <= 예상체결가 인 종목
	StartPrice int
	// 예상체결종료가격 예상체결가 <= yeeprice 인 종목
	EndPrice int
	// 예상체결량 예상체결량 >= yevolume 인 종목
	TransactionVolume int
}

func (ExpectedTransactionPriceChangeRateTopOption) Path() string {
	return "/stock/market-data"
}

func (ExpectedTransactionPriceChangeRateTopOption) String() string {
	return "t1488"
}

func (c ExpectedTransactionPriceChangeRateTopOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"sign": "%s",
	"jgubun": "%s",
	"jongchk": "%s",
	"idx": %d,
	"volume": "%s",
	"yesprice": %d,
	"yeeprice": %d,
	"yevolume": %d
  }
}`, c.String(),
		c.Market,
		c.UpDown,
		c.MarketType,
		c.TickerCheck,
		c.Index,
		c.Volume,
		c.StartPrice,
		c.EndPrice,
		c.TransactionVolume,
	)), nil
}

// API용주식멀티현재가조회
type MultiCurrentPriceOption struct {
	// 최대 50개까지
	Count int
	// 	구분자 없이 종목코드를 붙여서 입력 078020, 000660, 005930 을 전송시 '078020000660005930' 을 입력
	Tickers string
}

func (MultiCurrentPriceOption) Path() string {
	return "/stock/market-data"
}

func (MultiCurrentPriceOption) String() string {
	return "t8407"
}

func (c MultiCurrentPriceOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"nrec": %d,
	"shcode": "%s"
  }
}`, c.String(),
		c.Count,
		c.Tickers,
	)), nil
}

// 주식마스터조회API용
type StockMasterOption struct {
	// 구분(KSP:1KSD:2)
	Type string
}

func (StockMasterOption) Path() string {
	return "/stock/market-data"
}

func (StockMasterOption) String() string {
	return "t9945"
}

func (c StockMasterOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s"
  }
}`, c.String(),
		c.Type,
	)), nil
}
