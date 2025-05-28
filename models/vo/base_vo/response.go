package base_vo

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/SisyphusSQ/golib/utils/httputil"
)

type Resp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func SuccessResp(data any) *Resp {
	return &Resp{
		Code:    httputil.SUCCESS,
		Message: httputil.GetMsg(httputil.SUCCESS),
		Data:    data,
	}
}

func ErrorResp(err error) *Resp {
	return &Resp{
		Code:    httputil.InternalERROR,
		Message: err.Error(),
		Data:    nil,
	}
}

func AssertErrResp(err string) *Resp {
	return &Resp{
		Code:    httputil.InternalERROR,
		Message: err,
		Data:    nil,
	}
}

func CustomResp(code int, msg string, data any) *Resp {
	return &Resp{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

type ResponseError struct {
	Message string `json:"message"`
}

type SimpleResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func CommErrResp(c echo.Context, err error) error {
	return c.JSON(httputil.GetStatusCode(err), ErrorResp(err))
}

func CommSuccResp(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, SuccessResp(data))
}
