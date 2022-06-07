package queue

type Message struct {
	key   string
	value string
}

func NewMessage(key, value string) *Message {
	return &Message{key, value}
}
