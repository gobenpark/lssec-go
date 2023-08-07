package ebest_go

type EbestError struct {
	Code    string `json:"rsp_cd"`
	Message string `json:"rsp_msg"`
}

func (e EbestError) Error() string {
	return e.Message
}

var (
	ErrInvalidTRCD error = EbestError{Code: "IGW00215", Message: "유효하지 않은 TR CD 입니다."}
)
