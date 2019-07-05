package trello

import (
	"bytes"
	"encoding/json"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	accessToken = os.Getenv("TRELLO_ACCESS_TOKEN")
	apiKey      = os.Getenv("TRELLO_API_KEY")
	boardID     = os.Getenv("TRELLO_BOARD_ID")
	listID      = os.Getenv("TRELLO_LIST_ID")
	cardID      = os.Getenv("TRELLO_CARD_ID")
	username    = os.Getenv("TRELLO_USERNAME")
)

var _ = Describe("Get Card details from Trello without environment variables", func() {

	os.Setenv("ACCESS_TOKEN", "")
	os.Setenv("API_KEY", "")

	trello := TrelloArgs{BoardID: boardID}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getcards", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCards)
	handler.ServeHTTP(recorder, request)

	Describe("Get card details from board", func() {
		Context("card details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get Card details from Trello with invalid input", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getcards", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCards)
	handler.ServeHTTP(recorder, request)

	Describe("Get card details from board", func() {
		Context("card details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get Card details from Trello", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: boardID}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getcards", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCards)
	handler.ServeHTTP(recorder, request)

	Describe("Get card details from board", func() {
		Context("card details", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get Board details from Trello with param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getboard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBoard)
	handler.ServeHTTP(recorder, request)

	Describe("Get board details from board", func() {
		Context("board details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get Board details from Trello without board ID", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: ""}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getboard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBoard)
	handler.ServeHTTP(recorder, request)

	Describe("Get board details from board", func() {
		Context("board details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get Board details from Trello", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: boardID}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getboard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetBoard)
	handler.ServeHTTP(recorder, request)

	Describe("Get board details from board", func() {
		Context("board details", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get list details from Trello with param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getlists", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetLists)
	handler.ServeHTTP(recorder, request)

	Describe("Get list details from board", func() {
		Context("list details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get list details from Trello without board ID", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: ""}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getlists", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetLists)
	handler.ServeHTTP(recorder, request)

	Describe("Get list details from board", func() {
		Context("list details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get list details from Trello", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: boardID}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/getlists", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetLists)
	handler.ServeHTTP(recorder, request)

	Describe("Get list details from board", func() {
		Context("list details", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Add Card to list in Trello with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/addcard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AddCard)
	handler.ServeHTTP(recorder, request)

	Describe("Get card details from board", func() {
		Context("card details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Add Card to list in Trello without list ID", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: boardID, CardName: "Test Card", Description: "Test card description"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/addcard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AddCard)
	handler.ServeHTTP(recorder, request)

	Describe("Get card details from board", func() {
		Context("card details", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Add Card to list in Trello", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: boardID, ListID: listID, CardName: "Test Card", Description: "Test card description"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/addcard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(AddCard)
	handler.ServeHTTP(recorder, request)

	Describe("Get card details from board", func() {
		Context("card details", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create list in trello board with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createlist", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateList)
	handler.ServeHTTP(recorder, request)

	Describe("create list in board", func() {
		Context("create list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create list in trello board without board ID", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{ListName: "Test List"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createlist", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateList)
	handler.ServeHTTP(recorder, request)

	Describe("create list in board", func() {
		Context("create list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create list in trello board without list name", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: boardID}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createlist", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateList)
	handler.ServeHTTP(recorder, request)

	Describe("create list in board", func() {
		Context("create list", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create list in trello board", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: boardID, ListName: "Test List"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createlist", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateList)
	handler.ServeHTTP(recorder, request)

	Describe("create list in board", func() {
		Context("create list", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Move Card to list in Trello with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/movecard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MoveCard)
	handler.ServeHTTP(recorder, request)

	Describe("move card from list", func() {
		Context("move card", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Move Card to list in Trello without list id", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{CardID: cardID, CardName: "Test Card", Description: "Test card description"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/movecard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MoveCard)
	handler.ServeHTTP(recorder, request)

	Describe("move card from list", func() {
		Context("move card", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Move Card to list in Trello without card id", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{ListID: listID, CardName: "Test Card", Description: "Test card description"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/movecard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(MoveCard)
	handler.ServeHTTP(recorder, request)

	Describe("move card from list", func() {
		Context("move card", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Copy Card to list in Trello with invalid param", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/copycard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CopyCard)
	handler.ServeHTTP(recorder, request)

	Describe("copy card from list", func() {
		Context("copy card", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Copy Card to list in Trello without list id", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{CardID: cardID}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/copycard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CopyCard)
	handler.ServeHTTP(recorder, request)

	Describe("copy card from list", func() {
		Context("copy card", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Copy Card to list in Trello without card id", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{ListID: listID}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/copycard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CopyCard)
	handler.ServeHTTP(recorder, request)

	Describe("copy card from list", func() {
		Context("copy card", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Copy Card to list in Trello", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{CardID: cardID, ListID: listID}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/copycard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CopyCard)
	handler.ServeHTTP(recorder, request)

	Describe("copy card from list", func() {
		Context("copy card", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create board in Trello with invalid param ", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createboard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBoard)
	handler.ServeHTTP(recorder, request)

	Describe("create board in trello", func() {
		Context("create board", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create board in Trello without board name", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardName: ""}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createboard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBoard)
	handler.ServeHTTP(recorder, request)

	Describe("create board in trello", func() {
		Context("create board", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create board in Trello", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardName: "Test Board"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createboard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBoard)
	handler.ServeHTTP(recorder, request)

	Describe("create board in trello", func() {
		Context("create board", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete board in Trello with invalid param ", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/deleteboard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteBoard)
	handler.ServeHTTP(recorder, request)

	Describe("delete board in trello", func() {
		Context("delete board", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete board in Trello without board id", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{BoardID: ""}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/deleteboard", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteBoard)
	handler.ServeHTTP(recorder, request)

	Describe("delete board in trello", func() {
		Context("delete board", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Subscribe trello for card without board ID", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	data := TrelloArgs{ListID: listID}
	sub := Subscribe{Endpoint: "https://webhook.site/3cee781d-0a87-4966-bdec-9635436294e9",
		ID:        "1",
		IsTesting: true,
		Data:      data,
	}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(sub)
	if err != nil {
		fmt.Println(" request err :", err)
	}
	req, err := http.NewRequest("POST", "/subscribe", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SubscribeCard)
	handler.ServeHTTP(recorder, req)

	Describe("Subscribe", func() {
		Context("Subscribe", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Subscribe trello for card", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	data := TrelloArgs{BoardID: boardID, ListID: listID}
	sub := Subscribe{Endpoint: "https://webhook.site/3cee781d-0a87-4966-bdec-9635436294e9",
		ID:        "1",
		IsTesting: true,
		Data:      data,
	}
	requestBody := new(bytes.Buffer)
	err := json.NewEncoder(requestBody).Encode(sub)
	if err != nil {
		fmt.Println(" request err :", err)
	}
	req, err := http.NewRequest("POST", "/subscribe", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SubscribeCard)
	handler.ServeHTTP(recorder, req)

	Describe("Subscribe", func() {
		Context("Subscribe", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get boards for user in Trello with invalid param ", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/boardforuser", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBoards)
	handler.ServeHTTP(recorder, request)

	Describe("get boards for user in trello", func() {
		Context("get boards for user", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get boards for user in Trello without username", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{Username: ""}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/boardforuser", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBoards)
	handler.ServeHTTP(recorder, request)

	Describe("Get boards for user in trello", func() {
		Context("Get boards for user", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Get boards for user in Trello", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)
	os.Setenv("API_KEY", apiKey)

	trello := TrelloArgs{Username: username}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(trello)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/boardforuser", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllBoards)
	handler.ServeHTTP(recorder, request)

	Describe("Get boards for user in trello", func() {
		Context("Get boards for user", func() {
			It("Should result http.StatusStatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
