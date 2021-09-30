package entity

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type DefaultResponse struct {
	Data interface{} `json:"data"`
	Response
}
