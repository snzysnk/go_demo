package api

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
