package models

type Feedback struct {
	Id        int    `json:"id"`
	UserId    int    `json:"user_id"`
	Sentiment string `json:"sentiment"`
	Message   string `json:"message"`
}
