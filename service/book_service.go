package service

import (
	"fmt"
	"github.com/hikaru7719/bookshelf-bot/finder"
	"github.com/hikaru7719/bookshelf-bot/message"
	"strconv"
	"strings"
)

type BookService struct {
	finder finder.Finder
}

func NewService() (*BookService, error) {
	CSVFinder, err := finder.NewCSV()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &BookService{finder: CSVFinder}, nil
}

func (b *BookService) SendAnswer(query string, channelName string) {
	bookSlice, err := b.finder.Find(query)
	if err != nil {
		fmt.Println(err)
	}

	var sendMessage strings.Builder
	length := strconv.Itoa(len(bookSlice))
	fmt.Println(length)
	if len(bookSlice) > 0 {
		sendMessage.WriteString("本あったよ:sunglasses:\n")
		sendMessage.WriteString("検索結果/" + length + "件\n")
	} else {
		sendMessage.WriteString("残念だ...\n")
	}

	for _, book := range bookSlice {
		sendMessage.WriteString("```")
		sendMessage.WriteString(book.ToString())
		sendMessage.WriteString("```")
		sendMessage.WriteString("\n")
	}
	message.SendMessage(channelName, sendMessage.String())
	b.finder.Close()
}
