package base_vo

import "github.com/SisyphusSQ/golib/utils/httputil"

type BaseListReq struct {
	Page     int `json:"page" query:"page"`
	PageSize int `json:"pageSize" query:"pageSize"`
}

func ValidateBaseList(page, pageSize int) error {
	if page == 0 || pageSize == 0 {
		return httputil.ErrBadParamInput
	}

	return nil
}
