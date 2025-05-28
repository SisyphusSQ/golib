package lark_dto

type Content struct {
	Tag    string `json:"tag,omitempty"`
	Text   string `json:"text,omitempty"`
	Href   string `json:"href,omitempty"`
	UserId string `json:"user_id,omitempty"`
}

type ZnCn struct {
	Title   string      `json:"title"`
	Content [][]Content `json:"content"`
}

type Co struct {
	Post Post `json:"post"`
}

type Post struct {
	ZhCN ZnCn `json:"zh_cn"`
}

type BotMsg struct {
	MsgType string `json:"msg_type"`
	Content Co     `json:"content"`
}

func NewBotMsg(title string, contents []Content) BotMsg {
	return BotMsg{
		MsgType: "post",
		Content: Co{
			Post{
				ZhCN: ZnCn{
					Title:   title,
					Content: [][]Content{contents},
				},
			},
		},
	}
}

type BotMsgResp struct {
	Code Code   `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type Code int

const (
	Success Code = 0
	Failure Code = 9499
)
