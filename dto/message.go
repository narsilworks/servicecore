package dto

type Recipient struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}

type Message struct {
	To      []Recipient `json:"to,omitempty"`
	Cc      []Recipient `json:"cc,omitempty"`
	Bcc     []Recipient `json:"bcc,omitempty"`
	From    Recipient   `json:"from,omitempty"`
	ReplyTo []Recipient `json:"reply_to,omitempty"`
	Subject string      `json:"subject,omitempty"`
	Type    string      `json:"type,omitempty"`
	Format  string      `json:"format,omitempty"`
	body    string
}

func (m *Message) SetBody(t string) {
	m.body = t
}

func (m *Message) Body() string {
	return m.body
}
