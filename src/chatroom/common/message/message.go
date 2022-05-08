package message

const (
	LoginMsgType    = "LoginMsg"
	LoginResMsgType = "loginResMsg"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
type RegisterMsg struct {
	User User `json:"user"`
}

type RegisterResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
