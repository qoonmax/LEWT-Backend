package translateService

import (
	text "LEWT_Backend"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
)

const endPoint = "https://translate.api.cloud.yandex.net/translate/v2/translate"

var serviceAccountId string
var apiKey string

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading env variables")
	}

	serviceAccountId = os.Getenv("YANDEX_SERVICE_ACCOUNT_ID")
	apiKey = os.Getenv("YANDEX_API_KEY")
}

type yandexRequestData struct {
	ServiceAccountId   string `json:"serviceAccountId"`
	TargetLanguageCode string `json:"targetLanguageCode"`
	Texts              string `json:"texts"`
}

func Translate(data *text.Data) {

	if isNotEmptyString(data.GetInputText()) == false || isRussianString(data.GetInputText()) == false {
		return
	}

	if isWhiteString(data.GetInputText()) == false {
		data.ResultText = "Censored 🚫"
		data.ResetInputText()
		return
	}

	requestData := yandexRequestData{
		ServiceAccountId:   serviceAccountId,
		TargetLanguageCode: "en",
		Texts:              data.GetInputText(),
	}

	requestJsonData, err := json.Marshal(requestData)
	if err != nil {
		fmt.Println("Ошибка при преобразовании в JSON:", err)
		return
	}

	request, err := http.NewRequest("POST", endPoint, bytes.NewBuffer(requestJsonData))
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Api-Key "+apiKey)

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return
	}

	var yandexResponseData struct {
		Translations []struct {
			Text string `json:"text"`
		} `json:"translations"`
	}

	err = json.Unmarshal(responseData, &yandexResponseData)
	if err != nil {
		fmt.Println("Ошибка при разборе ответа:", err)
		return
	}
	if len(yandexResponseData.Translations) > 0 {
		data.ResultText = yandexResponseData.Translations[0].Text
	} else {
		fmt.Println("Ответ сервера: пустой")
	}
	return
}
