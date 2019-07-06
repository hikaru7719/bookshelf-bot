package handler

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/hikaru7719/bookshelf-bot/service"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	byteBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(byteBody, &jsonMap); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(jsonMap)

	token := jsonMap["token"].(string)
	if token != os.Getenv("SLACK_TOKEN") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	eventType := jsonMap["type"].(string)
	switch eventType {
	case "url_verification":
		challenge := jsonMap["challenge"].(string)
		w.WriteHeader(200)
		w.Write([]byte(challenge))
		return
	case "event_callback":
		event := jsonMap["event"].(map[string]interface{})
		eventTypeString := event["type"].(string)
		if eventTypeString == "app_mention" {
			eventText := event["text"].(string)
			stringReader := strings.NewReader(eventText)
			scanner := bufio.NewScanner(stringReader)
			scanner.Scan()
			scanner.Scan()
			text := scanner.Text()
			splitSlice := strings.Split(text, ":")
			if len(splitSlice) < 1 {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			switch splitSlice[0] {
			case "search":
				w.WriteHeader(200)
				if err != nil {
					fmt.Fprint(os.Stderr, err)
					return
				}
				service, err := service.NewService()
				if err != nil {
					fmt.Fprint(os.Stderr, err)
					return
				}
				channelName := event["channel"].(string)
				go service.SendAnswer(splitSlice[1], channelName)
				return
			}
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
