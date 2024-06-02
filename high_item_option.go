package lssec_go

import "fmt"

// 등락율상위
type UpDownRateTopOption struct {
	// 	0:전체 1:코스피 2:코스닥
	Market string
	// 0: 상승률 1: 하락률 2: 보합
	UpDown string
	// 	0: 당일 1: 전일
	DayOption string
	// 대상제외
	// 대상제외값 증거금50 : 0x00400000 증거금100 : 0x00800000 증거금50/100 : 0x00200000 관리종목 : 0x00000080 시장경보 : 0x00000100 거래정지 : 0x00000200 우선주 : 0x00004000 투자유의 : 0x04000000 정리매매 : 0x01000000 불성실공시 : 0x80000000
	ExcludeOption int
	// 현재가 >= StartPrice
	StartPrice int
	// 현재가 <= EndPrice
	EndPrice int
	Volume   int
	Index    int
	// 기본 => 000000000000 상장지수펀드 => 000000000001 선박투자회사 => 000000000002 스펙 => 000000000004 ETN => 000000000008(0x00000008) 투자주의 => 000000000016(0x00000010) 투자위험 => 000000000032(0x00000020) 위험예고 => 000000000064(0x00000040) 담보불가 => 000000000128(0x00000080) 두개 이상 제외시 해당 값을 합산한다.
	ExcludeOption2 int
}

func (UpDownRateTopOption) String() string {
	return "t1441"
}

func (UpDownRateTopOption) Path() string {
	return "/stock/high-item"
}

func (t UpDownRateTopOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun1": "%s",
	"gubun2": "%s",
	"gubun3": "%s",
	"jc_num": %d,
	"sprice": %d,
	"eprice": %d,
	"volume": %d,
	"idx": %d,
	"jc_num2": %d
  }
}`, t.String(),
		t.Market,
		t.UpDown,
		t.DayOption,
		t.ExcludeOption,
		t.StartPrice,
		t.EndPrice,
		t.Volume,
		t.Index,
		t.ExcludeOption2,
	)), nil
}

// 시가총액상위
type MarketCapTopOption struct {
	SectorCode
	Index int
}

func (MarketCapTopOption) String() string {
	return "t1444"
}

func (MarketCapTopOption) Path() string {
	return "/stock/high-item"
}

func (t MarketCapTopOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"upcode": "%s",
	"idx": %d
  }
}`, t.String(),
		t.SectorCode,
		t.Index,
	)), nil
}

// 거래량상위
type VolumeTopOption struct {
	//	0:전체 1:코스피 2:코스닥
	Market string
	// 1:당일 2:전일
	DayOption string
	// 시작등락율
	StartUpDownRate int
	// 종료등락율
	EndUpDownRate int
	// 대상제외
	// 대상제외값 (0x00000080)관리종목 => 000000000128 (0x00000100)시장경보 => 000000000256 (0x00000200)거래정지 => 000000000512 (0x00004000)우선주 => 000000016384 (0x00200000)증거금50 => 000008388608 (0x01000000)정리매매 => 000016777216 (0x04000000)투자유의 => 000067108864 (0x80000000)불성실공시 => -02147483648 두개 이상 제외시 해당 값을 합산한다 예)관리종목 + 시장경보 = 000000000128 + 000000000256 = 000000000384
	ExcludeOption int
	StartPrice    int
	EndPrice      int
	Volume        int
	Index         int
}

func (VolumeTopOption) String() string {
	return "t1452"
}

func (VolumeTopOption) Path() string {
	return "/stock/high-item"
}

func (t VolumeTopOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"jnilgubun": "%s",
	"sdiff": %d,
	"ediff": %d,
	"jc_num": %d,
	"sprice": %d,
	"eprice": %d,
	"volume": %d,
	"idx": %d
  }
}`, t.String(),
		t.Market,
		t.DayOption,
		t.StartUpDownRate,
		t.EndUpDownRate,
		t.ExcludeOption,
		t.StartPrice,
		t.EndPrice,
		t.Volume,
		t.Index,
	)), nil
}

// 거래대금상위
type AmountTopOption struct {
	//	0:전체 1:코스피 2:코스닥
	Market string
	// 0:당일 1:전일
	DayOption string
	// 대상제외
	// 	대상제외값 (0x00000080)관리종목 => 000000000128 (0x00000100)시장경보 => 000000000256 (0x00000200)거래정지 => 000000000512 (0x00004000)우선주 => 000000016384 (0x00200000)증거금50 => 000008388608 (0x01000000)정리매매 => 000016777216 (0x04000000)투자유의 => 000067108864 (0x80000000)불성실공시 => -02147483648 두개 이상 제외시 해당 값을 합산한다 예)관리종목 + 시장경보 = 000000000128 + 000000000256 = 000000000384
	ExcludeOption int
	StartPrice    int
	EndPrice      int
	Volume        int
	Index         int
	// 기본 => 000000000000 상장지수펀드 => 000000000001 선박투자회사 => 000000000002 스펙 => 000000000004 ETN => 000000000008(0x00000008) 투자주의 => 000000000016(0x00000010) 투자위험 => 000000000032(0x00000020) 위험예고 => 000000000064(0x00000040) 담보불가 => 000000000128(0x00000080) 두개 이상 제외시 해당 값을 합산한다.
	ExcludeOption2 int
}

func (AmountTopOption) String() string {
	return "t1463"
}

func (AmountTopOption) Path() string {
	return "/stock/high-item"
}

func (t AmountTopOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"jnilgubun": "%s",
	"jc_num": %d,
	"sprice": %d,
	"eprice": %d,
	"volume": %d,
	"idx": %d,
	"jc_num2": %d
  }
}`, t.String(),
		t.Market,
		t.DayOption,
		t.ExcludeOption,
		t.StartPrice,
		t.EndPrice,
		t.Volume,
		t.Index,
		t.ExcludeOption2,
	)), nil
}

// 전일동시간대비거래급증
type VolumeIncreaseOption struct {
	//	0:전체 1:코스피 2:코스닥
	Market string
	// 전일 거래량
	// 0:1주 이상 1:1만주 이상 2:5만주 이상 3:10만주 이상 4:20만주 이상 5:50만주 이상 6:100만주 이상
	PreDayVolume string
	// 거래급등율
	// 0:전체 1:2000%이하 2:1500%이하 3:1000%이하 4:500%이하 5:100%이하 6:50%이하
	TradeRateTopRate string
	// 	대상제외값 (0x00000080)관리종목 => 000000000128 (0x00000100)시장경보 => 000000000256 (0x00000200)거래정지 => 000000000512 (0x00004000)우선주 => 000000016384 (0x00200000)증거금50 => 000008388608 (0x01000000)정리매매 => 000016777216 (0x04000000)투자유의 => 000067108864 (0x80000000)불성실공시 => -02147483648 두개 이상 제외시 해당 값을 합산한다 예)관리종목 + 시장경보 = 000000000128 + 000000000256 = 000000000384
	ExcludeOption int
	StartPrice    int
	EndPrice      int
	Volume        int
	Index         int
	// 기본 => 000000000000 상장지수펀드 => 000000000001 선박투자회사 => 000000000002 스펙 => 000000000004 ETN => 000000000008(0x00000008) 투자주의 => 000000000016(0x00000010) 투자위험 => 000000000032(0x00000020) 위험예고 => 000000000064(0x00000040) 담보불가 => 000000000128(0x00000080) 두개 이상 제외시 해당 값을 합산한다.
	ExcludeOption2 int
}

func (VolumeIncreaseOption) String() string {
	return "t1466"
}

func (VolumeIncreaseOption) Path() string {
	return "/stock/high-item"
}

func (t VolumeIncreaseOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"type1": "%s",
	"type2": "%s",
	"jc_num": %d,
	"sprice": %d,
	"eprice": %d,
	"volume": %d,
	"idx": %d,
	"jc_num2": %d
  }
}`, t.String(),
		t.Market,
		t.PreDayVolume,
		t.TradeRateTopRate,
		t.ExcludeOption,
		t.StartPrice,
		t.EndPrice,
		t.Volume,
		t.Index,
		t.ExcludeOption2,
	)), nil
}

// 시간외등락율상위
type AfterMarketPriceTopOption struct {
	//	0:전체 1:코스피 2:코스닥
	Market string
	// 	0: 상승률 1: 하락률
	UpDownRate string
	// 종목체크
	// 	0: 전체 1: 우선제외 2: 관리제외 3: 우선관리제외
	TickerCheck string
	// 	0: 전체거래량 1: 1천주 이상 2: 5천주 이상 3: 1만주 이상 4: 5만주 이상 5: 10만주 이상 6: 50만주 이상 7: 100만주 이상
	Volume string
	Index  int
}

func (AfterMarketPriceTopOption) String() string {
	return "t1481"
}

func (AfterMarketPriceTopOption) Path() string {
	return "/stock/high-item"
}

func (t AfterMarketPriceTopOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"gubun2": "%s",
	"jongchk": "%s",
	"volume": "%s",
	"idx": %d
  }
}`, t.String(),
		t.Market,
		t.UpDownRate,
		t.TickerCheck,
		t.Volume,
		t.Index,
	)), nil
}

// 시간외거래량상위
type AfterMarketVolumeTopOption struct {
	//	0:전체 1:코스피 2:코스닥
	Market string
	// 	0: 전체 1: 우선제외 2: 관리제외 3: 우선관리제외
	Volume string
	Index  int
}

func (AfterMarketVolumeTopOption) String() string {
	return "t1482"
}

func (AfterMarketVolumeTopOption) Path() string {
	return "/stock/high-item"
}

func (t AfterMarketVolumeTopOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"jongchk": "%s",
	"idx": %d
  }
}`, t.String(),
		t.Market,
		t.Volume,
		t.Index,
	)), nil
}

// 예상체결량상위조회
type ExpectVolumeTopOption struct {
	//	0:전체 1:코스피 2:코스닥
	Market string
	// 장구분
	// 	0:장전 1:장후
	OnOffTime string
	// 대상제외값(설정시 저장됨) 증거금50 : 0x00400000 증거금100 : 0x00800000 증거금50/100 : 0x00200000 관리종목 : 0x00000080 시장경보 : 0x00000100 거래정지 : 0x00000200 우선주 : 0x00004000 투자유의 : 0x04000000 정리매매 : 0x01000000 불성실공시 : 0x80000000
	TickerCheck string
	Index       int
	// ExpectStartPrice <= 예상체결가 인 종목
	ExpectStartPrice int
	// ExpectEndPrice >= 예상체결가 인 종목
	ExpectEndPrice int
	// ExpectVolume <= 예상체결량 인 종목
	ExpectVolume int
}

func (ExpectVolumeTopOption) String() string {
	return "t1489"
}

func (ExpectVolumeTopOption) Path() string {
	return "/stock/high-item"
}

func (t ExpectVolumeTopOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun": "%s",
	"jgubun": "%s",
	"jongchk": "%s",
	"idx": %d,
	"yesprice": %d,
	"yeeprice": %d,
	"yevolume": %d
  }
}`, t.String(),
		t.Market,
		t.OnOffTime,
		t.TickerCheck,
		t.Index,
		t.ExpectStartPrice,
		t.ExpectEndPrice,
		t.ExpectVolume,
	)), nil
}

// 단일가예상등락율상위
type ExpectUpDownRateTopOption struct {
	//	0:전체 1:코스피 2:코스닥
	Market string
	// 	0: 상승률 1: 하락률
	UpDownRate string
	// 종목체크
	// 	0: 전체 1: 우선제외 2: 관리제외 3: 우선관리제외
	TickerCheck string
	// 	0: 전체거래량 1: 1천주 이상 2: 5천주 이상 3: 1만주 이상 4: 5만주 이상 5: 10만주 이상 6: 50만주 이상 7: 100만주 이상
	Volume string
	Index  int
}

func (ExpectUpDownRateTopOption) String() string {
	return "t1492"
}

func (ExpectUpDownRateTopOption) Path() string {
	return "/stock/high-item"
}

func (t ExpectUpDownRateTopOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{
  "%sInBlock": {
	"gubun1": "%s",
	"gubun2": "%s",
	"jongchk": "%s",
	"volume": "%s",
	"idx": %d
  }
}`, t.String(),
		t.Market,
		t.UpDownRate,
		t.TickerCheck,
		t.Volume,
		t.Index,
	)), nil
}
