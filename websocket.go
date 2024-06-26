package lssec_go

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var (
	ErrNotConnected          = errors.New("websocket not connected")
	ErrUrlEmpty              = errors.New("url can not be empty")
	ErrUrlWrongScheme        = errors.New("websocket uri must start with ws or wss scheme")
	ErrUrlNamePassNotAllowed = errors.New("user name and password are not allowed in websocket uri")
	//ErrCantConnect           = errors.New("websocket can't connect")
)

type WsOpts func(dl *websocket.Dialer)

type Websocket struct {
	// Websocket ID
	Id uint64
	// Websocket Meta
	Meta map[string]interface{}

	//Logger *logrus.Logger
	Logger *zap.Logger

	Errors chan<- error

	Reconnect bool

	// default to 2 seconds
	// default to 30 seconds
	ReconnectIntervalMax time.Duration
	// interval, default to 1.5
	ReconnectRandomizationFactor float64
	// default to 2 seconds
	HandshakeTimeout time.Duration
	// Verbose suppress connecting/reconnecting messages.
	Verbose bool

	// Cal function
	OnConnect    func(ws *Websocket)
	OnDisconnect func(ws *Websocket)

	OnConnectError    func(ws *Websocket, err error)
	OnDisconnectError func(ws *Websocket, err error)
	OnReadError       func(ws *Websocket, err error)
	OnWriteError      func(ws *Websocket, err error)

	dialer        *websocket.Dialer
	url           string
	requestHeader http.Header
	httpResponse  *http.Response
	mu            sync.Mutex
	dialErr       error
	isConnected   bool
	isClosed      bool

	*websocket.Conn
}

func (ws *Websocket) WriteJSON(v interface{}) error {
	err := ErrNotConnected
	if ws.IsConnected() {
		ws.mu.Lock()
		err = ws.Conn.WriteJSON(v)
		ws.mu.Unlock()
		if err != nil {
			if ws.OnWriteError != nil {
				ws.OnWriteError(ws, err)
			}
			ws.closeAndReconnect()
		}
	}

	return err
}

func (ws *Websocket) WriteMessage(messageType int, data []byte) error {
	err := ErrNotConnected
	if ws.IsConnected() {
		err = ws.Conn.WriteMessage(messageType, data)
		if err != nil {
			if ws.OnWriteError != nil {
				ws.OnWriteError(ws, err)
			}
			ws.closeAndReconnect()
		}
	}

	return err
}

func (ws *Websocket) ReadMessage() (messageType int, message []byte, err error) {
	err = ErrNotConnected
	if ws.IsConnected() {
		messageType, message, err = ws.Conn.ReadMessage()
		if err != nil {
			if ws.OnReadError != nil {
				ws.OnReadError(ws, err)
			}
			ws.closeAndReconnect()
		}
	}

	return
}

func (ws *Websocket) Close() {
	ws.mu.Lock()
	if ws.Conn != nil {
		err := ws.Conn.Close()
		if err == nil && ws.isConnected && ws.OnDisconnect != nil {
			ws.OnDisconnect(ws)
		}
		if err != nil && ws.OnDisconnectError != nil {
			ws.OnDisconnectError(ws, err)
		}
	}
	//ws.isClosed = true
	ws.isConnected = false
	ws.mu.Unlock()
}

func (ws *Websocket) closeAndReconnect() {
	ws.Close()
	ws.Connect()
}

func (ws *Websocket) Dial(urlStr string, reqHeader http.Header, opts ...WsOpts) error {
	_, err := parseUrl(urlStr)
	if err != nil {
		return err
	}

	ws.url = urlStr
	ws.setDefaults()

	ws.dialer = &websocket.Dialer{
		Proxy:            http.ProxyFromEnvironment,
		HandshakeTimeout: ws.HandshakeTimeout,
	}
	for _, opt := range opts {
		opt(ws.dialer)
	}

	hs := ws.HandshakeTimeout
	go ws.Connect()

	// wait on first attempt
	time.Sleep(hs)

	return nil
}

func (ws *Websocket) Connect() error {

	expb := backoff.NewExponentialBackOff()
	expb.MaxInterval = ws.ReconnectIntervalMax
	expb.RandomizationFactor = ws.ReconnectRandomizationFactor
	expb.MaxElapsedTime = time.Minute

	return backoff.Retry(func() error {
		wsConn, httpResp, err := ws.dialer.Dial(ws.url, ws.requestHeader)

		ws.mu.Lock()
		ws.Conn = wsConn
		ws.dialErr = err
		ws.isConnected = err == nil
		ws.httpResponse = httpResp
		ws.mu.Unlock()

		if err == nil {
			if ws.Verbose && ws.Logger != nil {
				ws.Logger.Info(fmt.Sprintf("Websocket[%d].Dial: connection was successfully established with %s", ws.Id, ws.url))
			}
			if ws.OnConnect != nil {
				ws.OnConnect(ws)
			}
			return nil
		} else {
			if ws.Verbose && ws.Logger != nil {
				ws.Logger.Error(fmt.Sprintf("Websocket[%d].Dial: can't connect to %s, will try again in %v", ws.Id, ws.url, expb.NextBackOff()))
			}
			if ws.OnConnectError != nil {
				ws.OnConnectError(ws, err)
			}
			return err
		}
	}, expb)
}

func (ws *Websocket) GetHTTPResponse() *http.Response {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	return ws.httpResponse
}

func (ws *Websocket) GetDialError() error {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	return ws.dialErr
}

func (ws *Websocket) IsConnected() bool {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	return ws.isConnected
}

func (ws *Websocket) setDefaults() {

	if ws.ReconnectIntervalMax == 0 {
		ws.ReconnectIntervalMax = 60 * time.Second
	}

	if ws.ReconnectRandomizationFactor == 0 {
		ws.ReconnectRandomizationFactor = 0.5
	}

	if ws.HandshakeTimeout == 0 {
		ws.HandshakeTimeout = 2 * time.Second
	}
}

func parseUrl(urlStr string) (*url.URL, error) {
	if urlStr == "" {
		return nil, ErrUrlEmpty
	}
	u, err := url.Parse(urlStr)

	if err != nil {
		return nil, err
	}

	if u.Scheme != "ws" && u.Scheme != "wss" {
		return nil, ErrUrlWrongScheme
	}

	if u.User != nil {
		return nil, ErrUrlNamePassNotAllowed
	}

	return u, nil
}
