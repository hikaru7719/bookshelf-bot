package message

import "testing"

func TestSendMessage(t *testing.T) {
	SendMessage("bot-test", "test test")
}
