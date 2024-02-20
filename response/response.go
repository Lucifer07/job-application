package response

type ResponseMsg struct {
	Message string
	Data    any `json:"data,omitempty"`
}
type ResponseMsgErr struct {
	StatusCode int `json:"-"`
	Message    string
}
type Id struct {
	Id int `json:"id"`
}
