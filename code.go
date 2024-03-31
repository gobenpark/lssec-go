package ebest_go

type TRCD string

const (
	//가격유형을 잘못 입력하셨습니다.
	InvalidPriceType TRCD = "01018"
	//가입대상(%%)만 가능합니다.
	OnlyForMembers TRCD = "01021"
	//개별 계좌번호를 입력하십시오.
	EnterIndividualAccountNumber TRCD = "01033"
	//개인 신용한도 초과 금액입니다.
	ExceedPersonalCreditLimit TRCD = "01036"
	//거래구분을 입력하십시오.
	EnterTradingType TRCD = "01060"
	//거래구분을 잘못 입력하셨습니다.
	InvalidTradingType TRCD = "01061"
	//거래량을 잘못 입력하셨습니다.
	InvalidTradingVolume TRCD = "01063"
	//거래번호를 잘못 입력하셨습니다.
	InvalidTradingNumber TRCD = "01064"
	//거래불가 종목입니다.
	NonTradableStock TRCD = "01065"
	//거래유형코드를 확인하십시오.
	CheckTradingTypeCode TRCD = "01067"
	//거래일자를 입력하십시오.
	EnterTradingDate TRCD = "01068"
	//거래정지 종목으로 주문이 불가능합니다.
	TradingHaltStock TRCD = "01069"
	//거래정지 종목이므로 인터넷상에서는 출고가 불가능합니다.
	TradingHaltStockNoInternet TRCD = "01070"
	//거래중지된 계좌입니다.
	TradingHaltAccount TRCD = "01072"
	//거부된 주문으로 정정 또는 취소 주문이 불가능합니다.
	RejectedOrder TRCD = "01081"
	//계좌구분을 입력하십시오.
	EnterAccountType TRCD = "01137"
	//계좌구분을 잘못 입력하셨습니다.
	InvalidAccountType TRCD = "01138"
	//계좌명을 입력하십시오.
	EnterAccountName TRCD = "01139"
	//계좌명을 잘못 입력하셨습니다.
	InvalidAccountName TRCD = "01140"
	//계좌번호가 중복됩니다.
	DuplicateAccountNumber TRCD = "01141"
	//계좌번호를 입력하십시오.
	EnterAccountNumber TRCD = "01142"
	//계좌번호를 잘못 입력하셨습니다.
	InvalidAccountNumber TRCD = "01143"
	//계좌번호를 확인 하십시오 . 폐쇄계좌가 아닙니다.
	CheckAccountNumber TRCD = "01144"
	//계좌번호를 확인하십시오. 위탁계좌가 아닙니다.
	CheckAccountNumber2 TRCD = "01145"
	//계좌번호를 확인하십시오. 이미 등록된 계좌입니다.
	CheckAccountNumber3 TRCD = "01146"
	//계좌번호를 확인하십시오. 해당 지점의 계좌번호가 아닙니다.
	CheckAccountNumber4 TRCD = "01147"
	//계좌번호의 지점번호를 잘못 입력하셨습니다.
	InvalidBranchNumber TRCD = "01148"
	//계좌별 1회 주문수량한도를 초과했습니다.
	ExceedOrderLimit TRCD = "01149"
	//계좌별 1회 주문한도를 초과했습니다.
	ExceedOrderLimit2 TRCD = "01150"
	//계좌별 순매도 한도를 초과했습니다.
	ExceedNetSellLimit TRCD = "01151"
	//계좌별 순매수 한도를 초과했습니다.
	ExceedNetBuyLimit TRCD = "01152"
	//계좌별 일일 주문한도를 초과하셨습니다.
	ExceedDailyOrderLimit TRCD = "01153"
	//계좌별 총 매매한도를 초과했습니다.
	ExceedTotalTradingLimit TRCD = "01154"
	//귀 계좌는 신용신규약관 미 등록계좌입니다.
	NotRegisteredCreditAccount TRCD = "01253"
	//귀 계좌는 장내, 코스닥 등록 주식만 매매가 가능합니다.
	OnlyTradeRegisteredStock TRCD = "01255"
	//귀 계좌는 주식거래만 가능합니다.
	OnlyTradeStock TRCD = "01256"
	//귀 계좌는 채권거래만 가능합니다.
	OnlyTradeBond TRCD = "01257"
	//귀 계좌는 통폐합된 계좌입니다.
	MergedAccount TRCD = "01258"
	//귀 계좌번호와 주문내역의 계좌번호가 일치하지 않습니다.
	AccountNumberMismatch TRCD = "01259"
	//귀 계좌에서는 거래할 수 없는 유가증권입니다.
	NonTradableStock2 TRCD = "01261"
	//귀 계좌의 금일 주문증거금이 매수가능금액을 초과합니다.
	ExceedOrderMargin TRCD = "01264"
	//담보금부족으로 신규 신용주문이 불가능합니다.
	InsufficientCollateral TRCD = "01389"
	//담보기간을 확인하십시오.
	CheckCollateralPeriod TRCD = "01390"
	//담보대출 유지비율 미만 상태로 되므로 처리 불가합니다.
	InsufficientCollateral2 TRCD = "01391"
	//담보대출 잔고가 평가금액을 초과하여 매수주문이 불가합니다.
	ExceedCollateralBalance TRCD = "01393"
	//담보대출 잔고가 평가금액을 초과하였습니다.(매수주문불가)
	ExceedCollateralBalance2 TRCD = "01394"
	//담보부족이거나 신용주문이 있는 계좌의 거래는 취소할 수 없습니다.
	CancelNonTradableAccount TRCD = "01396"
	//담보설정 해지 후 처리하십시오.
	CancelAfterCollateral TRCD = "01397"
	//담보유지비율 미만입니다.
	InsufficientCollateral3 TRCD = "01398"
	//등록되지 않은 계좌번호입니다.
	UnregisteredAccountNumber TRCD = "01511"
	//매도 결제일자 구분을 잘못 입력하셨습니다.
	InvalidSettlementDate TRCD = "01574"
	//매도 체결일을 입력하십시오.
	EnterSellSettlementDate TRCD = "01575"
	//매도구분을 잘못 입력하였습니다.
	InvalidSellType TRCD = "01576"
	//매도금액을 잘못 입력하셨습니다.
	InvalidSellAmount TRCD = "01577"
	//매도대금 담보대출 약정계좌가 아닙니다.
	NotCollateralAccount TRCD = "01578"
	//매도대금 담보대출 약정해지(취소) 후 가능합니다.
	CancelAfterSell TRCD = "01579"
	//매도수구분을 입력하십시오.
	EnterSellBuyType TRCD = "01580"
	//매도수구분을 잘못 입력하셨습니다.
	InvalidSellBuyType TRCD = "01581"
	//매도수량을 잘못 입력하셨습니다.
	InvalidSellVolume TRCD = "01582"
	//매도의뢰자 분류코드를 잘못 입력하셨습니다.
	InvalidSellRequesterCode TRCD = "01583"
	//매도잔고가 부족합니다.
	InsufficientSellBalance TRCD = "01584"
	//매도좌수가 잔고좌수를 초과합니다.
	ExceedSellBalance TRCD = "01585"
	//매도주문이 불가한 계좌입니다.
	NonTradableAccount TRCD = "01586"
	//매도출금은 수익증권만 가능합니다.
	OnlyProfitWithdrawal TRCD = "01587"
	//매도출금인 경우만 종목번호를 입력하십시오.
	EnterStockNumber TRCD = "01588"
	//매매가 불가능한 종목입니다.
	NonTradableStock3 TRCD = "01589"
	//매매가 제한된 계좌입니다.
	RestrictedAccount TRCD = "01590"
	//매매구분을 잘못 입력하셨습니다.
	InvalidTradingType2 TRCD = "01591"
	//매매금액을 입력하시오.
	EnterTradingAmount TRCD = "01592"
	//매매단가를 잘못 입력하셨습니다.
	InvalidTradingPrice TRCD = "01593"
	//매매보고서 통보지를 입력하십시오
	EnterTradingReport TRCD = "01594"
	//매매보고서 통보지를 잘못 입력하셨습니다.
	InvalidTradingReport TRCD = "01595"
	//매매수량을 잘못 입력하셨습니다.
	InvalidTradingVolume2 TRCD = "01596"
	//매매일을 잘못 입력하셨습니다.
	InvalidTradingDate TRCD = "01597"
	//매매조건을 잘못 입력하셨습니다.
	InvalidTradingCondition TRCD = "01601"
	//매수 계좌번호를 확인하십시오.
	CheckBuyAccountNumber TRCD = "01602"
	//매수 희망가를 입력하십시오.
	EnterBuyPrice TRCD = "01603"
	//매수가 불가한 종목입니다.
	NonTradableStock4 TRCD = "01604"
	//매수금액을 잘못 입력하셨습니다.
	InvalidBuyAmount TRCD = "01605"
	//매수수량을 잘못 입력하셨습니다.
	InvalidBuyVolume TRCD = "01606"
	//매수주문이 불가한 계좌입니다.
	NonTradableAccount2 TRCD = "01612"
	//매입자금대출 가능 종목이 아닙니다.
	NotCollateralStock TRCD = "01615"
	//매입자금대출 약정 등록 계좌가 아닙니다.
	NotCollateralAccount2 TRCD = "01616"
	//사고등록된 계좌임으로 처리가 불가능 합니다.
	NonTradableAccount3 TRCD = "01832"
	//사고등록된 계좌입니다.
	RegisteredAccount TRCD = "01833"
	//사고등록된 유가증권으로 주문처리가 불가능합니다.
	NonTradableStock5 TRCD = "01835"
	//사고신고 계좌입니다.
	ReportedAccount TRCD = "01839"
	//사고종목이므로 매매가 불가능합니다.
	NonTradableStock6 TRCD = "01840"
	//사고해지 사유코드를 확인하십시오.
	CheckCancelReason TRCD = "01841"
	//상장 주식수의 5%이상은 주문이 불가합니다.
	NonTradableStock7 TRCD = "01888"
	//상장구분을 잘못 입력하셨습니다.
	InvalidListingType TRCD = "01889"
	//상장종목이 아니므로 주문이 불가합니다.
	NonListingStock TRCD = "01890"
	//상장주식수 보다  주문수량이 큽니다.
	ExceedListingStock TRCD = "01891"
	//상장주식수의 1/3 이상은 주문이 불가합니다.
	NonTradableStock8 TRCD = "01892"
	//상장폐지 종목입니다.
	DelistedStock TRCD = "01893"
	//선물계좌가 없습니다. 확인하십시오.
	NoFuturesAccount TRCD = "01920"
	//선물계좌번호를 입력하십시오.
	EnterFuturesAccountNumber TRCD = "01922"
	//선물대용 수량이 일치하지 않습니다
	NonMatchingFutures TRCD = "01923"
	//선물대용 지정수량 존재합니다
	ExistingFutures TRCD = "01925"
	//선물수량이 부족합니다.
	InsufficientFutures TRCD = "01928"
	//수량을 10주 단위로 입력하십시오.
	EnterVolume10 TRCD = "01982"
	//수량을 잘못 입력하셨습니다.
	InvalidVolume TRCD = "01983"
	//시장가 또는 조건부 지정가 주문이 불가한 종목입니다.
	NonTradableStock9 TRCD = "02030"
	//시장가주문은 정정주문이 불가합니다.
	NonMarketOrder TRCD = "02031"
	//시장구분을 입력하십시오.
	EnterMarketType TRCD = "02032"
	//시장구분을 잘못 입력하셨습니다.
	InvalidMarketType TRCD = "02033"
	//신구주 병합대상 종목이므로 매수주문이 불가능합니다.
	NonTradableStock10 TRCD = "02037"
	//신구주 병합대상 종목이므로 신규융자 주문이 불가합니다.
	NonTradableStock11 TRCD = "02038"
	//신구주 합병대상 종목이므로 신규대주 주문이 불가합니다.
	NonTradableStock12 TRCD = "02039"
	//신규상장,관리지정,해제종목의 시가결정 전으로 정정,취소주문이 불가합니다.
	NonTradableStock13 TRCD = "02044"
	//신용평가기관 채무불이행계좌는 신규 신용주문이 불가합니다.
	NonTradableAccount4 TRCD = "02045"
	//신용 불량계좌는 처리가 불가합니다.
	NonTradableAccount5 TRCD = "02046"
	//신용 수량을 입력하시오.
	EnterCreditVolume TRCD = "02047"
	//신용거래 대주종목이 아닙니다.
	NonCreditStock TRCD = "02054"
	//신용거래 융자종목이 아닙니다.
	NonCreditStock2 TRCD = "02055"
	//신용거래 종목이 아닙니다.
	NonCreditStock3 TRCD = "02056"
	//신용계좌 입니다.
	CreditAccount TRCD = "02061"
	//신용계좌 입니다.
	CreditAccount2 TRCD = "02062"
	//신용계좌이므로 처리가 불가능합니다.
	NonTradableAccount6 TRCD = "02063"
	//신용구분을 입력하십시오.
	EnterCreditType TRCD = "02066"
	//신용구분을 잘못 입력하셨습니다.
	InvalidCreditType TRCD = "02067"
	//신용금액을 입력하시오.
	EnterCreditAmount TRCD = "02068"
	//신용금액을 잘못 입력하셨습니다. 10원단위로 입력하시오.
	InvalidCreditAmount TRCD = "02069"
	//신용단가내역 정리가 불가합니다.
	NonTradableCredit TRCD = "02070"
	//신용담보 부족계좌 및 신용이자미납 계좌이므로 만기일 정정 불가합니다.
	NonTradableAccount7 TRCD = "02071"
	//신용담보 부족계좌가 없습니다.
	NoCollateralAccount TRCD = "02072"
	//신용담보 부족계좌이므로 만기일 정정 불가합니다.
	NonTradableAccount8 TRCD = "02073"
	//신용담보 재계산이 불가합니다.
	NonTradableCollateral TRCD = "02074"
	//신용담보금 부족으로 말소처리가 불가능 합니다.
	NonTradableCollateral2 TRCD = "02075"
	//신용담보금 부족으로 출금처리가 불가능 합니다.
	NonTradableCollateral3 TRCD = "02076"
	//신용담보금이 부족한 계좌이므로 처리가 불가능 합니다.
	NonTradableCollateral4 TRCD = "02077"
	//신용등급을 잘못 입력하셨습니다.
	InvalidCreditRating TRCD = "02078"
	//신용주문이 불가능한 종목입니다.
	NonTradableStock14 TRCD = "02098"
	//실명 미등록 계좌입니다.
	UnregisteredRealNameAccount TRCD = "02123"
	//예탁담보 대출 담보설정 주식 매도 재사용이 불가합니다.
	NonTradableCollateral5 TRCD = "02211"
	//예탁담보 대출 약정계좌입니다. 약정 해지 후 처리하십시오.
	CancelAfterCollateral2 TRCD = "02212"
	//예탁증권 담보대출 약정등록 후 처리 가능합니다.
	NonTradableCollateral6 TRCD = "02213"
	//예탁증권 담보대출 약정한도금액과 동일한 한도금액을 입력하십시오.
	EnterCollateralLimit TRCD = "02214"
	//원주문가와 정정주문가가 동일하여 정정주문이 불가능합니다.
	NonTradableOrder TRCD = "02257"
	//원주문번호를 잘못 입력하셨습니다.
	InvalidOrderNumber TRCD = "02258"
	//원주문번호에 해당하는 주문내용이 없습니다. 확인바랍니다.
	NoOrder TRCD = "02259"
	//원주문번호와 거부 시 원주문번호가 다릅니다.
	OrderNumberMismatch TRCD = "02260"
	//원주문번호와 정정 시 원주문번호와 다릅니다.
	OrderNumberMismatch2 TRCD = "02261"
	//원주문번호와 취소 시 원주문번호가 다릅니다.
	OrderNumberMismatch3 TRCD = "02262"
	//원주문의 종목번호와 정정, 취소 시 종목번호가 다릅니다.
	StockNumberMismatch TRCD = "02263"
	//원주문이 거부된 주문이므로 한도 해지할 수 있습니다.
	CancelAfterRejectedOrder TRCD = "02264"
	//원주문이 취소주문으로 정정,취소 주문이 불가능합니다.
	NonCancelOrder TRCD = "02265"
	//위탁계좌가 아닙니다.
	NonEntrustAccount TRCD = "02278"
	//위탁계좌번호를 입력하십시오.
	EnterEntrustAccountNumber TRCD = "02279"
	//유가미수 계좌는 신용신규 주문이 불가합니다.
	NonTradableAccount9 TRCD = "02289"
	//유가미수 계좌이므로 유가미수종목 매수만 가능합니다.
	OnlyBuyUnsettled TRCD = "02290"
	//유가미수 발생계좌이므로 한도설정 또는 취소가 불가능합니다.
	NonTradableAccount10 TRCD = "02291"
	//유가미수 변제가 불가합니다.
	NonTradableUnsettled TRCD = "02292"
	//유가미수변제를 환원할 수 없습니다.
	NonTradableUnsettled2 TRCD = "02293"
	//유가증권 미수계좌입니다.
	Unsettled TRCD = "02294"
	//유가증권 미수처리 완료계좌입니다.
	Unsettled2 TRCD = "02295"
	//유가증권 잔고가 남아 있습니다.
	ExistingUnsettled TRCD = "02296"
	//유가증권 잔고가 없습니다.
	NoUnsettled TRCD = "02297"
	//유가증권 종류를 잘못 입력하셨습니다.
	InvalidUnsettled TRCD = "02298"
	//유가증권의 종류를 잘못 입력하셨습니다.
	InvalidUnsettled2 TRCD = "02299"
	//유상청약 금액(외화)을 입력 하십시오.
	EnterSubscriptionAmount TRCD = "02300"
	//유상청약 금액(원화)을 입력 하십시오.
	EnterSubscriptionAmount2 TRCD = "02301"
	//유상청약 이외에는 금액을 입력할 수 없습니다.
	NonSubscriptionAmount TRCD = "02302"
	//유통대주 주문은 처리가 불가합니다.
	NonTradableStock15 TRCD = "02303"
	//유통분 만기일정정은 최장 150일까지만 가능합니다. 150일 이내로 입력하시오.
	EnterMaturityDate TRCD = "02304"
	//유통융자 지점한도가 초과 되었습니다.
	ExceedBranchLimit TRCD = "02305"
	//이미 취소된 주문내역 입니다.
	CancelledOrder TRCD = "02453"
	//정정 가능한 수량을 초과하였습니다.
	ExceedCorrectableVolume TRCD = "02662"
	//정정가격을 잘못 입력하셨습니다.
	InvalidCorrectablePrice TRCD = "02663"
	//정정구분을 잘못 입력하셨습니다.
	InvalidCorrectableType TRCD = "02664"
	//정정처리가 불가능합니다. 취소 후 등록으로 처리하십시오.
	NonCorrectable TRCD = "02665"
	//정정할 가격을 입력하십시오.
	EnterCorrectablePrice TRCD = "02666"
	//정정할 수량을 입력하십시오.
	EnterCorrectableVolume TRCD = "02667"
	//현금/현물 상환내역이 없습니다.
	NoCashRepayment TRCD = "03074"
	//현금/현물 상환할 대출내역이 없습니다.
	NoCashRepayment2 TRCD = "03075"
	//현금미수 계좌는 신용신규 주문이 불가합니다.
	NonTradableAccount11 TRCD = "03076"
	//현금미수 계좌이므로 매도만 가능합니다.
	OnlySell TRCD = "03077"
	//현금미수 또는 신용이자 미납 계좌이므로 현금상환 처리가 불가합니다.
	NonTradableAccount12 TRCD = "03078"
	//현금미수 또는 유가증권 미수계좌입니다.
	Unsettled3 TRCD = "03079"
	//사고등록된 고객입니다.
	ReportedCustomer TRCD = "03481"
	//장종료매매는 시장조성회원만 가능합니다.
	OnlyMarketMaker TRCD = "03561"
	//장종료 매매시 상품은 소액채권 매도가 불가합니다.
	NonTradableStock16 TRCD = "03562"
	//정규매매장이 종료되었습니다. 시간외 호가구분을 입력하십시요.
	EnterAfterHours TRCD = "03563"
	//자기주식 매매가 종료되었습니다.
	NonTradableStock17 TRCD = "03564"
	//보유중인 미결제약정 수량이 부족합니다.
	Insufficient TRCD = "03565"
	//증거금 증가 청산주문이거나 추가증거금 발생계좌이므로 주문 불가합니다.
	NonTradableAccount13 TRCD = "03568"
	//주문수량이 1회한도 주문수량을 초과하였습니다.
	ExceedOrderLimit3 TRCD = "03570"
	//선물주문한도를 초과하였습니다.1001계약 미만으로 입력 요망
	ExceedFuturesOrderLimit TRCD = "03571"
	//옵션주문한도를 초과하였습니다. 5001계약 미만으로 입력 요망
	ExceedOptionOrderLimit TRCD = "03572"
	//옵션매수나 청산주문만 가능한 계좌입니다.
	OnlyOptionBuyOrSell TRCD = "03573"
	//해당 유가증권의 잔고를 확인하세요(보유수량/매수일자 등)
	CheckStockBalance TRCD = "03574"
	//증거금 한도 초과로 주문 불가합니다.
	ExceedMarginLimit TRCD = "03576"
	//해당종목의 체결 종료전 이므로 취소할 수 없습니다.
	NonTradableStock18 TRCD = "03588"
	//장개시 전이므로 주문 불가합니다.
	NonTradableStock19 TRCD = "03589"
	//장마감 후이므로 주문 불가합니다.
	NonTradableStock20 TRCD = "03590"
	//입력일자를 잘못 입력하셨습니다
	InvalidInputDate TRCD = "03591"
	//%% 확인하여 주십시오.
	PercentCheck TRCD = "03592"
	//%% 선택하여 주십시오.
	PercentCheck2 TRCD = "03593"
	//계좌비밀번호가 틀립니다.
	InvalidPassword TRCD = "03688"
	//계좌비밀번호를 잘못 입력하셨습니다. %%
	InvalidPassword2 TRCD = "03689"
	//계좌 종목신용한도 초과입니다.
	ExceedCreditLimit TRCD = "04055"
	//고객 신용한도 초과입니다.
	ExceedCustomerCreditLimit TRCD = "04056"
	//고객 신용유형별(자기/유통) 신용한도 초과입니다.
	ExceedCustomerCreditLimit2 TRCD = "04057"
	//고객 신용유형별(자기/유통) 종목신용한도 초과입니다.
	ExceedCustomerCreditLimit3 TRCD = "04058"
	//계좌 신용유형별(자기/유통) 신용한도 초과입니다.
	ExceedCreditLimit2 TRCD = "04059"
	//계좌 신용유형별(자기/유통) 종목 신용한도 초과입니다.
	ExceedCreditLimit3 TRCD = "04060"
	//계좌 신용유형별(자기/유통) 종목군 신용한도 초과입니다.
	ExceedCreditLimit4 TRCD = "04061"
	//증거금이 평가액의 배수를 초과해 주문 불가합니다.
	ExceedMarginMultiple TRCD = "04062"
	//파생상품 ETF 종목은 위험 고지 등록 후 매매 가능 합니다. %%
	NonTradableStock21 TRCD = "04386"
	//프로그램호가 신고구분코드를 잘못 입력하였습니다.
	InvalidProgram TRCD = "04387"
	//호가적합성 오류(장마감 동시호가시 조건부지정가 불가)
	NonTradableStock22 TRCD = "04389"
	//호가적합성 오류(동시호가시 최우선/최유리/IOC/FOK 불가)
	NonTradableStock23 TRCD = "04390"
	//호가적합성 오류(사전/사후 프로그램 호가시 시장가/조건부지정가 불가)
	NonTradableStock24 TRCD = "04391"
	//호가적합성 오류(시장가 주문이 불가한 종목)
	NonTradableStock25 TRCD = "04392"
	//호가적합성 오류(조건부지정가 주문이 불가한 종목)
	NonTradableStock26 TRCD = "04393"
	//호가적합성 오류(최유리지정가 주문이 불가한 종목)
	NonTradableStock27 TRCD = "04394"
	//호가적합성 오류(최우선지정가 주문이 불가한 종목)
	NonTradableStock28 TRCD = "04395"
	//투자주의 종목은 5일간 신규융자주문이 불가합니다.
	NonTradableStock29 TRCD = "04396"
)
