package lark_dto

import larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"

type LarkMsgReq struct {
	Contacts []string `json:"contacts"`
	Message  string   `json:"message"`
}

type LarkMsgResp struct {
	Resp []*larkim.CreateMessageResp `json:"resp"`
}
