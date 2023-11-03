package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// start main
func main() {
	botToken := os.Getenv("SECURITY_TOKEN")
	//https: //api.telegram.org/bot<token>/METHOD_NAME.
	botApi := "https://api.telegram.org/"
	botUrl := botApi + botToken
	//возрат бесконечного цикла for
	for {
		update, err := getUpdates(botUrl)
		if err != nil {
			log.Println("Errors", err)
		}
		fmt.Println(update)
	}
}

// request update
func getUpdates(botUrl string) ([]Update, error) {
	resp, err := http.Get(botUrl + "/getUpdates")
	if err != nil {
		return nil, err
	}
	//закрываем тело ответа после обработки
	defer resp.Body.Close()
	//io функция читает содержимое файла и возвращает его в виде среза байтов.
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
func respond(botUrl string, update Update) {
	var botMessage BotMessage
	botMessage.ChatId = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text
	buf, err := json.Marshal(botMessage)
	err = json.Unmarshal(body, &restResponce)
	if err != nil {
		return err
	}
	//contentType это хедер в http и его тип (например json)
	err := http.Post(botUrl+"/sendMessage", "application/json", &buf)
	if err != nil {
		return err
	}
	return nil
}
