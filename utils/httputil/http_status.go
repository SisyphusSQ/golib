package httputil

import (
	"errors"
	"net/http"
)

const (
	SUCCESS               = 200
	UpdatePasswordSuccess = 201
	NotExistInentifier    = 202
	InternalERROR         = 500
	InvalidParams         = 400
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Not Found")
	ErrConflict            = errors.New("Your Item already exist")
	ErrBadParamInput       = errors.New("Param Invalid")
)

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch {
	case errors.Is(err, ErrInternalServerError):
		return http.StatusInternalServerError
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

var MsgFlags = map[int]string{
	SUCCESS:               "success",
	UpdatePasswordSuccess: "修改密码成功",
	NotExistInentifier:    "该第三方账号未绑定",
	InternalERROR:         "failed",
	InvalidParams:         "请求参数错误",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[InternalERROR]
}
