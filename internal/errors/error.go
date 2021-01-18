package errors

type ErrorCode int

func (err ErrorCode) ExternalErrString() string {
	// Improve by DB
	switch err {
	case Success:
		return "Thành công"
	case ListTasksFailed:
		return "Không thể lấy danh sách công việc"
	case AddTaskFailed:
		return "Không thể thêm công việc mới"
	case LoginFailed:
		return "Không thể đăng nhập thành công"
	}
	return "Unknown"
}

const (
	Success ErrorCode = iota
	LoginFailed
	AddTaskFailed
	ListTasksFailed
)

type ErrorApplication struct {
	Code        ErrorCode `json:"code"`
	InternalMsg string    `json:"internal_msg"`
	ExternalMsg string    `json:"external_msg"`
}

func (e ErrorApplication) Error() string {
	return e.InternalMsg
}

func NewErrorMsg(errcode ErrorCode, internalMsg, externalMsg string) error {
	return &ErrorApplication{
		Code:        errcode,
		InternalMsg: internalMsg,
		ExternalMsg: externalMsg,
	}
}
