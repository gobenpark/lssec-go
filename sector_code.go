package ebest_go

type SectorCode string

const (
	KOSPI_CODE  SectorCode = "001"
	KOSDAQ_CODE SectorCode = "301"
	KRX100_CODE SectorCode = "501"
	KP200_CODE  SectorCode = "101"
	SRI_CODE    SectorCode = "515"
	// 코스닥 프리미어
	KOSDAQPREMIER_CODE SectorCode = "404"
	// KRX 보험
	KRXINSURANCE_CODE SectorCode = "516"
	// 운송
	KRXTRANSPORTATION_CODE SectorCode = "517"
)

//"hname": "종       합",
//"upcode": "001"
//"hname": "대   형  주",
//"upcode": "002"
//"hname": "중   형  주",
//"upcode": "003"
//"hname": "소   형  주",
//"upcode": "004"
//"hname": "음 식 료 업",
//"upcode": "005"
//"hname": "섬 유 의 복",
//"upcode": "006"
//"hname": "종 이 목 재",
//"upcode": "007"
//"hname": "화       학",
//"upcode": "008"
//"hname": "의   약  품",
//"upcode": "009"
//"hname": "비금속 광물",
//"upcode": "010"
//"hname": "철 강 금 속",
//"upcode": "011"
//"hname": "기       계",
//"upcode": "012"
//"hname": "전 기 전 자",
//"upcode": "013"
//"hname": "의 료 정 밀",
//"upcode": "014"
//"hname": "운 수 장 비",
//"upcode": "015"
//"hname": "유   통  업",
//"upcode": "016"
//"hname": "전 기 가 스",
//"upcode": "017"
//"hname": "건   설  업",
//"upcode": "018"
//"hname": "운 수 창 고",
//"upcode": "019"
//"hname": "통   신  업",
//"upcode": "020"
//"hname": "금   융  업",
//"upcode": "021"
//"hname": "증       권",
//"upcode": "024"
//"hname": "보   험  업",
//"upcode": "025"
//"hname": "서 비 스 업",
//"upcode": "026"
//"hname": "제   조  업",
//"upcode": "027"
//"hname": "KOSPI200",
//"upcode": "101"
//"hname": "KP200 건설",
//"upcode": "111"
//"hname": "KP200 중공업",
//"upcode": "112"
//"hname": "KP200 철강소재",
//"upcode": "113"
//"hname": "KP200 에너지화학",
//"upcode": "114"
//"hname": "KP200 정보기술",
//"upcode": "115"
//"hname": "KP200 금융",
//"upcode": "116"
//"hname": "KP200 생활소비재",
//"upcode": "117"
//"hname": "KP200 경기소비재",
//"upcode": "118"
//"hname": "KP200 산업재",
//"upcode": "119"
//"hname": "KP200 건강관리",
//"upcode": "120"
//"hname": "KP200 커뮤니케이션서",
//"upcode": "121"
//"hname": "KP200 초대형주제외지",
//"upcode": "122"
//"hname": "KOSPI100",
//"upcode": "201"
//"hname": "KOSPI50",
//"upcode": "202"
//"hname": "F-KOSPI200인버스",
//"upcode": "204"
//"hname": "VKOSPI",
//"upcode": "205"
//"hname": "F-KOSPI200",
//"upcode": "206"
//"hname": "KP200 동일가중",
//"upcode": "207"
//"hname": "KP100 동일가중",
//"upcode": "208"
//"hname": "KP 50 동일가중",
//"upcode": "209"
//"hname": "KRX100동일가중",
//"upcode": "210"
//"hname": "KP200 고배당지수",
//"upcode": "211"
//"hname": "KOSPI 고배당 50",
//"upcode": "213"
//"hname": "KOSPI 배당성장 50",
//"upcode": "214"
//"hname": "KOSPI 우선주",
//"upcode": "215"
//"hname": "KP200 중소형주지수",
//"upcode": "216"
//"hname": "KP200 필수소비재",
//"upcode": "217"
//"hname": "KP200 예측 고배당 50",
//"upcode": "222"
//"hname": "KP200 예측 배당성장",
//"upcode": "223"
//"hname": "KP200 예측 TR",
//"upcode": "224"
//"hname": "KP200 예측 NTR",
//"upcode": "225"
//"hname": "코스닥 종합",
//"upcode": "301"
//"hname": "기타 서비스",
//"upcode": "303"
//"hname": "IT   종  합",
//"upcode": "304"
//"hname": "제       조",
//"upcode": "306"
//"hname": "건       설",
//"upcode": "307"
//"hname": "유       통",
//"upcode": "308"
//"hname": "운       송",
//"upcode": "310"
//"hname": "금       융",
//"upcode": "311"
//"hname": "통 신/방 송",
//"upcode": "312"
//"hname": "IT   S / W",
//"upcode": "313"
//"hname": "IT   H / W",
//"upcode": "314"
//"hname": "음식료 담배",
//"upcode": "315"
//"hname": "섬 유/의 류",
//"upcode": "316"
//"hname": "종 이/목 재",
//"upcode": "317"
//"hname": "출 판 매 체",
//"upcode": "318"
//"hname": "화       학",
//"upcode": "319"
//"hname": "제       약",
//"upcode": "320"
//"hname": "비   금  속",
//"upcode": "321"
//"hname": "금       속",
//"upcode": "322"
//"hname": "기 계 장 비",
//"upcode": "323"
//"hname": "전 기/전 자",
//"upcode": "324"
//"hname": "의 료/정 밀",
//"upcode": "325"
//"hname": "운 송/부 품",
//"upcode": "326"
//"hname": "기 타 제 조",
//"upcode": "327"
//"hname": "통신 서비스",
//"upcode": "328"
//"hname": "방송 서비스",
//"upcode": "329"
//"hname": "인   터  넷",
//"upcode": "330"
//"hname": "디   지  털",
//"upcode": "331"
//"hname": "소프트 웨어",
//"upcode": "332"
//"hname": "컴퓨터서비스",
//"upcode": "333"
//"hname": "통 신 장 비",
//"upcode": "334"
//"hname": "정 보 기 기",
//"upcode": "335"
//"hname": "반   도  체",
//"upcode": "336"
//"hname": "IT   부  품",
//"upcode": "337"
//"hname": "오락/문화",
//"upcode": "338"
//"hname": "KOSDAQ100",
//"upcode": "339"
//"hname": "KOSDAQ MID 300",
//"upcode": "340"
//"hname": "KOSDAQ SMALL",
//"upcode": "341"
//"hname": "우 량 기 업",
//"upcode": "342"
//"hname": "벤 처 기 업",
//"upcode": "343"
//"hname": "중 견 기 업",
//"upcode": "344"
//"hname": "기술성장 기업",
//"upcode": "345"
//"hname": "코스닥 150",
//"upcode": "405"
//"hname": "K R X 1 0 0",
//"upcode": "501"
//"hname": "KRX자 동 차",
//"upcode": "502"
//"hname": "KRX반 도 체",
//"upcode": "503"
//"hname": "KRX바 이 오",
//"upcode": "504"
//"hname": "KRX금    융",
//"upcode": "505"
//"hname": "KRX에너지화학",
//"upcode": "507"
//"hname": "KRX철    강",
//"upcode": "508"
//"hname": "KRX방송통신",
//"upcode": "510"
//"hname": "KRX건    설",
//"upcode": "511"
//"hname": "KRX증    권",
//"upcode": "513"
//"hname": "KRX기계장비",
//"upcode": "514"
//"hname": "KRX보    험",
//"upcode": "516"
//"hname": "KRX운    송",
//"upcode": "517"
//"hname": "KRX ESG Leaders 150",
//"upcode": "523"
//"hname": "KRX Governance Leade",
//"upcode": "524"
//"hname": "KRX Eco Leaders 100",
//"upcode": "525"
//"hname": "KRX 최소변동성지수",
//"upcode": "526"
//"hname": "KRX 경기소비재",
//"upcode": "531"
//"hname": "KRX 생활소비재",
//"upcode": "532"
//"hname": "KRX 미디어&엔터테인�",
//"upcode": "533"
//"hname": "KRX 정보기술",
//"upcode": "534"
//"hname": "KRX 유틸리티",
//"upcode": "535"
//"hname": "KRX 금현물",
//"upcode": "536"
//"hname": "KRX 금현물 USD",
//"upcode": "537"
//"hname": "KRX 300",
//"upcode": "538"
//"hname": "KRX 300 정보기술",
//"upcode": "540"
//"hname": "KRX 300 금융",
//"upcode": "541"
//"hname": "KRX 300 자유소비재",
//"upcode": "542"
//"hname": "KRX 300 산업재",
//"upcode": "543"
//"hname": "KRX 300 헬스케어",
//"upcode": "544"
//"hname": "KRX 300 커뮤니케이션",
//"upcode": "545"
//"hname": "KRX 300 소재",
//"upcode": "546"
//"hname": "KRX 300 필수소비재",
//"upcode": "547"
//"hname": "KRX BBIG K-뉴딜",
//"upcode": "548"
//"hname": "KRX 2차전지 K-뉴딜",
//"upcode": "549"
//"hname": "KRX 바이오 K-뉴딜",
//"upcode": "550"
//"hname": "KRX 인터넷 K-뉴딜",
//"upcode": "551"
//"hname": "KRX 게임 K-뉴딜",
//"upcode": "552"
//"hname": "KOSPI200 기후변화",
//"upcode": "553"
//"hname": "KRX 300 기후변화",
//"upcode": "554"
//"hname": "KRX 기후변화 솔루션",
//"upcode": "555"
//"hname": "F-USDKRW",
//"upcode": "601"
//"hname": "F-USDKRW 인버스",
//"upcode": "602"
//"hname": "KOSPI200레버리지",
//"upcode": "603"
//"hname": "KP200 커버드콜 5OTM",
//"upcode": "604"
//"hname": "KP200 프로텍티브풋 5",
//"upcode": "605"
//"hname": "국채선물지수",
//"upcode": "606"
//"hname": "국채선물인버스",
//"upcode": "607"
//"hname": "F-LKTB",
//"upcode": "608"
//"hname": "F-LKTB 인버스",
//"upcode": "609"
//"hname": "주식미국채DAE지수",
//"upcode": "615"
//"hname": "COBIX",
//"upcode": "616"
//"hname": "MOBIX",
//"upcode": "617"
//"hname": "F-USDKRW 레버리지",
//"upcode": "618"
//"hname": "F-USDKRW 인버스 -2x",
//"upcode": "619"
//"hname": "F-USDKRW 인버스 -3x",
//"upcode": "620"
//"hname": "K200 USD 선물 바이셀",
//"upcode": "622"
//"hname": "USD K200 선물 바이셀",
//"upcode": "623"
//"hname": "K200선물매수 콜매도",
//"upcode": "624"
//"hname": "K200선물매도 풋매도",
//"upcode": "625"
//"hname": "K200 가치저변동성지�",
//"upcode": "630"
//"hname": "F-KP200 인버스 -2x",
//"upcode": "638"
//"hname": "F-KP200 인버스 -3x",
//"upcode": "639"
//"hname": "KTOP 30",
//"upcode": "640"
//"hname": "KTOP 30 레버리지",
//"upcode": "641"
//"hname": "필수소비재채권혼합지",
//"upcode": "642"
//"hname": "배당성장채권혼합지수",
//"upcode": "643"
//"hname": "KP200 에너지화학 레�",
//"upcode": "644"
//"hname": "KP200 정보기술 레버�",
//"upcode": "645"
//"hname": "KP200 금융 레버리지",
//"upcode": "646"
//"hname": "KP200 경기소비재 레�",
//"upcode": "647"
//"hname": "F-KP200 레버리지",
//"upcode": "648"
//"hname": "F-KQ150",
//"upcode": "649"
//"hname": "F-KQ150 인버스",
//"upcode": "650"
//"hname": "F-KQ150 인버스 -2x",
//"upcode": "651"
//"hname": "F-KQ150 인버스 -3x",
//"upcode": "652"
//"hname": "KQ150 레버리지",
//"upcode": "653"
//"hname": "KQ150 레버리지 0.5",
//"upcode": "654"
//"hname": "WISE-KAP 스마트베타Q",
//"upcode": "655"
//"hname": "MF-KP200",
//"upcode": "656"
//"hname": "MF-KP200 레버리지",
//"upcode": "657"
//"hname": "MF-KP200 인버스",
//"upcode": "658"
//"hname": "MF-KP200 인버스 -2x",
//"upcode": "659"
//"hname": "F-KQ150 레버리지",
//"upcode": "660"
//"hname": "KP200 커버드콜 ATM",
//"upcode": "661"
//"hname": "KP200 현선물레버리지",
//"upcode": "662"
//"hname": "한국대만IT프리미어",
//"upcode": "663"
//"hname": "F-KRX300",
//"upcode": "664"
//"hname": "F-KRX300 레버리지",
//"upcode": "665"
//"hname": "F-KRX300 인버스",
//"upcode": "666"
//"hname": "F-KRX300 인버스 -2X",
//"upcode": "667"
//"hname": "KRX300 레버리지",
//"upcode": "668"
//"hname": "F-KRX300 USD 혼합",
//"upcode": "669"
//"hname": "SF-KP200 목표변동성",
//"upcode": "670"
//"hname": "SF-KQ150 USD 혼합",
//"upcode": "671"
//"hname": "F-K200 에너지/화학",
//"upcode": "751"
//"hname": "F-K200 정보기술",
//"upcode": "752"
//"hname": "F-K200 금융",
//"upcode": "753"
//"hname": "F-K200 경기소비재",
//"upcode": "754"
//"hname": "F-K200 건설",
//"upcode": "755"
//"hname": "F-K200 중공업",
//"upcode": "756"
//"hname": "F-K200 헬스케어",
//"upcode": "757"
//"hname": "F-K200 생활소비재",
//"upcode": "758"
//"hname": "F-K200 철강/소재",
//"upcode": "759"
//"hname": "F-K200 산업재",
//"upcode": "760"
//"hname": "F-K200 에너지/화학 �",
//"upcode": "761"
//"hname": "F-K200 정보기술 레버",
//"upcode": "762"
//"hname": "F-K200 금융 레버리지",
//"upcode": "763"
//"hname": "F-K200 경기소비재 레",
//"upcode": "764"
//"hname": "F-K200 건설 레버리지",
//"upcode": "765"
//"hname": "F-K200 중공업 레버리",
//"upcode": "766"
//"hname": "F-K200 헬스케어 레버",
//"upcode": "767"
//"hname": "F-K200 생활소비재 레",
//"upcode": "768"
//"hname": "F-K200 철강/소재 레�",
//"upcode": "769"
//"hname": "F-K200 산업재 레버리",
//"upcode": "770"
//"hname": "F-K200 에너지/화학 �",
//"upcode": "771"
//"hname": "F-K200 정보기술 인버",
//"upcode": "772"
//"hname": "F-K200 금융 인버스",
//"upcode": "773"
//"hname": "F-K200 경기소비재 인",
//"upcode": "774"
//"hname": "F-K200 건설 인버스",
//"upcode": "775"
//"hname": "F-K200 중공업 인버스",
//"upcode": "776"
//"hname": "F-K200 헬스케어 인버",
//"upcode": "777"
//"hname": "F-K200 생활소비재 인",
//"upcode": "778"
//"hname": "F-K200 철강/소재 인�",
//"upcode": "779"
//"hname": "F-K200 산업재 인버스",
//"upcode": "780"
//"hname": "F-K200 에너지/화학 �",
//"upcode": "781"
//"hname": "F-K200 정보기술 인버",
//"upcode": "782"
//"hname": "F-K200 금융 인버스-2",
//"upcode": "783"
//"hname": "F-K200 경기소비재 인",
//"upcode": "784"
//"hname": "F-K200 건설 인버스-2",
//"upcode": "785"
//"hname": "F-K200 중공업 인버스",
//"upcode": "786"
//"hname": "F-K200 헬스케어 인버",
//"upcode": "787"
//"hname": "F-K200 생활소비재 인",
//"upcode": "788"
//"hname": "F-K200 철강/소재 인�",
//"upcode": "789"
//"hname": "F-K200 산업재 인버스",
//"upcode": "790"
//"hname": "KOSPI 총수익",
//"upcode": "791"
//"hname": "KOSDAQ 총수익",
//"upcode": "792"
//"hname": "KOSDAQ150 총수익",
//"upcode": "793"
//"hname": "KRX 300 총수익",
//"upcode": "794"
//"hname": "KOSPI 고배당50 총수�",
//"upcode": "796"
//"hname": "KSP배당성장50총수익",
//"upcode": "797"
//"hname": "KOSPI200 ESG",
//"upcode": "798"
//"hname": "F-KTB 3-10스티프닝",
//"upcode": "810"
//"hname": "F-KTB 3-10플래트닝",
//"upcode": "811"
//"hname": "F-KTB 3-10스티프닝2X",
//"upcode": "812"
//"hname": "F-KTB 3-10플래트닝2X",
//"upcode": "813"
//"hname": "KP200-KQ150 비중5:5",
//"upcode": "814"
//"hname": "KP200-KQ150 비중7:3",
//"upcode": "815"
//"hname": "KP200-KQ150 비중3:7",
//"upcode": "816"
//"hname": "KP200 L KQ150 S",
//"upcode": "817"
//"hname": "KQ150 L KP200 S",
//"upcode": "818"
//"hname": "KP200 L KQ150 0.5 S",
//"upcode": "819"
//"hname": "KQ150 L KP200 0.5 S",
//"upcode": "820"
