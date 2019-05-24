package trello

import (
	"encoding/json"
	//"fmt"
	trello "github.com/adlio/trello"
	result "github.com/heaptracetechnology/microservice-trello/result"
	"net/http"
	"os"
)

//TrelloArgs struct
type TrelloArgs struct {
	BoardID     string `json:"board_id,omitempty"`
	ListID      string `json:"list_id,omitempty"`
	CardID      string `json:"card_id,omitempty"`
	CardName    string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

//GetCards trello
func GetCards(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	var trelloCards []*trello.Card
	client := trello.NewClient(apiKey, token)
	board, err := client.GetBoard(param.BoardID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	lists, err := board.GetLists(trello.Defaults())

	for _, list := range lists {
		cards, err := list.GetCards(trello.Defaults())
		if err != nil {
			result.WriteErrorResponse(responseWriter, err)
			return
		}

		trelloCards = append(trelloCards, cards...)
	}
	bytes, _ := json.Marshal(trelloCards)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}

//AddCard trello
func AddCard(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var card TrelloArgs
	decodeErr := decoder.Decode(&card)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)

	list, err := client.GetList(card.ListID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}
	list.AddCard(&trello.Card{Name: card.CardName, Desc: card.Description}, trello.Defaults())

	message := Message{"true", "Card added successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}

//MoveCard trello
func MoveCard(responseWriter http.ResponseWriter, request *http.Request) {

	var apiKey = os.Getenv("API_KEY")
	var token = os.Getenv("ACCESS_TOKEN")

	decoder := json.NewDecoder(request.Body)

	var param TrelloArgs
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponse(responseWriter, decodeErr)
		return
	}

	client := trello.NewClient(apiKey, token)

	card, err := client.GetCard(param.CardID, trello.Defaults())
	if err != nil {
		result.WriteErrorResponse(responseWriter, err)
		return
	}

	moveErr := card.MoveToList(param.ListID, trello.Defaults())
	if moveErr != nil {
		result.WriteErrorResponse(responseWriter, moveErr)
		return
	}

	message := Message{"true", "Card added successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)

}
