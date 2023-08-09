package ebest_go

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
)

type Option interface {
	String() string
	Path() string
}

type Client struct {
	appKey      string
	appSecret   string
	accessToken string
	aCache      bool
	log         *zap.Logger
	expire      time.Duration
	cli         *resty.Client
	ws          *Websocket
	broadCaster *BroadCast
	one         sync.Once
}

func NewClient(options ...ClientOption) *Client {
	client := &Client{}

	defaultLogger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	client.log = defaultLogger

	for i := range options {
		options[i](client)
	}
	cli := resty.New()
	cli.SetBaseURL("https://openapi.ebestsec.co.kr:8080")
	cli.SetDebug(true)
	client.cli = cli

	if client.aCache {
		f, err := os.Open("token")
		if err != nil && errors.Is(err, os.ErrNotExist) {
			tf, err := os.Create("token")
			if err != nil {
				panic(err)
			}
			defer tf.Close()
			tk, err := client.AccessToken(context.Background())
			if err != nil {
				panic(err)
			}

			client.accessToken = string(tk)
			if _, err := tf.Write([]byte(tk)); err != nil {
				panic(err)
			}
		} else if err != nil {
			panic(err)
		} else {
			bt, err := io.ReadAll(f)
			if err != nil {
				panic(err)
			}

			tk, _ := jwt.Parse(string(bt), nil)
			d, err := tk.Claims.GetExpirationTime()
			if err != nil {
				panic(err)
			}
			if d.Before(time.Now()) {
				tk, err := client.AccessToken(context.Background())
				if err != nil {
					fmt.Println(err)
				}
				f, err := os.Create("token")
				if err != nil {
					fmt.Println(err)
				}
				client.accessToken = tk
				f.Write([]byte(tk))
			}
			client.accessToken = string(bt)
		}
		defer f.Close()

	}

	cli.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		switch {
		case client.appKey == "":
			return errors.New("empty app key")
		case client.appSecret == "":
			return errors.New("empty app secret")
		case client.accessToken == "":
			return errors.New("empty access token")
		}

		r.SetHeader("authorization", "Bearer "+client.accessToken)
		return nil
	})

	return client
}

func (c *Client) AccessToken(ctx context.Context) (string, error) {
	res, err := c.cli.R().SetFormData(map[string]string{
		"grant_type":   "client_credentials",
		"appkey":       c.appKey,
		"appsecretkey": c.appSecret,
		"scope":        "oob",
	}).Post("/oauth2/token")
	if err != nil {
		return "", err
	}

	return gjson.GetBytes(res.Body(), "access_token").String(), nil
}

func (c *Client) Execute(ctx context.Context, option Option) ([]byte, error) {
	headers := map[string]string{
		"tr_cd":        option.String(),
		"tr_cont":      "N",
		"content-type": "application/json; charset=utf-8",
	}
	res, err := c.cli.R().
		SetContext(ctx).
		SetHeaders(headers).
		SetBody(option).
		Post(option.Path())
	if err != nil {
		return nil, err
	}

	return res.Body(), nil
}

type SubscriptionTR string
type SubscriptionTRType string

const (
	// ETF호가잔량
	ETFAskingbalance SubscriptionTR = "B7_"
	// KOSPI시간외단일가호가잔량
	OverTimeKOSPIAskikngSingleBalance SubscriptionTR = "DH1"
	// KOSDAQ시간외단일가호가잔량
	OverTimeKOSDAQAskingSignleBalance SubscriptionTR = "DHA"
	// KOSDAQ시간외단일가체결
	OvertimeKOSDAQSinglePriceContract SubscriptionTR = "DK3"
	// KOSPI시간외단일가체결
	OvertimeKOSPISinglePriceContract SubscriptionTR = "DS3"
	// 시간외단일가VI발동해제
	OvertimeSinglePriceVIClear SubscriptionTR = "DVI"
	// KOSPI호가잔량
	KOSPIAskingBalance SubscriptionTR = "H1_"
	// KOSDAQ호가잔량
	KOSDAQAskingBalance SubscriptionTR = "HA_"
	// 코스피ETF종목실시간NAV
	KOSPIETFRealtimeNAV SubscriptionTR = "I5_"
	// 지수
	Index SubscriptionTR = "IJ_"
	// KOSPI거래원
	KOSPITrader SubscriptionTR = "K1_"
	// KOSDAQ체결
	KOSDAQContract SubscriptionTR = "K3_"
	// KOSDAQ프로그램매매종목별
	KOSDAQProgramTrade SubscriptionTR = "KH_"
	// KOSDAQ프로그램매매전체집계
	KOSDAQProgramTradeTotal SubscriptionTR = "KM_"
	// KOSDAQ우선호가
	KOSDAQPriorityAsking SubscriptionTR = "KS_"
	// KOSDAQ거래원
	KOSDAQTrader SubscriptionTR = "OK_"
	// KOSPI프로그램매매종목별
	KOSPIProgramTrade SubscriptionTR = "PH_"
	// KOSPI프로그램매매전체집계
	KOSPIProgramTradeTotal SubscriptionTR = "PM_"
	// KOSPI우선호가
	KOSPIPriorityAskingSubscriptionTR = "S2_"
	// KOSPI체결
	KOSPIContract SubscriptionTR = "S3_"
	// KOSPI기세
	KOSPIMomentum SubscriptionTR = "S4_"
	// 주식주문접수
	StockOrderAcceptance SubscriptionTR = "SC0"
	// 주식주문체결
	StockOrderContract SubscriptionTR = "SC1"
	// 주식주문정정
	StockOrderModify SubscriptionTR = "SC2"
	// 주식주문취소
	StockOrderCancel SubscriptionTR = "SC3"
	// 주식주문거부
	StockOrderReject SubscriptionTR = "SC4"
	// 상/하한가근접진입
	UpperLowerLimitNearApproach SubscriptionTR = "SHC"
	// 상/하한가근접이탈
	UpperLowerLimitNearDeparture SubscriptionTR = "SHD"
	// 상/하한가진입
	UpperLowerLimitApproach SubscriptionTR = "SHI"
	// 상/하한가이탈
	UpperLowerLimitDeparture SubscriptionTR = "SHO"
	// VI발동해제
	VIClear SubscriptionTR = "VI_"
	// 예상지수
	ExpectedIndex SubscriptionTR = "YJ_"
	// KOSDAQ예상체결
	KOSDAQExpectedContract SubscriptionTR = "YK3"
	// KOSPI예상체결
	KOSPIExpectedContract SubscriptionTR = "YS3"
	// 뉴ELW투자지표민감도
	NewELWInvestmentIndicatorSensitivity SubscriptionTR = "ESN"
	// ELW장전시간외호가잔량
	ELWPreMarketAskingBalance SubscriptionTR = "h2_"
	// ELW호가잔량
	ELWAskingBalance SubscriptionTR = "h3_"
	// ELW거래원
	ELWTrader SubscriptionTR = "k1_"
	// ELW우선호가
	ELWPriorityAsking SubscriptionTR = "s2_"
	// ELW체결
	ELWContract SubscriptionTR = "s3_"
	// ELW기세
	ELWMomentum SubscriptionTR = "s4_"
	// ELW예상체결
	ELWExpectedContract SubscriptionTR = "Ys3"

	AddAccountTRType    SubscriptionTRType = "1"
	DeleteAccountTRType SubscriptionTRType = "2"
	AddPriceTRType      SubscriptionTRType = "3"
	DeletePriceTRType   SubscriptionTRType = "4"
)

type SubscriptionContent struct {
	Type   SubscriptionTRType
	TRCD   SubscriptionTR
	Ticker string
}

type ReadtimeContent struct {
	Header struct {
		Token  string `json:"token"`
		TrType string `json:"tr_type"`
	} `json:"header"`
	Body struct {
		TrCD  string `json:"tr_cd"`
		TrKey string `json:"tr_key"`
	} `json:"body"`
}

func (c *Client) Subscribe(ctx context.Context, contents ...SubscriptionContent) (<-chan []byte, error) {
	ch := make(chan []byte, 1)
	ws := &Websocket{
		Id:                           0,
		Meta:                         nil,
		Logger:                       c.log,
		Errors:                       nil,
		Reconnect:                    false,
		ReconnectIntervalMax:         0,
		ReconnectRandomizationFactor: 0,
		HandshakeTimeout:             0,
		Verbose:                      true,
		OnConnect:                    nil,
		OnDisconnect:                 nil,
		OnConnectError:               nil,
		OnDisconnectError:            nil,
		OnReadError:                  nil,
		OnWriteError:                 nil,
	}
	if err := ws.Dial("wss://openapi.ebestsec.co.kr:9443/websocket", http.Header{"content-type": []string{"application/json; charset=utf-8"}}); err != nil {
		panic(err)
	}

	for _, content := range contents {
		r := ReadtimeContent{
			Header: struct {
				Token  string `json:"token"`
				TrType string `json:"tr_type"`
			}(struct {
				Token  string
				TrType string
			}{
				Token:  c.accessToken,
				TrType: string(content.Type),
			}),
			Body: struct {
				TrCD  string `json:"tr_cd"`
				TrKey string `json:"tr_key"`
			}(struct {
				TrCD  string
				TrKey string
			}{
				TrCD:  string(content.TRCD),
				TrKey: string(content.Ticker),
			}),
		}
		if err := ws.WriteJSON(r); err != nil {
			panic(err)
		}
	}

	go func() {
		defer close(ch)
		for {
			select {
			case <-ctx.Done():
				break
			default:
				mt, m, err := ws.ReadMessage()
				if err != nil {
					break
				}
				if mt == websocket.TextMessage {
					ch <- m
				}
			}
		}
	}()
	return ch, nil
}

type Code struct {
	Code     string
	Name     string
	Industry string
}

// Search stock code
func (c *Client) Kosdaq(ctx context.Context) ([]Code, error) {
	cli := resty.New()
	res, err := cli.
		R().
		SetContext(ctx).
		SetFormData(map[string]string{
			"bld":         "dbms/MDC/STAT/standard/MDCSTAT03901",
			"locale":      "ko_KR",
			"mktId":       "KSQ",
			"trdDd":       "20230416",
			"money":       "1",
			"csvxls_isNo": "false",
		}).
		Post("http://data.krx.co.kr/comm/bldAttendant/getJsonData.cmd")
	if err != nil {
		return nil, err
	}
	var codes []Code
	re := gjson.ParseBytes(res.Body())
	for _, i := range re.Get("block1").Array() {
		codes = append(codes, Code{
			Code:     i.Get("ISU_SRT_CD").String(),
			Name:     i.Get("ISU_ABBRV").String(),
			Industry: i.Get("IDX_IND_NM").String(),
		})
	}

	return codes, nil
}

func (c *Client) Kospi(ctx context.Context) ([]Code, error) {
	cli := resty.New()
	res, err := cli.
		R().
		SetContext(ctx).
		SetFormData(map[string]string{
			"bld":         "dbms/MDC/STAT/standard/MDCSTAT03901",
			"locale":      "ko_KR",
			"mktId":       "STK",
			"trdDd":       "20230416",
			"money":       "1",
			"csvxls_isNo": "false",
		}).
		Post("http://data.krx.co.kr/comm/bldAttendant/getJsonData.cmd")
	if err != nil {
		return nil, err
	}
	var codes []Code
	re := gjson.ParseBytes(res.Body())
	for _, i := range re.Get("block1").Array() {
		codes = append(codes, Code{
			Code:     i.Get("ISU_SRT_CD").String(),
			Name:     i.Get("ISU_ABBRV").String(),
			Industry: i.Get("IDX_IND_NM").String(),
		})
	}

	return codes, nil
}
