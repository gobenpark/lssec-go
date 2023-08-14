<p align="center">
<h1 align="center">ebest-go</h1>
<p align="center"><a href="https://ebestsec.co.kr/">ebest investment</a> Unofficial client library for Go </p>


## Installation

```bash
go get github.com/gobenpark/ebest-go
```


## Usage

> Return value is byte array so recommand to use [gjson](https://github.com/tidwall/gjson)

```go
cli := ebest_go.NewClient(
	WithAuth(key,secret),
	WithAutomaticTokenCache(true),
	)

// Do something
cli.AccessToken(context.TODO())

bt, err := cli.Execute(context.TODO(), ebest_go.AccountBalanceOption{
PriceType:       "",
ContractType:    "",
SinglePriceType: "",
Charge:          "",
Ticker:          "",
})
if err != nil {
	return
}
```

### API List

- #### Sector
  - IndustryTrendsOverTimeOption [업종 기간별 추이]
  - TotalIndustryOption [전체업종]
  - ExpectedStockIndexOption [예상지수]
  - CurrentPriceOfIndustryOption [업종 현재가]
  - PriceOfIndustryOption [업종별 종목시세]

- #### Account
  - AccountTradeHistoryOption [계좌 거래내역]
  - AccountOrderHistoryOption [현물계좌 주문체결내역 조회(API)]
  - AccountOrderAvailableOption [현물계좌증거금률별주문가능수량조회]
  - AccountProfitDetailOption [주식계좌 기간별수익률 상세]
  - AccountDailyTradeOption [주식당일매매일지/수수료]
  - AccountBalanceOption [주식잔고2]
  - AccountContractOption [주식체결/미체결]

  미구현
  - 계좌별신용한도조회
  - 현물계좌예수금 주문가능금액 총평가 조회
  - BEP단가조회
  - 현물계좌예수금 주문가능금액 총평가2
  - 주식당일매매일지/수수료(전일)

- #### Market
  - UpDownRateTopOption [등락율상위]
  - MarketCapTopOption [시가총액상]
  - VolumeTopOption [거래량상위]
  - AmountTopOption [거래대금상위]
  - VolumeIncreaseOption [전일동시간대비거래급증]
  - AfterMarketPriceTopOption [시간외등락율상위]
  - AfterMarketVolumeTopOption [시간외거래량상위]
  - ExpectVolumeTopOption [예상체결량상위]
  - ExpectUpDownRateTopOption [단일가예상등락율상위]

- #### MarketData
  - CurrentAskingPriceOption [현재가 호가 조회]
  - CurrentPriceOption [현재가]
  - CurrentPriceMemoOption [주식현재가시세메모]
  - SearchPivotDemarkOption [주식피봇/디마크조회]
  - OvertimeTransactionCountOption [시간외체결량]
  - TimeOfDayTransactionOption [주식시간대별체결조회]
  - MinuteOfDayPriceOption [주식분별주가조회]
  - PeriodPriceOption [기간별주가p]
  - TimeOfDayTransactionChartOtion [주식시간대별체결조회챠트]
  - DayMinuteTickOption [주식당일전일분틱조회]
  - ManagementUnfaithfulInvestmentCautionOption [관리/불성실/투자유의조회]
  - InvestmentWarningStopClearanceOption [투자경고/매매정지/정리매매조회] 
  - LowLiquidityOption [초저유동성조회] 
  - UpperLowerLimitOption [상/하한]
  - UpperLowerLimitBeforeOption [상/하한가직전] 
  - NewHighLowOption [신고/신저가]
  - PriceRangeTransactionRatioOption [가격대별매매비중조회] 
  - TimeOfDayAskingPriceOption [시간대별호가잔량추이]
  - TransactionIntensityOption [체결강도추이] 
  - TimeOfDayExpectedTransactionPriceOption [시간별예상체결가]
  - ExpectedTransactionPriceChangeRateTopOption [예상체결가등락율상위조회] 
  - MultiCurrentPriceOption [API용주식멀티현재가조회] 
  - StockMasterOption [주식마스터조회API용] 

- #### Exchange
  - TopMemberCompanyOption [종목별 상위 거래원]
  - MemberShipTrandsByStockOption [종목별 회원사 추이]
  - MemberShipListOption [회원사 리스트] 

- #### Program
  - ProgramTradeSummaryOption [프로그램매매종합조회]
  - ProgramTradeTimeSeriesOption [시간대별프로그램매매추이]
  - ProgramTradeTrendOption [종목별프로그램매매동향]
  - ProgramTradeTimeSeriesByPeriodOption [기간별프로그램매매추이] 
  - ProgramTradeTrendByStockOption [종목별프로그램매매추이]
  - ProgramTradeSummaryMiniOption [프로그램매매종합조회(미니)]
  - ProgramTradeTimeSeriesChartOption [시간대별프로그램매매추이(차트)]


- #### Investor
  - TotalInvestorOption [투자자별종합]
  - InvestorTradeTimeSeriesOption [시간대별투자자매매추이]
  - InvestorTradeTimeSeriesDetailOption [시간대별투자자매매추이상세]
  - InvestorTradeSummaryOption1 [투자자매매종합1]
  - InvestorTradeSummaryOption2 [투자자매매종합2]
  - InvestorTradeSectorTimeSeriesOption [업종별분별투자자매매동향(챠트용)]
  - InvestorTradeSummaryChartOption [투자자매매종합(챠트)]

- #### Foreigner&Agency
  - ForeignerAgencyOption1 [외인기관종목별 동향]

  미구현
  - 외인기관종목별동향 [t1716,t1717]

- #### Thema
  - ThemeStockOption [테마별종목]
  - StockThemeOption [종목별테마]
  - SpecialThemeOption [특이테마]
  - ThemeStockQuoteOption [테마종목별시세조회]
  - ThemeStockQuoteOption [전체테마]

- #### HighItem
  - UpDownRateTopOption [등락율상위]
  - MarketCapTopOption [시가총액상위]
  - VolumeTopOption [거래량상위]
  - AmountTopOption [거래대금상위]
  - VolumeIncreaseOption [전일동시간대비거래급증]
  - AfterMarketPriceTopOption [시간외등락율상위]
  - AfterMarketVolumeTopOption [시간외거래량상위]
  - ExpectVolumeTopOption [예상체결량상위조회]
  - ExpectUpDownRateTopOption [단일가예상등락율상위]

- #### Chart
  - TickSectorChartOption [업종차트(틱/n틱)]
  - InvestorTrendChartOption [기간별투자자매매추이(챠트)]
  - StockChartOption [API전용주식챠트(일주월년)]
  - TickChartOption [주식챠트(틱/n틱)]
  - MinChartOption [주식챠트(N분)]
  - TotalChartOption [업종챠트(종합)]
  - DaySectorChartOption [업종챠트(일주월)]
  - MinSectorChartOption [업종챠트(N분)] 

- #### ETC
  - NewTickerOption [신규상장종목조회]
  - CollateralLoanableStockOption [예탁담보융자가능종목현황조회]
  - MarginRateByTickerOption [증거금율별종목조회]
  - StockRemainderOption [종목별잔량/사전공시]
  - CreditTradeTrendOption [신용거래동향]

  미구현
  - 종목별신용정보
  - 공매도일별추이
  - 종목별대차거래일간추이
  - 주식종목조회
  - 주식종목조회 API용

- #### Order
  - OrderOption [주식주문]
  - OrderCancelOption [주식주문취소]
  - OrderModifyOption [주식주문정정]

- ### InvestInformation


## TODO

- 연속조회

