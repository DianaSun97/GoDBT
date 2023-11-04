package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// start main
func main() {
	botToken := os.Getenv("SECURITY_TOKEN")
	//https: //api.telegram.org/bot<token>/METHOD_NAME.
	botApi := "https://api.telegram.org/" + botToken
	//возрат бесконечного цикла for
	for {
		updates, err := getUpdates(botApi)
		if err != nil {
			log.Println("Error:", err)
		}

		for _, update := range updates {
			err := respond(botApi, update)
			if err != nil {
				log.Println("Error responding:", err)
			}
		}
		// Ждем 1 секунду перед следующим запросом
		time.Sleep(1 * time.Second)
	}
}

// request update
func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
	if err != nil {
		return nil, err
	}
	//close body after updates
	defer resp.Body.Close()
	//io функция читает содержимое файла и возвращает его в виде массив байтов.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var restResponce RestResponse
	//указатель на структуру &??
	err = json.Unmarshal(body, &restResponce)
	if err != nil {
		return nil, err
	}

	//возращает массив обьекта
	return restResponce.Result, nil
}

// request
func respond(botUrl string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = "Received: " + update.Message.Text

	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	_, err = http.Post(botUrl+"/sendMessage", "application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}

	return nil
}
