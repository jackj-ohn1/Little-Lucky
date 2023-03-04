package handler

type Identification struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

type ErrorInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
