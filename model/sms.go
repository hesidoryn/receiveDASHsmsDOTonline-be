package model

// Sms contains information about text message
type Sms struct {
	From   string `json:"numberFrom"`
	To     string `json:"numberTo"`
	Body   string `json:"body"`
	SentAt int64  `json:"sentAt"`
}
