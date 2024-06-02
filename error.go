package lssec_go

type LssecError struct {
	Code    string `json:"rsp_cd"`
	Message string `json:"rsp_msg"`
}

func (e LssecError) Error() string {
	return e.Message
}

var (
	ErrInvalidTRCD error = LssecError{Code: "IGW00215", Message: "유효하지 않은 TR CD 입니다."}
)
