package route

import (
    "github.com/gorilla/mux"
    trello "github.com/heaptracetechnology/microservice-trello/trello"
    "net/http"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "HealthCheck",
        "GET",
        "/health",
        trello.HealthCheck,
    },
    Route{
        "GetCards",
        "POST",
        "/getCards",
        trello.GetCards,
    },
    Route{
        "GetBoard",
        "POST",
        "/getBoard",
        trello.GetBoard,
    },
    Route{
        "GetLists",
        "POST",
        "/getLists",
        trello.GetLists,
    },
    Route{
        "AddCard",
        "POST",
        "/addCard",
        trello.AddCard,
    },
    Route{
        "MoveCard",
        "POST",
        "/moveCard",
        trello.MoveCard,
    },
    Route{
        "SubscribeCard",
        "POST",
        "/subscribe",
        trello.SubscribeCard,
    },
    Route{
        "CopyCard",
        "POST",
        "/copyCard",
        trello.CopyCard,
    },
    Route{
        "CreateBoard",
        "POST",
        "/createBoard",
        trello.CreateBoard,
    },
    Route{
        "DeleteBoard",
        "POST",
        "/deleteBoard",
        trello.DeleteBoard,
    },
    Route{
        "CreateList",
        "POST",
        "/createList",
        trello.CreateList,
    },
    Route{
        "GetAllBoards",
        "POST",
        "/getallBoards",
        trello.GetAllBoards,
    },
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }
    return router
}
