package trello

import (
	"encoding/json"
	//"fmt"
	trello "github.com/adlio/trello"
	result "github.com/heaptracetechnology/microservice-trello/result"
	"net/http"
	"os"
)

//TrelloArgs
type TrelloArgs struct {
	BoardID string `json:"board_id,omitempty"`
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
