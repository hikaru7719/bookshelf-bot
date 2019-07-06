package service

import "testing"

func TestBookService_SendAnswer(t *testing.T) {
	service, _ := NewService()
	service.SendAnswer("Amazon", "bot-test")
}
