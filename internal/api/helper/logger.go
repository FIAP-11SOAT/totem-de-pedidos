package helper

type LoggerInfo struct {
	Scope   string `json:"category"`
	Message string `json:"message"`
}

type HttpResponse struct {
	Message string `json:"message"`
}
