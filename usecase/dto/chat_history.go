package dto

type ChatHistoryResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
}
