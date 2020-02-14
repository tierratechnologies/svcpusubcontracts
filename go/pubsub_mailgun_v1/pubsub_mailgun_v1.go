package pubsub_mailgun_v1

var Topic = "send_email_v1"

type SendEmailMsgV1 struct {
	Topic string `json:"topic"`

	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Text    string `json:"message"`
}