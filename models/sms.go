package models

type SendSmsRequestToPM struct {
	Messages []*Message `json:"messages"`
}

type Message struct {
	Recipient string `json:"recipient"`
	MessageID string `json:"message-id"`
	Sms       *Sms   `json:"sms"`
}

type Sms struct {
	Originator string   `json:"originator"`
	Content    *Content `json:"content"`
}

type Content struct {
	Text string `json:"text"`
}
