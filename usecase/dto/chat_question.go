package dto

type ChatQuestionRequest struct {
	Question string `json:"question"`
}

type ChatQuestionResponse struct {
	Answer string `json:"answer"`
}
